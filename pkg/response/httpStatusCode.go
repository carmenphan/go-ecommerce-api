package response

const (
	ErrorCodeSuccess      = 20001
	ErrorCodeFail         = 40001
	ErrorCodeParamInvalid = 2003
)

var msg = map[int]string{
	ErrorCodeFail:         "Fail",
	ErrorCodeSuccess:      "Success",
	ErrorCodeParamInvalid: "Email Is Invalid",
}
