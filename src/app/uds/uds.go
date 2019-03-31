package uds

//Respons 定义统一的返回格式
type Respons struct {
	Code int
	Msg  string
	Data interface{}
}
