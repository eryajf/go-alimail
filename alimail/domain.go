package alimail

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

// DomainService 域名服务
type DomainService struct{ *Client }

type Domain struct {
	ID             string     `json:"id"`             // 域名ID
	Domain         string     `json:"domain"`         // 域名
	Aliases        []string   `json:"aliases"`        // 域别名
	MasterDomain   string     `json:"masterDomain"`   // 主域名
	MasterDomainID string     `json:"masterDomainId"` // 主域名ID
	IsMaster       bool       `json:"isMaster"`       // 是否是主域名
	IsMxVerified   bool       `json:"isMxVerified"`   // 是否MX记录已验证
	CreatedTime    time.Time  `json:"createdTime"`    // 创建时间
	AliasInfos     []struct { // 域别名信息
		Alias      string `json:"alias"`      // 域别名
		IsVerified bool   `json:"isVerified"` // 是否已验证
	} `json:"aliasInfos"`
	IsVerified bool `json:"isVerified"` // 是否已验证
}

// listDomainsResponse 定义用于解析 API 返回的 JSON
type listDomainsResponse struct {
	Domains []Domain `json:"domains"`
}

// List 列出组织下所有域名信息
func (d *DomainService) List(ctx context.Context) ([]Domain, error) {
	path := "/v2/domains"

	resp, err := d.doRequest(ctx, MethodGet, path, BaseHeader, nil)
	if err != nil {
		return nil, fmt.Errorf("request failed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	var listResp listDomainsResponse
	if err := json.NewDecoder(resp.Body).Decode(&listResp); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	return listResp.Domains, nil
}

// Get 根据域名（或域别名）获取域信息
func (d *DomainService) Get(ctx context.Context, identifier string) (*Domain, error) {
	if identifier == "" {
		return nil, fmt.Errorf("identifier is required")
	}
	path := fmt.Sprintf("/v2/domains/%s", identifier)

	resp, err := d.doRequest(ctx, MethodGet, path, BaseHeader, nil)
	if err != nil {
		return nil, fmt.Errorf("request failed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	var dataObj Domain
	if err := json.NewDecoder(resp.Body).Decode(&dataObj); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	return &dataObj, nil
}
