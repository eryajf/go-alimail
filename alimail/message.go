package alimail

import (
	"time"
)

// MessageService 邮件内容服务
type MessageService struct{ *Client }

type Message struct {
	InternetMessageID      string `json:"internetMessageId"`
	Subject                string `json:"subject"`
	Summary                string `json:"summary"`
	Priority               string `json:"priority"`
	IsReadReceiptRequested bool   `json:"isReadReceiptRequested"`
	From                   struct {
		Email string `json:"email"`
		Name  string `json:"name"`
	} `json:"from"`
	ToRecipients []struct {
		Email string `json:"email"`
		Name  string `json:"name"`
	} `json:"toRecipients"`
	CcRecipients []struct {
		Email string `json:"email"`
		Name  string `json:"name"`
	} `json:"ccRecipients"`
	BccRecipients []struct {
		Email string `json:"email"`
		Name  string `json:"name"`
	} `json:"bccRecipients"`
	Sender struct {
		Email string `json:"email"`
		Name  string `json:"name"`
	} `json:"sender"`
	ReplyTo []struct {
		Email string `json:"email"`
		Name  string `json:"name"`
	} `json:"replyTo"`
	Body struct {
		BodyText string `json:"bodyText"`
		BodyHTML string `json:"bodyHtml"`
	} `json:"body"`
	InternetMessageHeaders struct {
		Property1 string `json:"property1"`
		Property2 string `json:"property2"`
	} `json:"internetMessageHeaders"`
	FolderID             string    `json:"folderId"`
	ID                   string    `json:"id"`
	HasAttachments       bool      `json:"hasAttachments"`
	IsRead               bool      `json:"isRead"`
	ConversationID       string    `json:"conversationId"`
	SentDateTime         time.Time `json:"sentDateTime"`
	LastModifiedDateTime time.Time `json:"lastModifiedDateTime"`
	ReceivedDateTime     time.Time `json:"receivedDateTime"`
	Tags                 []string  `json:"tags"`
}
