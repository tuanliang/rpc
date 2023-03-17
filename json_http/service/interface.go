package service

const (
	SERVICE_NAME = "HelloService"
)

type CalcRequest struct {
	A int `json:"a"`
	B int `json:"b"`
}
type HelloService interface {
	Hello(request string, response *string) error
	Calc(req *CalcRequest, response *int) error
}
