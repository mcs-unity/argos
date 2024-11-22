package ocpp

func Response(id UniqueId, p Payload) []any {
	return []any{RESPONSE, id, p}
}

func ResponseConv(p []any) (*ResponseData, error) {
	if len(p) < 3 {
		return nil, &OCPPError{}
	}
	r := &ResponseData{}

	call, ok := p[0].(float64)
	if !ok {
		return nil, &OCPPError{}
	}
	r.Call = call

	uuid, ok := p[1].(string)
	if !ok {
		return nil, &OCPPError{}
	}
	r.UUID = uuid

	r.Data = p[2]
	return r, nil
}

func Request(id UniqueId, a Action, p Payload) []any {
	return []any{CALL, id, a, p}
}

func RequestConv(p []any) (*RequestData, error) {
	if len(p) < 4 {
		return nil, &OCPPError{}
	}
	r := &RequestData{}

	call, ok := p[0].(float64)
	if !ok {
		return nil, &OCPPError{}
	}
	r.Call = call

	uuid, ok := p[1].(string)
	if !ok {
		return nil, &OCPPError{}
	}
	r.UUID = uuid

	action, ok := p[2].(Action)
	if !ok {
		return nil, &OCPPError{}
	}
	r.Action = action

	r.Data = p[3]
	return r, nil
}
