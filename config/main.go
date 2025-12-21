package config

import (
	"path"

	"github.com/nicola-strappazzon/pm/env"
)

const DATA_DIR = ".password-manager"

// func GetHomeDir() string {
// 	base, err := os.UserHomeDir()
// 	check.Check(err)

// 	return base
// }

// rename to GetAbsolutePath
func GetDataDirectoryFrom(in string) string {
	// return path.Join(GetHomeDir(), env.Get("PM_PATH", DATA_DIR), in)
	return path.Join(env.Get("PM_PATH", DATA_DIR), in)
}

func GetDataDirectory() string {
	return GetDataDirectoryFrom("")
}

func GetPrivateKeyPath() string {
	return env.Get("PM_PRIVATEKEY", "")
}

func GetPublicKeyPath() string {
	return env.Get("PM_PUBLICKEY", "")
}
