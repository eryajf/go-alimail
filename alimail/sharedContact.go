package alimail

import (
	"time"
)

// SharedContactService 公共联系人服务
type SharedContactService struct{ *Client }

type SharedContact struct {
	ID           string    `json:"id"`
	Name         string    `json:"name"`
	FolderID     string    `json:"folderId"`
	Email        string    `json:"email"`
	WorkPhone    string    `json:"workPhone"`
	Phone        string    `json:"phone"`
	HomeAddress  string    `json:"homeAddress"`
	CreatedTime  time.Time `json:"createdTime"`
	ModifiedTime time.Time `json:"modifiedTime"`
	CompanyName  string    `json:"companyName"`
	JobTitle     string    `json:"jobTitle"`
	WorkAddress  string    `json:"workAddress"`
	FolderPath   []string  `json:"folderPath"`
	Nickname     string    `json:"nickname"`
	Manager      string    `json:"manager"`
	Info         string    `json:"info"`
	IsHidden     bool      `json:"isHidden"`
}
