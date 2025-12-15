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
	} `yaml:"aws"`
}

func New(in string) (c Card) {
	if err := yaml.Unmarshal([]byte(in), &c); err != nil {
		log.Fatal(err)
	}

	return
}

func (c Card) ToString() string {
	out, err := yaml.Marshal(&c)

	if err != nil {
		log.Fatal(err)
	}

	return string(out)
}

func (c Card) Fields() []string {
	return []string{
		"certificate",
		"email",
		"host",
		"name",
		"notes",
		"otp",
		"password",
		"port",
		"recovery_codes",
		"recovery_key",
		"secret_key",
		"serial",
		"token",
		"url",
		"username",
		"aws.region",
		"aws.account_id",
		"aws.access_key",
		"aws.secret_access_key",
	}
}

func (c Card) GetValue(in string) (out string) {
	switch in {
	case "certificate":
		out = c.Certificate
	case "email":
		out = c.Email
	case "host":
		out = c.Host
	case "name":
		out = c.Name
	case "notes":
		out = c.Notes
	case "otp":
		out = c.OTP
	case "password":
		out = c.Password
	case "port":
		out = c.Port
	case "recovery_codes":
		out = c.RecoveryCodes
	case "recovery_key":
		out = c.RecoveryKey
	case "secret_key":
		out = c.SecretKey
	case "serial":
		out = c.Serial
	case "token":
		out = c.Token
	case "url":
		out = c.URL
	case "username":
		out = c.Username
	case "aws.region":
		out = c.AWS.Region
	case "aws.account_id":
		out = c.AWS.AccountId
	case "aws.access_key":
		out = c.AWS.AccessKey
	case "aws.secret_access_key":
		out = c.AWS.SecretAccessKey
	}

	return out
}

func (c *Card) SetValue(key, value string) {
	switch key {
	case "certificate":
		c.Certificate = value
	case "email":
		c.Email = value
	case "host":
		c.Host = value
	case "name":
		c.Name = value
	case "notes":
		c.Notes = value
	case "otp":
		c.OTP = value
	case "password":
		c.Password = value
	case "port":
		c.Port = value
	case "recovery_codes":
		c.RecoveryCodes = value
	case "recovery_key":
		c.RecoveryKey = value
	case "secret_key":
		c.SecretKey = value
	case "serial":
		c.Serial = value
	case "token":
		c.Token = value
	case "url":
		c.URL = value
	case "username":
		c.Username = value
	case "aws.region":
		c.AWS.Region = value
	case "aws.account_id":
		c.AWS.AccountId = value
	case "aws.access_key":
		c.AWS.AccessKey = value
	case "aws.secret_access_key":
		c.AWS.SecretAccessKey = value
	}
}
