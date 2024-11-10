package charger

//	allow charger to reboot (reset connector state and follow OCPP procedure)

func (c *Charger) Reboot(t RebootType) error {
	defer c.lock.Unlock()
	c.lock.Lock()

	if t == HARD {
		// simulate rebooting the hardware long reboot
		// hard reboot don't be smooth
		// set flag hard boot
	} else {
		// simulate rebooting the firmware short reboot
		// also known as a graceful reboot
		// set flag soft boot
	}

	// Retry start procedure

	return nil
}
