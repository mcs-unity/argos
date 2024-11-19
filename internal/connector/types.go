package connector

type IState interface{}

type State struct{}

type IConnector interface{}

type Connector struct {
	connector uint8
	state     IState
}
