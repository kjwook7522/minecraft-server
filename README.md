# Minecraft Server Manager

GCP 기반의 마인크래프트 서버 관리 시스템입니다. 비용 효율적인 운영을 위해 필요할 때만 서버를 구동할 수 있습니다.

## 주요 기능

- GCP Spot VM을 활용한 저비용 마인크래프트 서버 운영
- Docker 기반의 마인크래프트 서버 관리
- 웹 기반 서버 관리 인터페이스
- 서버 시작/중지 자동화

## 기술 스택

- Backend: Go (Cloud Run)
- Frontend: Vue3 + Vite
- Infrastructure: Terraform, Docker
- Cloud: Google Cloud Platform (GCP)

## 시작하기

### 사전 요구사항

- Go 1.21 이상
- Node.js 18 이상
- Docker & Docker Compose
- GCP 계정 및 프로젝트 설정
- Terraform

### 설치 및 실행

1. 백엔드 서버 실행
```bash
cd backend
go mod download
go run cmd/main.go
```

2. 프론트엔드 개발 서버 실행
```bash
cd frontend
npm install
npm run dev
```

3. 인프라 설정
```bash
cd infra
terraform init
terraform apply
```

## 프로젝트 구조

- `backend/`: Go 기반 API 서버
- `frontend/`: Vue3 기반 관리 웹 UI
- `infra/`: Terraform 설정 및 Docker 구성
