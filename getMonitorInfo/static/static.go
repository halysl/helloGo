package static

// SQL
var (
	UserName = "user"
	Password = "passwd"
	IP       = "*"
	Port     = "3306"
	DbName   = "*"
)

var (
	baseURL       = "http://*:9090"
	baseTargetURL = baseURL + "/api/v1/targets"
	baseQueryURL  = baseURL + "/api/v1/query?query=%s&time=%s"
)
