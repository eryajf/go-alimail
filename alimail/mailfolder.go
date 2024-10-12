package alimail

// MailFolderService 邮件文件夹服务
type MailFolderService struct{ *Client }

type MailFolder struct {
	ID               string `json:"id"`
	DisplayName      string `json:"displayName"`
	ParentFolderID   string `json:"parentFolderId"`
	ChildFolderCount int    `json:"childFolderCount"`
	TotalItemCount   int    `json:"totalItemCount"`
	UnreadItemCount  int    `json:"unreadItemCount"`
	Extensions       struct {
		Property1 string `json:"property1"`
		Property2 string `json:"property2"`
	} `json:"extensions"`
}
