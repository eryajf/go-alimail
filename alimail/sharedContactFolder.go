package alimail

import (
	"time"
)

// SharedContactFolderService 公共联系人分组服务
type SharedContactFolderService struct{ *Client }

type SharedContactFolder struct {
	ID                       string    `json:"id"`
	Name                     string    `json:"name"`
	ParentID                 string    `json:"parentId"`
	CreatedTime              time.Time `json:"createdTime"`
	HiddenExcludeUsers       []string  `json:"hiddenExcludeUsers"`
	HiddenExcludeDepartments []string  `json:"hiddenExcludeDepartments"`
	IsHidden                 bool      `json:"isHidden"`
	ContactCount             int64     `json:"contactCount"`
	ChildFolderCount         int64     `json:"childFolderCount"`
}
