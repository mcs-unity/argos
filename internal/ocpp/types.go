package ocpp

const (
	CALL     CallType = 2
	RESPONSE CallType = 3
	ERROR    CallType = 4
)

type CallType = float64
type UniqueId = string
type Action string
type Payload = any

type ErrorCode = string
type ErrorDescription = string
type ErrorDetails = any

type Frame struct {
	Call CallType
	UUID UniqueId
}

type ResponseData struct {
	Frame
	Data Payload
}

type RequestData struct {
	Action Action
	ResponseData
}

type OCPPError struct {
	Frame
	ErrorCode
	ErrorDescription
	ErrorDetails
}
