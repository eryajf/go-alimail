package alimail

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

// DepartmentService 部门服务
type DepartmentService struct{ *Client }

type Department struct {
	ID                       string    `json:"id"`                       // 部门ID，根部门id为$root
	Name                     string    `json:"name"`                     // 部门名称
	ParentID                 string    `json:"parentId"`                 // 父部门 id
	CreatedTime              time.Time `json:"createdTime"`              // 创建时间
	HasUsers                 bool      `json:"hasUsers"`                 // 是否包含用户
	HasSubDepartments        bool      `json:"hasSubDepartments"`        // 是否包含子部门
	CanManage                bool      `json:"canManage"`                // 当前登录角色是否可管理该部门
	HiddenExcludeUsers       []string  `json:"hiddenExcludeUsers"`       // 部门隐藏后，哪些白名单帐号 id 可访问该部门，仅管理员或授权应用可访问，可修改
	HiddenExcludeDepartments []string  `json:"hiddenExcludeDepartments"` // 部门隐藏后，哪些部门 id 可访问该部门，仅管理员或授权应用可访问，可修改
	IsHidden                 bool      `json:"isHidden"`                 // 是否隐藏，仅管理员或授权应用可访问，可修改
	Managers                 []string  `json:"managers"`                 // 部门主管的 id 列表，仅管理员或授权应用可访问，可修改
	Email                    string    `json:"email"`                    // 部门邮件组地址，仅管理员或授权应用可修改
}

// Get 获取部门信息，需要传入部门ID，其中根部门ID为$root
func (d *DepartmentService) Get(ctx context.Context, deptId string) (*Organization, error) {
	if deptId == "" {
		return nil, fmt.Errorf("domain is required")
	}
	path := fmt.Sprintf("/v2/departments/%s", deptId)

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

type listDepartmentByIdsRsp struct {
	Departments []Department `json:"departments"`
}

// ListByIds 根据部门id获取部门的基本信息，最大支持 100
func (d *DepartmentService) ListByIds(ctx context.Context, ids []string) ([]Department, error) {
	path := "/v2/departments/listByIds"
	if len(ids) == 0 {
		return nil, fmt.Errorf("ids can't be empty")
	}
	if len(ids) > 100 {
		return nil, fmt.Errorf("ids can't be more than 100")
	}

	body, err := json.Marshal(map[string][]string{"ids": ids})
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request body: %w", err)
	}
	resp, err := d.doRequest(ctx, MethodGet, path, BaseHeader, body)
	if err != nil {
		return nil, fmt.Errorf("request failed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		var dataObj listDepartmentByIdsRsp
		if err := json.NewDecoder(resp.Body).Decode(&dataObj); err != nil {
			return nil, fmt.Errorf("failed to decode response: %w", err)
		}
		return dataObj.Departments, nil
	}
	return nil, parseAPIError(resp)
}

type BaseModifyReq struct {
	Name                     string   `json:"name"`                     // 部门名称
	ParentID                 string   `json:"parentId"`                 // 父部门 id
	HiddenExcludeUsers       []string `json:"hiddenExcludeUsers"`       // 部门隐藏后，哪些白名单帐号 id 可访问该部门，仅管理员或授权应用可访问，可修改
	HiddenExcludeDepartments []string `json:"hiddenExcludeDepartments"` // 部门隐藏后，哪些部门 id 可访问该部门，仅管理员或授权应用可访问，可修改
	IsHidden                 bool     `json:"isHidden"`                 // 是否隐藏，仅管理员或授权应用可访问，可修改
	Managers                 []string `json:"managers"`                 // 部门主管的 id 列表，仅管理员或授权应用可访问，可修改
	Email                    string   `json:"email"`                    // 部门邮件组地址，仅管理员或授权应用可修改
}

type CreateDepartmentReq struct {
	BaseModifyReq
}

// Create 创建部门
func (u *DepartmentService) Create(ctx context.Context, req CreateDepartmentReq) (*Department, error) {
	path := "/v2/departments"

	body, err := json.Marshal(req.BaseModifyReq)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request: %w", err)
	}

	resp, err := u.doRequest(ctx, MethodPost, path, BaseHeader, body)
	if err != nil {
		return nil, fmt.Errorf("request failed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		var dataObj Department
		if err := json.NewDecoder(resp.Body).Decode(&dataObj); err != nil {
			return nil, fmt.Errorf("failed to decode response: %w", err)
		}
		return &dataObj, nil
	}
	return nil, parseAPIError(resp)
}

type UpdateDepartmentReq struct {
	ID string `json:"id"`
	BaseModifyReq
}

// Update 更新部门信息
func (d *DepartmentService) Update(ctx context.Context, req UpdateDepartmentReq) error {
	if req.ID == "" {
		return fmt.Errorf("id  can't be empty at the same time")
	}
	path := "/v2/departments/" + req.ID

	body, err := json.Marshal(req.BaseModifyReq)
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

// Delete 删除部门
func (d *DepartmentService) Delete(ctx context.Context, deptId string) error {
	if deptId == "" {
		return fmt.Errorf("id  can't be empty at the same time")
	}
	path := "/v2/departments/" + deptId
	resp, err := d.doRequest(ctx, MethodDelete, path, BaseHeader, nil)
	if err != nil {
		return fmt.Errorf("request failed: %w", err)
	}
	defer resp.Body.Close()

	return parseAPIError(resp)
}
