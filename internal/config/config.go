package config

import (
	"os"
	"strings"
)

// Config 구조체는 애플리케이션 설정을 저장합니다.
type Config struct {
	Port            string
	TrustedProxies  []string
	GinMode         string
}

// NewConfig는 환경 변수에서 설정을 로드하여 새 Config 인스턴스를 반환합니다.
func NewConfig() *Config {
	return &Config{
		Port:           getEnv("PORT", "8080"),
		TrustedProxies: getTrustedProxies(),
		GinMode:        getEnv("GIN_MODE", "debug"),
	}
}

// getEnv는 환경 변수 값을 가져오거나 기본값을 반환합니다.
func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

// getTrustedProxies는 TRUSTED_PROXIES 환경 변수에서 신뢰할 수 있는 프록시 목록을 가져옵니다.
func getTrustedProxies() []string {
	proxiesStr := getEnv("TRUSTED_PROXIES", "192.168.1.2")
	if proxiesStr == "" {
		return []string{}
	}
	return strings.Split(proxiesStr, ",")
}