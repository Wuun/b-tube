package serializer

import (
	"btube/conf"
	"encoding/json"
	"fmt"

	"gopkg.in/go-playground/validator.v8"
)

const (
	//ServerUnmarshalJSONErr is the err code when server
	//unmarshal JSON format data.
	ServerUnmarshalJSONErr = 50001
	//ServerQueryMySQLErr is the serr code when server
	//want to get data from mysqlbut get an error.
	ServerQueryMySQLErr = 5002
	//ServerQueryRedisErr is an error code when server
	//want to get data from redis but get an error.
	ServerQueryRedisErr = 5003
	//ParamsErr ,error by params invalid.
	ParamsErr = 5004
	//OthersErr ,error occure by other reason.
	OthersErr = 5005
)

//ErrResponse is a factory of error response by using validator
//which is a helpful package for validate struct's type we want in go.
func ErrResponse(err error) Response {
	if ve, ok := err.(validator.ValidationErrors); ok {
		for _, e := range ve {
			field := conf.Translate(fmt.Sprintf("Field.%s", e.Field))
			tag := conf.Translate(fmt.Sprintf("Tag.Valid.%s", e.Tag))
			return Response{
				StatusCode: OthersErr,
				Msg:        fmt.Sprintf("%s%s", field, tag),
				Error:      fmt.Sprint(err),
			}
		}
	}
	if _, ok := err.(*json.UnmarshalTypeError); ok {
		return Response{
			StatusCode: ServerUnmarshalJSONErr,
			Msg:        "JSON is not a valid format.",
			Error:      fmt.Sprint(err),
		}
	}

	return Response{
		StatusCode: ParamsErr,
		Msg:        "params not match.",
		Error:      fmt.Sprint(err),
	}
}
