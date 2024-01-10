package common

import "os"

const (
	BLACK int = 0x000000
	WHITE int = 0xffffff
)

func CheckErr(err error) {
	if err != nil {
		panic(NewErrorResponse(0, "An Error occurs", err.Error()))
	}
}

type ErrorResponse struct {
	Code    int
	Message string
	Details string
}

func NewErrorResponse(code int, msg string, details string) *ErrorResponse {
	return &ErrorResponse{Code: code, Message: msg, Details: details}
}

func GetDecodedFileContent(path string) (string, error) {
	bytes2, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}
	return string(bytes2), err
}
func GetEncodedFileContent(path string) ([]byte, error) {
	bytes2, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}
	return bytes2, err
}
