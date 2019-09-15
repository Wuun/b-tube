package serializer

import (
	"btube/conf"
	"encoding/json"
	"fmt"

	validator "gopkg.in/go-playground/validator.v8"
)

//Response is use to format  server's response.
type Response struct {
	StatusCode int
	Msg        string
	Error      string
	Data       interface{}
}

// ErrorResponse return the information of error.
func ErrorResponse(err error) Response {
	if ve, ok := err.(validator.ValidationErrors); ok {
		for _, e := range ve {
			field := conf.Translate(fmt.Sprintf("Field.%s", e.Field))
			tag := conf.Translate(fmt.Sprintf("Tag.Valid.%s", e.Tag))
			return Response{
				StatusCode: 40001,
				Msg:        fmt.Sprintf("%s%s", field, tag),
				Error:      fmt.Sprint(err),
			}
		}
	}
	if _, ok := err.(*json.UnmarshalTypeError); ok {
		return Response{
			StatusCode: 40001,
			Msg:        "JSON类型不匹配",
			Error:      fmt.Sprint(err),
		}
	}

	return Response{
		StatusCode: 40001,
		Msg:        "参数错误",
		Error:      fmt.Sprint(err),
	}
}
