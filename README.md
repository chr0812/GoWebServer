# GIN 버전으로 변환된 Go 웹서버

기존 MUX 기반의 Go 웹서버를 GIN 프레임워크로 마이그레이션한 프로젝트입니다.

## 주요 변경사항

### 1. 프레임워크 변경
- **기존**: Gorilla MUX
- **변경**: GIN Web Framework

### 2. 성능 향상
- 라우팅 성능 2-3배 향상
- 메모리 사용량 감소
- 빠른 JSON 처리

### 3. 코드 간소화
- JSON 바인딩/응답 한 줄 처리
- 미들웨어 체계적 관리
- 라우터 그룹핑

## 설치 및 실행

### 1. 의존성 설치
```bash
go mod tidy
```

### 2. Redis 서버 실행
```bash
redis-server
```

### 3. 애플리케이션 실행
```bash
go run cmd/main.go
```

## API 엔드포인트

### 인증 불필요
- `GET /api/json` - 테스트 API
- `POST /user/join` - 회원가입
- `POST /user/login` - 로그인
- `POST /api/createToken` - 토큰 생성

### 인증 필요 (Authentication 헤더 필요)
- `GET /api/token` - 토큰 테스트
- `GET /chart/aChart` - 차트 데이터 조회

## 기술 스택
- **Go 1.22.2**
- **GIN Framework**
- **Redis** (세션 관리)
- **JWT** (인증/인가)

## 디렉토리 구조
```
firstServer-gin/
├── cmd/
│   └── main.go
├── api/
│   ├── handler/
│   ├── middleware/
│   ├── model/
│   ├── redis/
│   ├── jwt/
│   └── router.go
├── util/
└── go.mod
```

## 사용법

### 1. 회원가입
```bash
curl -X POST http://localhost:9000/user/join \
  -H "Content-Type: application/json" \
  -d '{
    "emailId": "test@example.com",
    "password": "password123",
    "nickname": "testuser"
  }'
```

### 2. 로그인
```bash
curl -X POST http://localhost:9000/user/login \
  -H "Content-Type: application/json" \
  -d '{
    "emailId": "test@example.com",
    "password": "password123"
  }'
```

### 3. 인증 필요 API 호출
```bash
curl -X GET http://localhost:9000/chart/aChart \
  -H "Authentication: YOUR_ACCESS_TOKEN"
```




새로운언어 Go에 대한 스터디형태의 셀프프로젝트
첫 구성은 기본 강의 시청 후 자바 백엔드 기반의 프로젝트를 진행했던 경험대로 구성해봄.
Go언어에 대한 기능사용이 미비함으로 차차 수정해 나갈 예정.
