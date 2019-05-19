package common

import (
	"time"
)

type CommonError struct {
	Errors map[string]interface{} `json:"errors"`
}

// Warp the error info in a object
func NewError(key string, err error) CommonError {
	res := CommonError{}
	res.Errors = make(map[string]interface{})
	res.Errors[key] = err.Error()
	return res
}

func MakeTimeStamp() int64 {
	return time.Now().UnixNano() / int64(time.Second)
}
