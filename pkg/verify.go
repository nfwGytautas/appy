package appy_pkg

import (
	"fmt"

	appy_share "github.com/nfwGytautas/appy/share"
)

func Verify() error {
	// Try to read config
	_, err := appy_share.ReadConfig()
	if err != nil {
		return err
	}

	fmt.Println("Configuration is valid!")

	return nil
}
