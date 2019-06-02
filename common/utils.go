package common

import (
	"log"
	"os"
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

// func myLogger() (*log.Logger, error) {

// 	f, err := os.OpenFile("text.log",
// 		os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
// 	if err != nil {
// 		log.Println(err)
// 	}
// 	defer f.Close()

// 	logger := log.New(f, "prefix", log.LstdFlags)
// 	logger.Println("text to append")
// 	logger.Println("more text to append")

// 	return logger, err

// }

func GetOrCreateDir(path string) string {

	fileInfo, err := os.Stat(path)

	if os.IsNotExist(err) {

		err = os.Mkdir(path, 0644)
		log.Println("Create dir", fileInfo)
		return path
	}
	return ""

}

func GetOrCreateDirs(path string) string {

	fileInfo, err := os.Stat(path)

	if os.IsNotExist(err) {

		err = os.MkdirAll(path, 0644)
		log.Println("Create dir", fileInfo)
		return path
	}
	return ""

}

func Trace(msg string) func() {
	start := time.Now()
	log.Printf("enter %s", msg)
	return func() { log.Printf("exit %s (%s)", msg, time.Since(start)) }
}
