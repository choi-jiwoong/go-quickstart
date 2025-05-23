# Go Quickstart

간단한 Go 웹 서버 프로젝트입니다. Gin 프레임워크를 사용하여 구현되었습니다.

## 프로젝트 구조

```
/
├── cmd/                  # 애플리케이션의 메인 진입점
│   └── server/           # 서버 애플리케이션
│       └── main.go       # 메인 애플리케이션 코드
├── internal/             # 외부에서 임포트할 수 없는 패키지
│   ├── api/              # API 핸들러
│   ├── middleware/       # 미들웨어
│   └── config/           # 설정 관련 코드
├── pkg/                  # 외부에서 임포트할 수 있는 패키지
│   └── utils/            # 유틸리티 함수
├── .env.example          # 환경 변수 예제 파일
├── Dockerfile            # Docker 이미지 빌드 설정
├── docker-compose.yml    # Docker Compose 설정
├── go.mod                # Go 모듈 정의
├── go.sum                # Go 모듈 체크섬
├── Makefile              # 빌드 스크립트
└── .gitignore            # Git 무시 파일 목록
```

## 시작하기

### 필수 조건

- Go 1.22 이상
- Docker (선택 사항)

### 설치 및 실행

#### 로컬 실행

1. 저장소 클론:
   ```
   git clone https://github.com/yourusername/go-quickstart.git
   cd go-quickstart
   ```

2. 의존성 설치:
   ```
   go mod tidy
   ```

3. 환경 변수 설정:
   ```
   cp .env.example .env
   # .env 파일을 필요에 따라 수정
   ```

4. 애플리케이션 실행:
   ```
   make run
   ```

#### Docker로 실행

1. Docker 이미지 빌드:
   ```
   make docker-build
   ```

2. Docker 컨테이너 실행:
   ```
   make docker-run
   ```

또는 Docker Compose 사용:

```
make docker-compose-up
```

## 사용 가능한 명령어

### 개발 명령어
- `make build`: 애플리케이션 빌드
- `make test`: 테스트 실행
- `make run`: 애플리케이션 실행
- `make clean`: 빌드 결과물 정리
- `make tidy`: Go 모듈 의존성 정리
- `make coverage`: 테스트 커버리지 보고서 생성

### Docker 명령어
- `make docker-build`: Docker 이미지 빌드
- `make docker-run`: Docker 컨테이너 실행
- `make docker-compose-up`: Docker Compose로 서비스 시작
- `make docker-compose-down`: Docker Compose로 서비스 중지

## API 엔드포인트

- `GET /`: 클라이언트 IP 정보 반환
- `GET /ping`: 상태 확인 엔드포인트 (pong 응답)

## 환경 변수

- `PORT`: 서버 포트 (기본값: 8080)
- `GIN_MODE`: Gin 모드 설정 (debug, release, test)
- `TRUSTED_PROXIES`: 신뢰할 수 있는 프록시 IP 목록 (쉼표로 구분)

## 테스트

프로젝트의 모든 패키지에 대한 테스트를 실행하려면:

```
go test ./...
```

특정 패키지의 테스트만 실행하려면:

```
go test ./internal/api
go test ./internal/middleware
go test ./internal/config
go test ./pkg/utils
```

테스트 커버리지 보고서를 생성하려면:

```
go test ./... -coverprofile=coverage.out
go tool cover -html=coverage.out
```