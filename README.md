# Go Quickstart

간단한 Go 웹 서버 프로젝트입니다. Gin 프레임워크와 GORM을 사용하여 구현되었습니다.

## 프로젝트 구조

```
/
├── cmd/                  # 애플리케이션의 메인 진입점
│   └── server/           # 서버 애플리케이션
│       └── main.go       # 메인 애플리케이션 코드
├── internal/             # 외부에서 임포트할 수 없는 패키지
│   ├── api/              # API 핸들러
│   ├── database/         # 데이터베이스 연결 관리
│   ├── middleware/       # 미들웨어
│   ├── models/           # 데이터 모델
│   ├── repository/       # 데이터 접근 레이어
│   └── config/           # 설정 관련 코드
├── pkg/                  # 외부에서 임포트할 수 있는 패키지
│   └── utils/            # 유틸리티 함수
├── scripts/              # 스크립트 파일
│   └── init_db.sql       # 데이터베이스 초기화 스크립트
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
- MySQL 5.7 이상
- Docker (선택 사항)

### 데이터베이스 설정

1. MySQL 서버에 접속:
   ```
   mysql -u root -p
   ```

2. 초기화 스크립트 실행:
   ```
   source scripts/init_db.sql
   ```

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

## API 엔드포인트

### 기본 API
- `GET /`: 클라이언트 IP 정보 반환
- `GET /ping`: 상태 확인 엔드포인트 (pong 응답)

### 인증 API
- `POST /login`: 사용자 로그인

### 사용자 관리 API (인증 필요)
- `GET /user/:id`: 특정 ID의 사용자 정보 조회 (관리자: 모든 사용자, 일반 사용자: 본인만)
- `PUT /user/:id`: 사용자 정보 업데이트 (관리자: 모든 사용자, 일반 사용자: 본인만)

### 관리자 전용 API (관리자 권한 필요)
- `GET /users`: 모든 사용자 목록 조회
- `POST /user`: 새 사용자 생성
- `DELETE /user/:id`: 사용자 삭제

## 권한 관리

애플리케이션은 두 가지 사용자 역할을 지원합니다:

1. **ADMIN**: 모든 API에 접근 가능
2. **USER**: 본인의 정보만 조회 및 수정 가능

## API 요청 예시

### 로그인 (POST /login)
```json
{
  "username": "admin",
  "password": "password123"
}
```

응답:
```json
{
  "id": 1,
  "username": "admin",
  "email": "admin@example.com",
  "role": "ADMIN",
  "token": "admin-token"
}
```

### 인증이 필요한 API 호출
```
GET /user/1
Authorization: Bearer admin-token
```

### 사용자 생성 (POST /user) - 관리자 전용
```
POST /user
Authorization: Bearer admin-token
Content-Type: application/json

{
  "username": "newuser",
  "email": "newuser@example.com",
  "password": "password123",
  "role": "USER"
}
```

## 환경 변수

- `PORT`: 서버 포트 (기본값: 8080)
- `GIN_MODE`: Gin 모드 설정 (debug, release, test)
- `TRUSTED_PROXIES`: 신뢰할 수 있는 프록시 IP 목록 (쉼표로 구분)
- `DB_HOST`: 데이터베이스 호스트 (기본값: localhost)
- `DB_PORT`: 데이터베이스 포트 (기본값: 3306)
- `DB_USER`: 데이터베이스 사용자 (기본값: root)
- `DB_PASSWORD`: 데이터베이스 비밀번호 (기본값: rootpassword)
- `DB_NAME`: 데이터베이스 이름 (기본값: MAIN)