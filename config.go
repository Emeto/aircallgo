package aircallgo

import (
	"encoding/json"
	"log"
	"os"
)

type Config struct {
	ID    string `json:"ID"`
	Token string `json:"Token"`
}

func ParseConfig() *Config {
	buf, err := os.ReadFile("./config.json")
	if err != nil {
		log.Fatal("cannot read ./config.json. do you have permission or is it missing?")
	}
	var config Config
	err = json.Unmarshal(buf, &config)
	if err != nil {
		log.Fatal("could not parse config.json")
	}
	return &config
}
