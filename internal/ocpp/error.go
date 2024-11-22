package ocpp

const (
	ERROR CallType = 4
)

type ErrorCode = string
type ErrorDescription = string
type ErrorDetails = any

type OCPPError struct {
	CallType
	UniqueId
	ErrorCode
	ErrorDescription
	ErrorDetails
}
