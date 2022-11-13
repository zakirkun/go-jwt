package helper

import "fmt"

type Response struct {
	Message string      `json:"message"`
	Errors  interface{} `json:"errors"`
	Data    interface{} `json:"data"`
}

type Obj struct {
}

func BuildResponse(message string, data interface{}) Response {
	res := Response{
		Message: message,
		Errors:  Obj{},
		Data:    data,
	}
	return res
}

func BuildErrorResponse(message string, err ...string) Response {

	for _, s := range err {
		fmt.Println(s)
	}
	res := Response{
		Message: message,
		Errors:  err,
		Data:    Obj{},
	}
	return res
}
