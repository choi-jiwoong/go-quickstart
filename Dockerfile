FROM golang:1.24-alpine AS builder

WORKDIR /app

# 의존성 파일 복사 및 다운로드
COPY go.mod go.sum ./
RUN go mod download

# 소스 코드 복사
COPY . .

# 애플리케이션 빌드
RUN CGO_ENABLED=0 GOOS=linux go build -o /app/server ./cmd/server

# 실행 이미지
FROM alpine:latest

WORKDIR /app

# 타임존 설정
RUN apk --no-cache add tzdata
ENV TZ=Asia/Seoul

# 빌드된 바이너리 복사
COPY --from=builder /app/server .
COPY .env.example .env

# 포트 노출
EXPOSE 8080

# 애플리케이션 실행
CMD ["./server"]