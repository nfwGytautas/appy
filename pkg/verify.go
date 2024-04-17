package appy_pkg

func Verify() error {
	// Try to read config
	_, err := ReadConfig()
	if err != nil {
		return err
	}

	return nil
}
