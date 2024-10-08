package alimail

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

// OrganizationService 组织服务
type OrganizationService struct{ *Client }

type Organization struct {
	ID                string    `json:"id"`                // 组织ID
	Name              string    `json:"name"`              // 组织名称
	Introduction      string    `json:"introduction"`      // 组织简介
	Telephone         string    `json:"telephone"`         // 组织电话
	Address           string    `json:"address"`           // 组织地址
	PreferredLanguage string    `json:"preferredLanguage"` // 首选语言
	Domain            string    `json:"domain"`            // 组织域名
	CreatedTime       time.Time `json:"createdTime"`       // 创建时间
}

// Get 获取组织信息
func (d *OrganizationService) Get(ctx context.Context) (*Organization, error) {
	path := "/v2/organization/$current"

	resp, err := d.doRequest(ctx, MethodGet, path, BaseHeader, nil)
	if err != nil {
		return nil, fmt.Errorf("request failed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		var dataObj Organization
		if err := json.NewDecoder(resp.Body).Decode(&dataObj); err != nil {
			return nil, fmt.Errorf("failed to decode response: %w", err)
		}
		return &dataObj, nil
	}
	return nil, parseAPIError(resp)
}

type UpdateOrganizationReq struct {
	Name              string `json:"name"`
	Introduction      string `json:"introduction"`
	Telephone         string `json:"telephone"`
	Address           string `json:"address"`
	PreferredLanguage string `json:"preferredLanguage"`
}

// update 更新组织信息
func (d *OrganizationService) Update(ctx context.Context, req UpdateOrganizationReq) error {
	path := "/v2/organization/$current"

	body, err := json.Marshal(req)
	if err != nil {
		return fmt.Errorf("failed to marshal request: %w", err)
	}

	resp, err := d.doRequest(ctx, MethodPatch, path, BaseHeader, body)
	if err != nil {
		return fmt.Errorf("request failed: %w", err)
	}
	defer resp.Body.Close()

	return parseAPIError(resp)
}
