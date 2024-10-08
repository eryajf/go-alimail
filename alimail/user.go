package alimail

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

// UserService 用户服务
type UserService struct{ *Client }

// User 用户信息
type User struct {
	ID             string             `json:"id"`             // 用户ID
	Email          string             `json:"email"`          // 用户邮箱
	EmailAliases   []string           `json:"emailAliases"`   // 用户邮箱别名
	Name           string             `json:"name"`           // 用户姓名
	Nickname       string             `json:"nickname"`       // 用户昵称
	EmployeeNo     string             `json:"employeeNo"`     // 员工编号
	JobTitle       string             `json:"jobTitle"`       // 职位
	WorkLocation   string             `json:"workLocation"`   // 工作地点
	OfficeLocation string             `json:"officeLocation"` // 办公地点
	HomeLocation   string             `json:"homeLocation"`   // 家庭住址
	DepartmentIds  []string           `json:"departmentIds"`  // 部门ID列表
	Phone          string             `json:"phone"`          // 手机号
	WorkPhone      string             `json:"workPhone"`      // 工作电话
	Status         EmailAccountStatus `json:"status"`         // 用户状态
	Avatar         struct {           // 用户头像
		URL     string `json:"url"`     // 头像URL
		BgColor string `json:"bgColor"` // 头像背景颜色
	} `json:"avatar"`
	CreatedTime time.Time `json:"createdTime"` // 创建时间
	FreezedTime time.Time `json:"freezedTime"` // 冻结时间
	CustomName  string    `json:"customName"`  // 自定义名称
	ManagerInfo struct {  // 上级信息
		ID       string `json:"id"`       // 上级ID
		Email    string `json:"email"`    // 上级邮箱
		Name     string `json:"name"`     // 上级姓名
		Nickname string `json:"nickname"` // 上级昵称
	} `json:"managerInfo"`
	ManagerEmail   string   `json:"managerEmail"` // 上级邮箱
	DepartmentPath struct { // 部门路径
		Paths []struct { // 路径
			Segments []string `json:"segments"` // 路径段
		} `json:"paths"`
	} `json:"departmentPath"`
	EnableClientPassword           bool             `json:"enableClientPassword"`           // 是否启用客户端密码
	LastLoginTime                  time.Time        `json:"lastLoginTime"`                  // 最后登录时间
	LastPasswordChangeTime         time.Time        `json:"lastPasswordChangeTime"`         // 最后修改密码时间
	IsInitialPassword              bool             `json:"isInitialPassword"`              // 是否初始密码
	IsHidden                       bool             `json:"isHidden"`                       // 是否隐藏
	LastLoginIP                    string           `json:"lastLoginIp"`                    // 最后登录IP
	Info                           string           `json:"info"`                           // 信息
	ForceChangePasswordNextSignIn  bool             `json:"forceChangePasswordNextSignIn"`  // 下次登录是否强制修改密码
	EmployeeType                   EmailAccountType `json:"employeeType"`                   // 员工类型
	PasswordExpireTime             time.Time        `json:"passwordExpireTime"`             // 密码过期时间
	LastLoginTime4StandardProtocol time.Time        `json:"lastLoginTime4StandardProtocol"` // 最后登录时间（标准协议）
}

type BaseUserReq struct {
	ID    string `json:"id"`    // 用户ID
	Email string `json:"email"` // 用户邮箱
}

// Get 根据id或email获取用户信息,参数传入其一即可
func (d *UserService) Get(ctx context.Context, req BaseUserReq) (*User, error) {
	if req.Email == "" && req.ID == "" {
		return nil, fmt.Errorf("id and email can't be empty at the same time")
	}
	var path string
	if req.Email != "" {
		path = fmt.Sprintf("/v2/users/%s", req.Email)
	} else {
		path = fmt.Sprintf("/v2/users/%s", req.ID)
	}

	resp, err := d.doRequest(ctx, MethodGet, path, BaseHeader, nil)
	if err != nil {
		return nil, fmt.Errorf("request failed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		var dataObj User
		if err := json.NewDecoder(resp.Body).Decode(&dataObj); err != nil {
			return nil, fmt.Errorf("failed to decode response: %w", err)
		}
		return &dataObj, nil
	}
	return nil, parseAPIError(resp)
}

type listUserByIdsResponse struct {
	Users []User `json:"users"`
}

// ListByIds 根据 id 列表批量获取帐号信息，最大支持 100
func (d *UserService) ListByIds(ctx context.Context, ids []string) ([]User, error) {
	path := "/v2/users/listByIds"
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
		var dataObj listUserByIdsResponse
		if err := json.NewDecoder(resp.Body).Decode(&dataObj); err != nil {
			return nil, fmt.Errorf("failed to decode response: %w", err)
		}
		return dataObj.Users, nil
	}
	return nil, parseAPIError(resp)
}

