package alimail

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"sync"
	"time"

	"golang.org/x/time/rate"
)

// Client AliMail API客户端
type Client struct {
	appID     string
	appSecret string

	token       string
	tokenExpiry time.Time
	tokenMutex  sync.RWMutex

	httpClient *http.Client

	// Services
	Domain              *DomainService
	User                *UserService
	Organization        *OrganizationService
	Department          *DepartmentService
	Group               *GroupService
	MailFolder          *MailFolderService
	Message             *MessageService
	SharedContact       *SharedContactService
	SharedContactFolder *SharedContactFolderService
}

// NewClient 创建一个新的Client实例
func NewClient(appID, appSecret string) *Client {
	c := &Client{
		appID:      appID,
		appSecret:  appSecret,
		httpClient: &http.Client{Timeout: 10 * time.Second},
	}
	c.Domain = &DomainService{c}
	c.User = &UserService{c}
	c.Organization = &OrganizationService{c}
	c.Department = &DepartmentService{c}
	c.Group = &GroupService{c}
	c.MailFolder = &MailFolderService{c}
	c.Message = &MessageService{c}
	c.SharedContact = &SharedContactService{c}
	c.SharedContactFolder = &SharedContactFolderService{c}
	return c
}

// TokenResponse 表示获取Token的响应结构
type TokenResponse struct {
	TokenType   string `json:"token_type"`
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
}

// getToken 获取或者刷新Token
func (c *Client) getToken(ctx context.Context) (string, error) {
	// 读取Token时加读锁
	c.tokenMutex.RLock()
	if c.token != "" && time.Now().Before(c.tokenExpiry) {
		token := c.token
		c.tokenMutex.RUnlock()
		return token, nil
	}
	c.tokenMutex.RUnlock()

	// 获取新Token时加写锁
	c.tokenMutex.Lock()
	defer c.tokenMutex.Unlock()

	// 再次检查是否其他协程已经刷新了Token
	if c.token != "" && time.Now().Before(c.tokenExpiry) {
		return c.token, nil
	}

	data := url.Values{}
	data.Set("grant_type", "client_credentials")
	data.Set("client_id", c.appID)
	data.Set("client_secret", c.appSecret)

	req, err := http.NewRequestWithContext(ctx, MethodPost, TokenURL, bytes.NewBufferString(data.Encode()))
	if err != nil {
		return "", err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		bodyBytes, _ := io.ReadAll(resp.Body)
		return "", fmt.Errorf("failed to get token: status %d, body: %s", resp.StatusCode, string(bodyBytes))
	}

	var tokenResp TokenResponse
	if err := json.NewDecoder(resp.Body).Decode(&tokenResp); err != nil {
		return "", err
	}

	if tokenResp.AccessToken == "" {
		return "", errors.New("empty access_token received")
	}

	// 提前60秒过期
	c.token = tokenResp.AccessToken
	c.tokenExpiry = time.Now().Add(time.Duration(tokenResp.ExpiresIn-60) * time.Second)

	return c.token, nil
}

// 全局速率限制器（单个域内所有应用共享）
var (
	globalRateLimiter   *rate.Limiter
	rateLimiterOnce     sync.Once
	rateLimiterMaxBurst = 40
	rateLimiterRate     = rate.Limit(40) // 每秒40次
)

// 初始化全局速率限制器
func getGlobalRateLimiter() *rate.Limiter {
	rateLimiterOnce.Do(func() {
		globalRateLimiter = rate.NewLimiter(rateLimiterRate, rateLimiterMaxBurst)
	})
	return globalRateLimiter
}

// doRequest 发送API请求，自动处理Token和全局速率限制
func (c *Client) doRequest(ctx context.Context, method, path string, headers map[string]string, body []byte) (*http.Response, error) {
	// 全局速率限制
	limiter := getGlobalRateLimiter()
	if err := limiter.Wait(ctx); err != nil {
		return nil, err
	}

	// 获取Token
	token, err := c.getToken(ctx)
	if err != nil {
		return nil, err
	}

	// 构建完整URL
	fullURL := BaseUrl + path
	req, err := http.NewRequestWithContext(ctx, method, fullURL, bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", fmt.Sprintf("bearer %s", token))
	for k, v := range headers {
		req.Header.Set(k, v)
	}

	// 发送请求
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	// 检查速率限制响应
	if resp.StatusCode == 429 {
		bodyBytes, _ := io.ReadAll(resp.Body)
		resp.Body.Close()
		return nil, fmt.Errorf("rate limit exceeded: HTTP 429, body: %s", string(bodyBytes))
	}

	return resp, nil
}
