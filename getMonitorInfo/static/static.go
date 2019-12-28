package static

// SQL
var (
	UserName = "user"
	Password = "passwd"
	Ip = "*"
	Port = "3306"
	DbName = "*"
)

var (
	BaseUrl = "http://*:9090"
	BaseTargetUrl = BaseUrl + "/api/v1/targets"
	BaseQueryUrl = BaseUrl + "/api/v1/query?query=%s&time=%s"
)
