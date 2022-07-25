package config

import (
	"encoding/json"
	"os"
)

type ConfigDatabase struct {
	Host string `json:"host"`
	Port int    `json:"port"`
	User string `json:"user"`
	Pass string `json:"pass"`
	Name string `json:"name"`
}

type Config struct {
	Database struct {
		SurfTimer ConfigDatabase `json:"surf"`
		VIP       ConfigDatabase `json:"vip"`
	} `json:"database"`
	Directory struct {
		Maps string `json:"maps"`
	} `json:"directory"`
	Notifications struct {
		Kofi struct {
			OnPayment string `json:"on_payment"`
		} `json:"kofi"`
	} `json:"notifications"`
	Secrets struct {
		Kofi struct {
			VerificationToken string `json:"verification_token"`
		} `json:"kofi"`
	} `json:"secrets"`
	Server struct {
		Mode           string   `json:"mode"`
		Port           int      `json:"port"`
		TrustedProxies []string `json:"trusted_proxies"`
	} `json:"server"`
}

func NewConfig(configPath string) (*Config, error) {
	cfg := &Config{}

	file, err := os.Open(configPath)

	if err != nil {
		return cfg, err
	}

	defer file.Close()

	d := json.NewDecoder(file)

	if err := d.Decode(&cfg); err != nil {
		return cfg, err
	}

	return cfg, nil
}
