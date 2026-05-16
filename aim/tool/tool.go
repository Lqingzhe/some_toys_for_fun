package tool

import (
	newerror "aim/pkg/error"
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"net/http"
	"unicode/utf8"
)

func AddSaltByByteLength(length int64) (string, error) {
	saltBytes := make([]byte, length)
	_, err := rand.Read(saltBytes)
	if err != nil {
		return "", newerror.MakeError(http.StatusInternalServerError, newerror.CodeInternalError, "Internal Service Error", fmt.Errorf("Add Salt Generation Failed: %v", err), newerror.LevelFatal)
	}
	return hex.EncodeToString(saltBytes), nil
}
func TypeAssert[T any](input any) (output *T, err error) {
	output, ok := input.(*T)
	if !ok {
		return nil, newerror.MakeError(http.StatusInternalServerError, newerror.CodeInternalError, "Internal Service Error", fmt.Errorf("Database Struct Type Error"), newerror.LevelFatal)
	}
	return output, nil
}
func Encrypt(input string) string {
	h := sha256.New()
	h.Write([]byte(input))
	return hex.EncodeToString(h.Sum(nil))
}
func CalculateLength[T []any | string](input T) int64 {
	switch v := any(input).(type) {
	case []any:
		return int64(len(v))
	case string:
		return int64(utf8.RuneCountInString(v))
	}
	return 0
}
func CalculateByteLength[T []any | string](input T) int64 {
	return int64(len(input))
}
