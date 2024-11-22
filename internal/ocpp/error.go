package ocpp

func (err OCPPError) Error() string {
	return err.ErrorDescription
}
