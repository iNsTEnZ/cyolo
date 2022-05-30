package configuration

import (
	"os"
	"strconv"
)

func EnvVerInt(env string, defaultValue int) int {
	tm := os.Getenv(env)

	if tm == "" {
		return defaultValue
	}

	i, err := strconv.Atoi(tm)

	if err != nil {
		return defaultValue
	}

	return i
}

func EnvVerStr(env string, defaultValue string) string {
	envVar := os.Getenv(env)

	if envVar == "" {
		return defaultValue
	}

	return envVar
}

func EnvVerBool(env string, defaultValue bool) bool {
	envVar := os.Getenv(env)

	if envVar == "" {
		return defaultValue
	}

	b, err := strconv.ParseBool(envVar)

	if err != nil {
		return defaultValue
	}

	return b
}
