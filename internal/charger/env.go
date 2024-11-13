package charger

import (
	"fmt"
	"os"
)

func loadEnv(dir string) error {
	b, err := os.ReadFile(fmt.Sprintf("%s/config/config.json", dir))
	if err != nil {
		return err
	}
	fmt.Println(string(b))
	return nil
}