type CreateUserReq struct {
	Email                         string             `json:"email"`                                   // 用户邮箱
	Password                      string             `json:"password"`                                // 用户密码
	Name                          string             `json:"name"`                                    // 用户姓名
	Nickname                      string             `json:"nickname,omitempty"`                      // 用户昵称
	EmployeeNo                    string             `json:"employeeNo,omitempty"`                    // 员工编号
	JobTitle                      string             `json:"jobTitle,omitempty"`                      // 职位
	WorkLocation                  string             `json:"workLocation,omitempty"`                  // 工作地点
	OfficeLocation                string             `json:"officeLocation,omitempty"`                // 办公地点
	HomeLocation                  string             `json:"homeLocation,omitempty"`                  // 家庭住址
	DepartmentIds                 []string           `json:"departmentIds"`                           // 部门ID列表
	Phone                         string             `json:"phone,omitempty"`                         // 手机号
	WorkPhone                     string             `json:"workPhone,omitempty"`                     // 工作电话
	Status                        EmailAccountStatus `json:"status,omitempty"`                        // 用户状态
	CustomName                    string             `json:"customName,omitempty"`                    // 自定义名称
	ManagerEmail                  string             `json:"managerEmail,omitempty"`                  // 上级邮箱
	IsHidden                      bool               `json:"isHidden,omitempty"`                      // 是否隐藏
	Info                          string             `json:"info,omitempty"`                          // 信息
	ForceChangePasswordNextSignIn bool               `json:"forceChangePasswordNextSignIn,omitempty"` // 下次登录是否强制修改密码
	EmployeeType                  EmailAccountType   `json:"employeeType,omitempty"`                  // 员工类型
}

// Create 创建用户
func (u *UserService) Create(ctx context.Context, req CreateUserReq) (*User, error) {
	path := "/v2/users"

	headers := map[string]string{
		"Content-Type": "application/json",
	}

	body, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request: %w", err)
	}

	resp, err := u.doRequest(ctx, MethodPost, path, headers, body)
	if err != nil {
		return nil, fmt.Errorf("request failed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		var dataObj User
		if err := json.NewDecoder(resp.Body).Decode(&dataObj); err != nil {
			return nil, fmt.Errorf("failed to decode response: %w", err)
		}
		return &dataObj, nil
	}
	return nil, parseAPIError(resp)
}

type UpdateUserReq struct {
	BaseUserReq                                      // 用户ID或邮箱
	Name                          string             `json:"name"`                                    // 用户姓名
	Nickname                      string             `json:"nickname,omitempty"`                      // 用户昵称
	EmployeeNo                    string             `json:"employeeNo,omitempty"`                    // 员工编号
	JobTitle                      string             `json:"jobTitle,omitempty"`                      // 职位
	WorkLocation                  string             `json:"workLocation,omitempty"`                  // 工作地点
	OfficeLocation                string             `json:"officeLocation,omitempty"`                // 办公地点
	HomeLocation                  string             `json:"homeLocation,omitempty"`                  // 家庭住址
	DepartmentIds                 []string           `json:"departmentIds"`                           // 部门ID列表
	Phone                         string             `json:"phone,omitempty"`                         // 手机号
	WorkPhone                     string             `json:"workPhone,omitempty"`                     // 工作电话
	Status                        EmailAccountStatus `json:"status,omitempty"`                        // 用户状态
	CustomName                    string             `json:"customName,omitempty"`                    // 自定义名称
	ManagerEmail                  string             `json:"managerEmail,omitempty"`                  // 上级邮箱
	IsHidden                      bool               `json:"isHidden,omitempty"`                      // 是否隐藏
	Info                          string             `json:"info,omitempty"`                          // 信息
	ForceChangePasswordNextSignIn bool               `json:"forceChangePasswordNextSignIn,omitempty"` // 下次登录是否强制修改密码
}

// Update 更新用户信息
func (u *UserService) Update(ctx context.Context, req UpdateUserReq) (*User, error) {
	if req.Email == "" && req.ID == "" {
		return nil, fmt.Errorf("id and email can't be empty at the same time")
	}
	var path string
	if req.Email != "" {
		path = fmt.Sprintf("/v2/users/%s", req.Email)
	} else {
		path = fmt.Sprintf("/v2/users/%s", req.ID)
	}

	body, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request: %w", err)
	}

	resp, err := u.doRequest(ctx, MethodPatch, path, BaseHeader, body)
	if err != nil {
		return nil, fmt.Errorf("request failed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		var dataObj User
		if err := json.NewDecoder(resp.Body).Decode(&dataObj); err != nil {
			return nil, fmt.Errorf("failed to decode response: %w", err)
		}
		return &dataObj, nil
	}
	return nil, parseAPIError(resp)
}

