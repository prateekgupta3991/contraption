package configs

import (
	"encoding/json"
	"os"
)

type Conf struct {
	ServerPort string
}

func InitConfig(file string) (*Conf, error) {
	if file, err := os.Open(file); err != nil {
		return nil, err
	} else {
		decoder := json.NewDecoder(file)
		con := new(Conf)
		if err = decoder.Decode(&con); err != nil {
			return nil, err
		} else {
			return con, nil
		}
	}
}
