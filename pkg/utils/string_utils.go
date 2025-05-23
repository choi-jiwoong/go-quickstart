package utils

import "strings"

// IsEmpty는 문자열이 비어있는지 확인합니다.
func IsEmpty(s string) bool {
	return len(strings.TrimSpace(s)) == 0
}

// IsNotEmpty는 문자열이 비어있지 않은지 확인합니다.
func IsNotEmpty(s string) bool {
	return !IsEmpty(s)
}

// DefaultIfEmpty는 문자열이 비어있으면 기본값을 반환합니다.
func DefaultIfEmpty(s, defaultValue string) string {
	if IsEmpty(s) {
		return defaultValue
	}
	return s
}