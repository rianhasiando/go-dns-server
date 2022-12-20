package header

type ResponseCode int

const (
	ResponseCodeNoError        ResponseCode = 0
	ResponseCodeFormatError    ResponseCode = 1
	ResponseCodeServerFailure  ResponseCode = 2
	ResponseCodeNameError      ResponseCode = 3
	ResponseCodeNotImplemented ResponseCode = 4
	ResponseCodeRefused        ResponseCode = 5
)
