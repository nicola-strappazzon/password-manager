package card

import (
	"log"

	"github.com/goccy/go-yaml"
)

type Card struct {
	Host     string `yaml:"host"`
	OTP      string `yaml:"otp"`
	Password string `yaml:"password"`
	Port     string `yaml:"port"`
	URL      string `yaml:"url"`
	Username string `yaml:"username"`
	AWS      struct {
		Region          string `yaml:"region"`
		AccountId       string `yaml:"account_id"`
		AccessKey       string `yaml:"access_key"`
		SecretAccessKey string `yaml:"secret_access_key"`
	}
}

func New(in string) (c Card) {
	if err := yaml.Unmarshal([]byte(in), &c); err != nil {
		log.Fatal(err)
	}

	return
}
