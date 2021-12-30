package envs

import (
	"fmt"
	"os"
	"strconv"
)

func MustGetEnv(key string) string {
	val := os.Getenv(key)

	if val == "" {
		panic(fmt.Errorf("expected an ENV value for: %s", key))
	}

	return val
}

func MustGetIntEnv(key string) int64 {
	sVal := MustGetEnv(key)

	iVal, err := strconv.Atoi(sVal)

	if err != nil {
		panic(fmt.Errorf("expected an integer value for: %s, received: %s", key, sVal))
	}

	return int64(iVal)
}
