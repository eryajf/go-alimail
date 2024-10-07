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

// 邮箱账号状态
const (
	NORMAL = "NORMAL" // 正常
	FREEZE = "FREEZE" // 被冻结
)

// 邮箱账号类型
const (
	EMPLOYEE = "EMPLOYEE" // 员工
	SERVICE  = "SERVICE"  // 共享账号
)
