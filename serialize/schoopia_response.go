package serialize

type SchoolResponse struct {
	// Data 用户基本数据
	Data interface{} `json:"data"`
	// Code 请求状态码
	Code int `json:"code"`
	// Time 时间
	Time int64 `json:"time"`
	// Description 描述
	Description string `json:"desc"`
}
