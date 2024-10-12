package alimail

import (
	"time"
)

// GroupService 邮件组服务
type GroupService struct{ *Client }

type Group struct {
	ID                        string    `json:"id"`
	Name                      string    `json:"name"`
	Email                     string    `json:"email"`
	Status                    string    `json:"status"`
	CreatedTime               int64     `json:"createdTime"`
	ItemCount                 int64     `json:"itemCount"`
	Admins                    []string  `json:"admins"`
	NeedMessageDeliveryReview bool      `json:"needMessageDeliveryReview"`
	MessageDeliveryReviewers  []string  `json:"messageDeliveryReviewers"`
	AllowedSenders            []string  `json:"allowedSenders"`
	AllowedSenderPolicy       string    `json:"allowedSenderPolicy"`
	Type                      string    `json:"type"`
	DynamicMatchRule          string    `json:"dynamicMatchRule"`
	ItemCountLimit            int64     `json:"itemCountLimit"`
	IsHidden                  bool      `json:"isHidden"`
	Creator                   string    `json:"creator"`
	LastRecvMessageTime       time.Time `json:"lastRecvMessageTime"`
	IgnoreAutoReplyOfMember   bool      `json:"ignoreAutoReplyOfMember"`
}
