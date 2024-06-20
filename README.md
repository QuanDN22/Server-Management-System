# Server-Management-System

## Yêu cầu

- Với hệ điều hành Windows cần có Docker desktop, Docker Compose

- Cài đặt extension [Swagger UI](https://chrome.google.com/webstore/detail/swagger-ui/liacakmdhalagfjlfdofigfoiocghoej) trên Google Chrome

## Các bước chạy

### Bước 1: Build Docker Compose

`docker-compose.yml`: docker-compose build

### Bước 2: Khởi động các container

`docker-compose up -d`

### Bước 3

`go run ./static`

### Bước 4

Mở extension Swagger UI
Copy và dán các đường link sau vào hai tab riêng biệt trong extension Swagger UI:
Tab 1: `http://localhost:3000/server.swagger.json`
Tab 2: `http://localhost:3000/openapiv2/auth/auth.swagger.json`
