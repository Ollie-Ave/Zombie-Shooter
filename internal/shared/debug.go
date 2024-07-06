package shared

import "os"

func IsDebugMode() bool {
	return os.Getenv(DebugModeEnvironmentVariable) == "true"
}
