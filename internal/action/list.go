package action

import "github.com/mcs-unity/ocpp-simulator/internal/ocpp"

var list map[ocpp.Action]IAction

func init() {
	list := make(map[ocpp.Action]IAction, 1)
	list[HEARTBEAT] = heartbeat{}
	list[BOOTNOTIFICATION] = bootNotification{}

}
