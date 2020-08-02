package global

func Init(path string) error {
	err := InitConfig(path)
	if err != nil {
		return err
	}
	return nil
}
