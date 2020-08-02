package global

import "github.com/BurntSushi/toml"

type config struct {
	Host host `toml:"host"`
	AccountInfo accountInfo`toml:"account_info"`
}

type accountInfo struct {
	Token string `toml:"token"`
	AesKey string `toml:"aes_key"`
}

type host struct {
	Ip string `toml:"ip"`
	Port uint32 `toml:"port"`
}

var conf config

func GetConfig() *config {
	return &conf
}

func InitConfig(path string) error {
	_, err := toml.DecodeFile(path, &conf)
	if err != nil {
		return err
	}
	return nil
}