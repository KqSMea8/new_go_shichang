package common

type Conf struct {
	ColumnMap    map[string]string `yaml:"ColumnMap"`
}

// Response 接口输出
// Code:错误 -1 ，正常 0
type Response struct {
	Code    int         `json:"code"`
	Data    interface{} `json:"data"`
	Message interface{} `json:"message"`
}