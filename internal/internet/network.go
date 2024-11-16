package network

/*
   contain nic controller for handling internet
   this controller will be responsible for keeping
   the charger online. add event capable
*/

func (i *network) Connect(url []byte) error {
	return nil
}
func (i *network) Disconnect() error {
	return nil
}

func New() INetwork {
	return &network{
		READY,
		nil,
	}
}
