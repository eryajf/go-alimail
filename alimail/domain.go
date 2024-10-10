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

type BaseDomainReq struct {
	ID     string `json:"id"`     // 用户ID
	Domain string `json:"domain"` // 域名
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

	if resp.StatusCode == http.StatusOK {
		var dataObj listDomainsResponse
		if err := json.NewDecoder(resp.Body).Decode(&dataObj); err != nil {
			return nil, fmt.Errorf("failed to decode response: %w", err)
		}
		return dataObj.Domains, nil
	}
	return nil, parseAPIError(resp)
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

	if resp.StatusCode == http.StatusOK {
		var dataObj Domain
		if err := json.NewDecoder(resp.Body).Decode(&dataObj); err != nil {
			return nil, fmt.Errorf("failed to decode response: %w", err)
		}
		return &dataObj, nil
	}
	return nil, parseAPIError(resp)
}

// Create 创建新的域名信息
func (d *DomainService) Create(ctx context.Context, domain string) (*Domain, error) {
	path := "/v2/domains"

	body, err := json.Marshal(map[string]string{"domain": domain})
	if err != nil {
		return nil, err
	}
	resp, err := d.doRequest(ctx, MethodPost, path, BaseHeader, body)
	if err != nil {
		return nil, fmt.Errorf("request failed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		var dataObj Domain
		if err := json.NewDecoder(resp.Body).Decode(&dataObj); err != nil {
			return nil, fmt.Errorf("failed to decode response: %w", err)
		}
		return &dataObj, nil
	}
	return nil, parseAPIError(resp)
}

// Delete 删除域名信息
func (d *DomainService) Delete(ctx context.Context, domain string) (*Domain, error) {
	if domain == "" {
		return nil, fmt.Errorf("domain is required")
	}
	path := fmt.Sprintf("/v2/domains/%s", domain)

	resp, err := d.doRequest(ctx, MethodDelete, path, BaseHeader, nil)
	if err != nil {
		return nil, fmt.Errorf("request failed: %w", err)
	}
	defer resp.Body.Close()

	return nil, parseAPIError(resp)
}

type verifyDomainRsp struct {
	IsOwner bool `json:"isOwner"`
}

// Verify 验证域名信息
func (d *DomainService) Verify(ctx context.Context, domain string) (*verifyDomainRsp, error) {
	if domain == "" {
		return nil, fmt.Errorf("domain is required")
	}
	path := fmt.Sprintf("/v2/domains/%s/verify", domain)

	resp, err := d.doRequest(ctx, MethodPost, path, BaseHeader, nil)
	if err != nil {
		return nil, fmt.Errorf("request failed: %w", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode == http.StatusOK {
		var dataObj verifyDomainRsp
		if err := json.NewDecoder(resp.Body).Decode(&dataObj); err != nil {
			return nil, fmt.Errorf("failed to decode response: %w", err)
		}
		return &dataObj, nil
	}
	return nil, parseAPIError(resp)
}

type AddDoaminAliasReq struct {
	BaseDomainReq
	Alias string `json:"alias"`
}

// AddDomainAlias 添加域别名
func (d *DomainService) AddDomainAlias(ctx context.Context, req AddDoaminAliasReq) error {
	if req.ID == "" && req.Domain == "" {
		return fmt.Errorf("id and domain can't be empty at the same time")
	}
	var path string
	if req.Domain != "" {
		path = fmt.Sprintf("/v2/domains/%s/aliases", req.Domain)
	} else {
		path = fmt.Sprintf("/v2/domains/%s/aliases", req.ID)
	}

	body, err := json.Marshal(req)
	if err != nil {
		return fmt.Errorf("failed to marshal request: %w", err)
	}

	resp, err := d.doRequest(ctx, MethodPost, path, BaseHeader, body)
	if err != nil {
		return fmt.Errorf("request failed: %w", err)
	}
	defer resp.Body.Close()

	return parseAPIError(resp)
}

type DeleteDoaminAliasReq struct {
	BaseDomainReq
	Alias string `json:"alias"`
}

// DeleteDoaminAliasReq 删除域别名
func (d *DomainService) DeleteDoaminAliasReq(ctx context.Context, req AddDoaminAliasReq) error {
	if req.ID == "" && req.Domain == "" {
		return fmt.Errorf("id and domain can't be empty at the same time")
	}
	var path string
	if req.Domain != "" {
		path = fmt.Sprintf("/v2/domains/%s/aliases/%s", req.Domain, req.Alias)
	} else {
		path = fmt.Sprintf("/v2/domains/%s/aliases/%s", req.ID, req.Alias)
	}

	body, err := json.Marshal(req)
	if err != nil {
		return fmt.Errorf("failed to marshal request: %w", err)
	}

	resp, err := d.doRequest(ctx, MethodDelete, path, BaseHeader, body)
	if err != nil {
		return fmt.Errorf("request failed: %w", err)
	}
	defer resp.Body.Close()

	return parseAPIError(resp)
}

type renameDomainRsp struct {
	RenameQuota int `json:"renameQuota"` // 剩余更换域名的次数
}

// Rename 变更域名
func (d *DomainService) Rename(ctx context.Context, domain string) (*renameDomainRsp, error) {
	path := "/v2/domains/rename"

	resp, err := d.doRequest(ctx, MethodPost, path, BaseHeader, nil)
	if err != nil {
		return nil, fmt.Errorf("request failed: %w", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode == http.StatusOK {
		var dataObj renameDomainRsp
		if err := json.NewDecoder(resp.Body).Decode(&dataObj); err != nil {
			return nil, fmt.Errorf("failed to decode response: %w", err)
		}
		return &dataObj, nil
	}
	return nil, parseAPIError(resp)
}

// Release释放域名，主域名将重置为默认域名
func (d *DomainService) Release(ctx context.Context, domain string) error {
	path := "/v2/domains/release"

	resp, err := d.doRequest(ctx, MethodPost, path, BaseHeader, nil)
	if err != nil {
		return fmt.Errorf("request failed: %w", err)
	}
	defer resp.Body.Close()

	return parseAPIError(resp)
}
