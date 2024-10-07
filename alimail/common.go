package alimail

const (
	BaseUrl  = "https://alimail-cn.aliyuncs.com"
	TokenURL = BaseUrl + "/oauth2/v2.0/token"

	SchemeHTTP  = "http"
	SchemeHTTPS = "https"

	MethodGet    = "GET"
	MethodPut    = "PUT"
	MethodPost   = "POST"
	MethodDelete = "DELETE"
	MethodPatch  = "PATCH"
	MethodHead   = "HEAD"
)

var BaseHeader = map[string]string{
	"Content-Type": "application/json",
}

type EmailAccountStatus string

// 邮箱账号状态
const (
	NORMAL EmailAccountStatus = "NORMAL" // 正常
	FREEZE EmailAccountStatus = "FREEZE" // 被冻结
)

type EmailAccountType string

// 邮箱账号类型
const (
	EMPLOYEE EmailAccountType = "EMPLOYEE" // 员工
	SERVICE  EmailAccountType = "SERVICE"  // 共享账号
)
