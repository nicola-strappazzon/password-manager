package card

import (
	"log"

	"github.com/goccy/go-yaml"
)

type Card struct {
	Certificate   string `yaml:"certificate"`
	Email         string `yaml:"email"`
	Host          string `yaml:"host"`
	Name          string `yaml:"name"`
	Notes         string `yaml:"notes"`
	OTP           string `yaml:"otp"`
	Password      string `yaml:"password"`
	Port          string `yaml:"port"`
	RecoveryCodes string `yaml:"recovery_codes"`
	RecoveryKey   string `yaml:"recovery_key"`
	SecretKey     string `yaml:"secret_key"`
	Serial        string `yaml:"serial"`
	Token         string `yaml:"token"`
	URL           string `yaml:"url"`
	Username      string `yaml:"username"`
	AWS           struct {
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
