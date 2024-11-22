package ocpp

const (
	CALL     CallType = 2
	RESPONSE CallType = 3
)

type CallType = float64
type UniqueId = string
type Action string
type Payload = any

type Frame struct {
	Call CallType
	UUID UniqueId
	Data Payload
}

type ResponseData struct {
	Frame
}

type RequestData struct {
	Action Action
	Frame
}
