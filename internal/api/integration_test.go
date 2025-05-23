package api_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/choi-jiwoong/go-quickstart/internal/api"
	"github.com/gin-gonic/gin"
)

// 통합 테스트를 위한 라우터 설정
func setupTestRouter() *gin.Engine {
	gin.SetMode(gin.TestMode)
	router := gin.New()
	
	// 라우트 등록
	router.GET("/", api.RootHandler)
	router.GET("/ping", api.PingHandler)
	
	return router
}

// 여러 엔드포인트를 연속적으로 호출하는 통합 테스트
func TestAPIIntegration(t *testing.T) {
	router := setupTestRouter()
	
	// 테스트 케이스
	testCases := []struct {
		name           string
		method         string
		path           string
		expectedStatus int
		checkResponse  func(*testing.T, *httptest.ResponseRecorder)
	}{
		{
			name:           "루트 경로",
			method:         "GET",
			path:           "/",
			expectedStatus: http.StatusOK,
			checkResponse: func(t *testing.T, w *httptest.ResponseRecorder) {
				var response map[string]interface{}
				if err := json.Unmarshal(w.Body.Bytes(), &response); err != nil {
					t.Fatalf("응답 본문 파싱 실패: %v", err)
				}
				
				if _, exists := response["clientIP"]; !exists {
					t.Error("응답에 'clientIP' 필드가 없습니다")
				}
			},
		},
		{
			name:           "핑 경로",
			method:         "GET",
			path:           "/ping",
			expectedStatus: http.StatusOK,
			checkResponse: func(t *testing.T, w *httptest.ResponseRecorder) {
				var response map[string]interface{}
				if err := json.Unmarshal(w.Body.Bytes(), &response); err != nil {
					t.Fatalf("응답 본문 파싱 실패: %v", err)
				}
				
				if message, exists := response["message"]; !exists || message != "pong" {
					t.Errorf("응답의 'message' 필드가 'pong'이 아닙니다: %v", message)
				}
			},
		},
		{
			name:           "존재하지 않는 경로",
			method:         "GET",
			path:           "/not-found",
			expectedStatus: http.StatusNotFound,
			checkResponse:  nil,
		},
	}
	
	// 테스트 실행
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			req := httptest.NewRequest(tc.method, tc.path, nil)
			w := httptest.NewRecorder()
			
			router.ServeHTTP(w, req)
			
			if w.Code != tc.expectedStatus {
				t.Errorf("예상 상태 코드 %d, 실제 %d", tc.expectedStatus, w.Code)
			}
			
			if tc.checkResponse != nil {
				tc.checkResponse(t, w)
			}
		})
	}
}