package helper

import (
	"crypto/sha256"
	"encoding/base64"
	"math/rand"
	"regexp"
	"strings"
)

var defaultPool = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
// Function to generate a random string with a given length from a given pool of characters

func GenerateRandomString(length int, pool *string) string {
	if pool == nil {
		pool = &defaultPool
	}
	result := make([]byte, length)
	for i := 0; i < length; i++ {
		result[i] = (*pool)[rand.Intn(len((*pool)))]
	}
	return string(result)
}

func HashString(input string) string {
	hash := sha256.New()
	hash.Write([]byte(input))
	hashBytes := hash.Sum(nil)
	return base64.StdEncoding.EncodeToString(hashBytes)
}


var emailRegex = regexp.MustCompile(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`)

// IsValidEmail checks if the input string is a valid email address format
func IsValidEmail(email string) bool {
	return emailRegex.MatchString(email)
}

func EnsureTrailingSlash(address string) string {
	if strings.HasSuffix(address, "/") {
		return address
	}
	return address + "/"
}

func RemoveTrailingSlash(s string) string {
	if strings.HasSuffix(s, "/"){
		return strings.TrimSuffix(s, "/")
	}

	return s
}

func GetRelativePathComponents(path string) []string {
	// Trim leading and trailing slashes
	trimmed := strings.Trim(path, "/")

	// Split by slash
	return strings.Split(trimmed, "/")
}