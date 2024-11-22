package ocpp

func (err OCPPError) Format(uuid UniqueId) []Payload {
	return []Payload{ERROR, uuid, err.ErrorCode, err.ErrorDescription, err.ErrorDetails}
}

func (err OCPPError) Error() string {
	return err.ErrorDescription
}
