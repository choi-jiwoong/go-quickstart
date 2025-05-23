-- 데이터베이스 생성
CREATE DATABASE IF NOT EXISTS MAIN;
USE MAIN;

-- 사용자 테이블 생성
CREATE TABLE IF NOT EXISTS users (
    id         BIGINT AUTO_INCREMENT PRIMARY KEY,
    created_at DATETIME(6) DEFAULT CURRENT_TIMESTAMP(6),
    email      VARCHAR(100) NOT NULL,
    password   VARCHAR(255) NOT NULL,
    role       VARCHAR(20)  NOT NULL,
    updated_at DATETIME(6) DEFAULT CURRENT_TIMESTAMP(6) ON UPDATE CURRENT_TIMESTAMP(6),
    username   VARCHAR(50)  NOT NULL,
    CONSTRAINT UK_username UNIQUE (username)
);

-- 로그인 히스토리 테이블 생성
CREATE TABLE IF NOT EXISTS login_history (
    id         BIGINT AUTO_INCREMENT PRIMARY KEY,
    created_at DATETIME(6) DEFAULT CURRENT_TIMESTAMP(6),
    ip_address VARCHAR(50)  NULL,
    login_time DATETIME(6)  NOT NULL,
    success    BIT          NOT NULL,
    user_agent VARCHAR(255) NULL,
    user_id    BIGINT       NOT NULL,
    CONSTRAINT FK_login_history_user FOREIGN KEY (user_id) REFERENCES users (id)
);

-- 샘플 데이터 삽입
INSERT INTO users (email, password, role, username)
VALUES 
    ('admin@example.com', '$2a$10$N9qo8uLOickgx2ZMRZoMyeIjZAgcfl7p92ldGxad68LJZdL17lhWy', 'ADMIN', 'admin'),
    ('user@example.com', '$2a$10$N9qo8uLOickgx2ZMRZoMyeIjZAgcfl7p92ldGxad68LJZdL17lhWy', 'USER', 'user');