# Server-Management-System

## Yêu cầu

- Với hệ điều hành Windows cần có Docker desktop, Docker Compose

- Cài đặt extension Swagger UI trên Google Chrome: `https://chrome.google.com/webstore/detail/swagger-ui/liacakmdhalagfjlfdofigfoiocghoej`

## Các bước chạy system

### Bước 1: Build Docker Compose

`docker-compose.yml`: docker-compose build

### Bước 2: Khởi động các container

docker-compose up -d

### Bước 3

go run ./static

### Bước 4

Mở extension Swagger UI
Copy và dán các đường link sau vào hai tab riêng biệt trong extension Swagger UI:
Tab 1: `http://localhost:3000/server.swagger.json`
Tab 2: `http://localhost:3000/openapiv2/auth/auth.swagger.json`

### Giải thích chi tiết các bước

- **Yêu cầu**: Hướng dẫn người dùng cài đặt trước extension Swagger UI trên Google Chrome, cung cấp liên kết tới Chrome Web Store để dễ dàng cài đặt.
- **Bước 1: Build Docker Compose**: Sử dụng lệnh `docker-compose build` để build tất cả các container được định nghĩa trong tệp `docker-compose.yml`.
- **Bước 2: Khởi động các container**: Chạy các container ở chế độ nền với lệnh `docker-compose up -d`.
- **Bước 3: Chạy ứng dụng Go**: Chạy ứng dụng Go bằng lệnh `go run ./static`.
- **Bước 4: Cấu hình Swagger UI**: Hướng dẫn người dùng mở extension Swagger UI và dán các URL của file Swagger JSON vào hai tab khác nhau để cấu hình Swagger UI.
