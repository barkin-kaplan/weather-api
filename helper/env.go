package helper

import (
	"fmt"
	"os"
	"strconv"
)
func CheckAndGetEnvString(key string) string {
	value := os.Getenv(key)
	if value == ""{
		panic(fmt.Sprintf("Environment variable %s is not set", key))
	}

	return value
}

func CheckAndGetEnvInteger(key string) (int, error) {
	valueString := CheckAndGetEnvString(key)
	value, err := strconv.Atoi(valueString)
	if err != nil {
		fmt.Printf("Invalid %s %d. Has to be an integer", key, value)
		return 0, err
	}

	return value, nil
}

func CheckAndGetEnvIntegerWithDefault(key string, defaultValue int) (int, error) {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue, nil
	}

	valueInt, err := strconv.Atoi(value)
	if err != nil {
		fmt.Printf("Invalid %s %s. Has to be an integer", key, value)
		return defaultValue, err
	}
	return valueInt, nil
}