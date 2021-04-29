package serialize

import (
	"school-walker/walker/serialize/ecode"
)

type Response struct {
	Data interface{} `json:"data"`
	Msg  string `json:"msg"`
	Code int    `json:"code"`
}

func BuildErrorResponse(code int) Response {
	return Response{
		Data: "",
		Msg:  ecode.GetMsg(code),
		Code: code,
	}
}

func BuildCustomResponse(data interface{}, msg string, code int) Response {
	return Response{
		Data: data,
		Msg:  msg,
		Code: code,
	}
}
