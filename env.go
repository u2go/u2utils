package u2utils

import "os"

func GetEnv(key, defValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defValue
	}
	return value
}