// Delete 删除用户
func (d *UserService) Delete(ctx context.Context, req BaseUserReq) (*User, error) {
	if req.Email == "" && req.ID == "" {
		return nil, fmt.Errorf("id and email can't be empty at the same time")
	}
	var path string
	if req.Email != "" {
		path = fmt.Sprintf("/v2/users/%s", req.Email)
	} else {
		path = fmt.Sprintf("/v2/users/%s", req.ID)
	}

	resp, err := d.doRequest(ctx, MethodDelete, path, BaseHeader, nil)
	if err != nil {
		return nil, fmt.Errorf("request failed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		var dataObj User
		if err := json.NewDecoder(resp.Body).Decode(&dataObj); err != nil {
			return nil, fmt.Errorf("failed to decode response: %w", err)
		}
		return &dataObj, nil
	}
	return nil, parseAPIError(resp)
}

type ChangeUserPasswordReq struct {
	BaseUserReq
	Old string `json:"old"`
	New string `json:"new"`
}

// ChangePassword 修改用户密码
func (d *UserService) ChangePassword(ctx context.Context, req ChangeUserPasswordReq) (*User, error) {
	if req.ID == "" && req.Email == "" {
		return nil, fmt.Errorf("id and email can't be empty at the same time")
	}
	var path string
	if req.Email != "" {
		path = fmt.Sprintf("/v2/users/%s/changePassword", req.Email)
	} else {
		path = fmt.Sprintf("/v2/users/%s/changePassword", req.ID)
	}

	body, err := json.Marshal(map[string]string{"old": req.Old, "new": req.New})
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request: %w", err)
	}

	resp, err := d.doRequest(ctx, MethodDelete, path, BaseHeader, body)
	if err != nil {
		return nil, fmt.Errorf("request failed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		var dataObj User
		if err := json.NewDecoder(resp.Body).Decode(&dataObj); err != nil {
			return nil, fmt.Errorf("failed to decode response: %w", err)
		}
		return &dataObj, nil
	}
	return nil, parseAPIError(resp)
}

type ResetUserPasswordReq struct {
	BaseUserReq
	Password                      string `json:"password"`
	ForceChangePasswordNextSignIn bool   `json:"forceChangePasswordNextSignIn"` // 下次登录是否强制修改密码
}

// ResetPassword 重置用户密码
func (d *UserService) ResetPassword(ctx context.Context, req ResetUserPasswordReq) (*User, error) {
	if req.ID == "" && req.Email == "" {
		return nil, fmt.Errorf("id and email can't be empty at the same time")
	}
	var path string
	if req.Email != "" {
		path = fmt.Sprintf("/v2/users/%s/resetPassword", req.Email)
	} else {
		path = fmt.Sprintf("/v2/users/%s/resetPassword", req.ID)
	}

	body, err := json.Marshal(map[string]string{
		"password":                      req.Password,
		"forceChangePasswordNextSignIn": fmt.Sprintf("%t", req.ForceChangePasswordNextSignIn),
	})
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request: %w", err)
	}

	resp, err := d.doRequest(ctx, MethodPost, path, BaseHeader, body)
	if err != nil {
		return nil, fmt.Errorf("request failed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		var dataObj User
		if err := json.NewDecoder(resp.Body).Decode(&dataObj); err != nil {
			return nil, fmt.Errorf("failed to decode response: %w", err)
		}
		return &dataObj, nil
	}
	return nil, parseAPIError(resp)
}

type AddEmailAliasReq struct {
	BaseUserReq
	Alias     string `json:"alias"`
	AsDefault bool   `json:"asDefault"` // 是否作为默认展示的别名
}

// AddEmailAlias 添加邮箱别名
func (d *UserService) AddEmailAlias(ctx context.Context, req AddEmailAliasReq) (*User, error) {
	if req.ID == "" && req.Email == "" {
		return nil, fmt.Errorf("id and email can't be empty at the same time")
	}
	var path string
	if req.Email != "" {
		path = fmt.Sprintf("/v2/users/%s/emailAliases", req.Email)
	} else {
		path = fmt.Sprintf("/v2/users/%s/emailAliases", req.ID)
	}

	body, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request: %w", err)
	}

	resp, err := d.doRequest(ctx, MethodPost, path, BaseHeader, body)
	if err != nil {
		return nil, fmt.Errorf("request failed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		var dataObj User
		if err := json.NewDecoder(resp.Body).Decode(&dataObj); err != nil {
			return nil, fmt.Errorf("failed to decode response: %w", err)
		}
		return &dataObj, nil
	}
	return nil, parseAPIError(resp)
}

type DeleteEmailAliasReq struct {
	BaseUserReq
	Alias string `json:"alias"`
}

// DeleteEmailAlias 删除邮箱别名
func (d *UserService) DeleteEmailAlias(ctx context.Context, req DeleteEmailAliasReq) (*User, error) {
	if req.ID == "" && req.Email == "" {
		return nil, fmt.Errorf("id and email can't be empty at the same time")
	}
	var path string
	if req.Email != "" {
		path = fmt.Sprintf("/v2/users/%s/emailAliases", req.Email)
	} else {
		path = fmt.Sprintf("/v2/users/%s/emailAliases", req.ID)
	}

	body, err := json.Marshal(map[string]string{"alias": req.Alias})
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request: %w", err)
	}

	resp, err := d.doRequest(ctx, MethodDelete, path, BaseHeader, body)
	if err != nil {
		return nil, fmt.Errorf("request failed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		var dataObj User
		if err := json.NewDecoder(resp.Body).Decode(&dataObj); err != nil {
			return nil, fmt.Errorf("failed to decode response: %w", err)
		}
		return &dataObj, nil
	}
	return nil, parseAPIError(resp)
}
