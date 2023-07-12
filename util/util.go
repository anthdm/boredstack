package util

import "os"

func IsProd() bool {
	return os.Getenv("PRODUCTION") == "true"
}

func IsDev() bool {
	return !IsProd()
}

func AppEnv() string {
	if IsProd() {
		return "production"
	}
	return "development"
}
