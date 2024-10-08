package alimail

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const (
	BaseUrl  = "https://alimail-cn.aliyuncs.com"
	TokenURL = BaseUrl + "/oauth2/v2.0/token"

	SchemeHTTP  = "http"
	SchemeHTTPS = "https"

	MethodGet    = "GET"
	MethodPut    = "PUT"
	MethodPost   = "POST"
	MethodDelete = "DELETE"
	MethodPatch  = "PATCH"
	MethodHead   = "HEAD"
)

var BaseHeader = map[string]string{
	"Content-Type": "application/json",
}

type EmailAccountStatus string

// 邮箱账号状态
const (
	NORMAL EmailAccountStatus = "NORMAL" // 正常
	FREEZE EmailAccountStatus = "FREEZE" // 被冻结
)

type EmailAccountType string

// 邮箱账号类型
const (
	EMPLOYEE EmailAccountType = "EMPLOYEE" // 员工
	SERVICE  EmailAccountType = "SERVICE"  // 共享账号
)

// APIError 自定义的API错误类型
type APIError struct {
	StatusCode       int    // HTTP状态码
	DetailErrorCode  string `json:"detailErrorCode"`
	Message          string `json:"message"`
	DeveloperMessage string `json:"developerMessage"`
}

// Error 实现 error 接口
func (e *APIError) Error() string {
	return fmt.Sprintf("StatusCode: %d, DetailErrorCode: %s, Message: %s, DeveloperMessage: %s",
		e.StatusCode, e.DetailErrorCode, e.Message, e.DeveloperMessage)
}

// parseAPIError 解析API返回的错误响应
func parseAPIError(resp *http.Response) error {
	// 如果不需要再返回内容，则可直接调用此方法
	if resp.StatusCode == http.StatusOK {
		return nil
	}
	var apiErr APIError
	apiErr.StatusCode = resp.StatusCode
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("failed to read error response body: %w", err)
	}
	if err := json.Unmarshal(body, &apiErr); err != nil {
		return fmt.Errorf("failed to parse error response: %s", string(body))
	}
	return &apiErr
}
