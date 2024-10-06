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
