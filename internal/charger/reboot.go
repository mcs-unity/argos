package charger

import (
	"fmt"
	"time"
)

//	allow charger to reboot (reset connector state and follow OCPP procedure)

func (c *Charger) ResetConnectors() {
	for _, con := range c.connectors {
		t := con.GetTransaction()

		if t == -1 {
			continue
		}

		if err := con.StopTransaction(t); err != nil {
			fmt.Println("err")
		}
	}
}

func hard(c *Charger) {
	if err := c.Stop(); err != nil {
		panic(err)
	}

	time.Sleep(30 * time.Second)
	c.boot = HARD
	// the charger should stop transaction after reboot
}

func soft(c *Charger) {
	if err := c.socket.Close(); err != nil {
		panic(err)
	}

	time.Sleep(10 * time.Second)
	c.boot = SOFT
}

func (c *Charger) Reboot(t RebootType) error {
	defer c.lock.Unlock()
	c.lock.Lock()

	if t == HARD {
		hard(c)
	} else {
		soft(c)
	}

	if err := c.Start(); err != nil {
		panic(err)
	}

	return nil
}
