package card

import (
	// "fmt"
	"log"
	"reflect"
	"strings"

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

func (c Card) Fields() []string {
	return yamlFields(reflect.TypeOf(c), "")
}

func (c Card) Field(in string) string {
	parts := strings.Split(in, ".")
	val := reflect.ValueOf(c)
	typ := reflect.TypeOf(c)

	for _, p := range parts {
		for i := 0; i < typ.NumField(); i++ {
			if typ.Field(i).Tag.Get("yaml") == p {
				val = val.Field(i)
				typ = typ.Field(i).Type
				break
			}
		}
	}

	if val.Kind() != reflect.String {
		return ""
	}

	return val.String()
}

func yamlFields(t reflect.Type, prefix string) []string {
	fields := []string{}

	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}

	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)

		if f.PkgPath != "" {
			continue
		}

		tag := f.Tag.Get("yaml")
		if tag == "" || tag == "-" {
			continue
		}

		name := tag
		if prefix != "" {
			name = prefix + "." + tag
		}

		if f.Type.Kind() == reflect.Struct {
			fields = append(fields, yamlFields(f.Type, name)...)
			continue
		}

		fields = append(fields, name)
	}

	return fields
}

func findByYAMLTag(v reflect.Value, prefix, target string) (string, bool) {
	v = reflect.Indirect(v)
	if !v.IsValid() || v.Kind() != reflect.Struct {
		return "", false
	}

	t := v.Type()

	for i := 0; i < v.NumField(); i++ {
		sf := t.Field(i)
		fv := v.Field(i)

		// si el campo no es exportado, sáltalo
		if sf.PkgPath != "" {
			continue
		}

		tag := sf.Tag.Get("yaml")
		tagName := strings.Split(tag, ",")[0] // por si hay ",omitempty"
		if tagName == "" || tagName == "-" {
			continue
		}

		full := tagName
		if prefix != "" {
			full = prefix + "." + tagName
		}

		// 1) match directo
		if full == target {
			fv = reflect.Indirect(fv)
			if fv.IsValid() && fv.Kind() == reflect.String {
				return fv.String(), true
			}
			// si algún día quieres soportar otros tipos, aquí es donde convertirías
			return "", false
		}

		// 2) bajar recursivamente si es struct
		fv = reflect.Indirect(fv)
		if fv.IsValid() && fv.Kind() == reflect.Struct {
			if s, ok := findByYAMLTag(fv, full, target); ok {
				return s, true
			}
		}
	}

	return "", false
}
