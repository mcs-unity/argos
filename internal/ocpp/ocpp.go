package ocpp

func Response(id UniqueId, p Payload) []any {
	return []any{RESPONSE, id, p}
}

func ResponseConv(p []any) (*ResponseData, error) {
	return &ResponseData{}, nil
}

func Request(id UniqueId, a Action, p Payload) []any {
	return []any{CALL, id, a, p}
}

func RequestConv(p []any) (*RequestData, error) {
	return &RequestData{}, nil
}
