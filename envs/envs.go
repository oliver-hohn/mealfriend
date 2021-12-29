package envs

import (
	"fmt"
	"os"
)

func MustGetEnv(key string) string {
	val := os.Getenv(key)

	if val == "" {
		panic(fmt.Errorf("expected an ENV value for: %s", key))
	}

	return val
}
