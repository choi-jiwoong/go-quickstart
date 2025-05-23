package config

import (
	"os"
	"reflect"
	"testing"
)

func TestGetEnv(t *testing.T) {
	// 환경 변수 설정
	os.Setenv("TEST_KEY", "test_value")
	defer os.Unsetenv("TEST_KEY")
	
	// 테스트 케이스
	tests := []struct {
		name         string
		key          string
		defaultValue string
		expected     string
	}{
		{
			name:         "환경 변수가 설정된 경우",
			key:          "TEST_KEY",
			defaultValue: "default",
			expected:     "test_value",
		},
		{
			name:         "환경 변수가 설정되지 않은 경우",
			key:          "NON_EXISTENT_KEY",
			defaultValue: "default",
			expected:     "default",
		},
	}
	
	// 테스트 실행
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := getEnv(tt.key, tt.defaultValue)
			if result != tt.expected {
				t.Errorf("getEnv(%q, %q) = %q, expected %q", tt.key, tt.defaultValue, result, tt.expected)
			}
		})
	}
}

func TestGetTrustedProxies(t *testing.T) {
	// 테스트 케이스
	tests := []struct {
		name           string
		envValue       string
		expectedResult []string
	}{
		{
			name:           "쉼표로 구분된 여러 프록시",
			envValue:       "192.168.1.1,10.0.0.1",
			expectedResult: []string{"192.168.1.1", "10.0.0.1"},
		},
		{
			name:           "단일 프록시",
			envValue:       "192.168.1.1",
			expectedResult: []string{"192.168.1.1"},
		},
		{
			name:           "환경 변수가 설정되지 않은 경우 기본값 사용",
			envValue:       "",
			expectedResult: []string{"192.168.1.2"},
		},
	}
	
	// 테스트 실행
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// 환경 변수 설정
			if tt.envValue != "" {
				os.Setenv("TRUSTED_PROXIES", tt.envValue)
				defer os.Unsetenv("TRUSTED_PROXIES")
			} else {
				os.Unsetenv("TRUSTED_PROXIES")
			}
			
			result := getTrustedProxies()
			if !reflect.DeepEqual(result, tt.expectedResult) {
				t.Errorf("getTrustedProxies() = %v, expected %v", result, tt.expectedResult)
			}
		})
	}
}

func TestNewConfig(t *testing.T) {
	// 환경 변수 설정
	os.Setenv("PORT", "9000")
	os.Setenv("GIN_MODE", "release")
	os.Setenv("TRUSTED_PROXIES", "10.0.0.1,10.0.0.2")
	defer func() {
		os.Unsetenv("PORT")
		os.Unsetenv("GIN_MODE")
		os.Unsetenv("TRUSTED_PROXIES")
	}()
	
	// 설정 생성
	cfg := NewConfig()
	
	// 검증
	if cfg.Port != "9000" {
		t.Errorf("Expected Port to be '9000', got %q", cfg.Port)
	}
	
	if cfg.GinMode != "release" {
		t.Errorf("Expected GinMode to be 'release', got %q", cfg.GinMode)
	}
	
	expectedProxies := []string{"10.0.0.1", "10.0.0.2"}
	if !reflect.DeepEqual(cfg.TrustedProxies, expectedProxies) {
		t.Errorf("Expected TrustedProxies to be %v, got %v", expectedProxies, cfg.TrustedProxies)
	}
}