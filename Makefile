gen:	
	protoc -I ./proto \
   	--go_out ./proto --go_opt paths=source_relative \
   	--go-grpc_out ./proto --go-grpc_opt paths=source_relative \
	--grpc-gateway_out ./proto --grpc-gateway_opt paths=source_relative \
   	./proto/auth/auth.proto

gen-key:
	openssl genpkey -algorithm ED25519 -outform pem -out auth.ed
	openssl pkey -in auth.ed -pubout > auth.ed.pub

# run server:
run-auth-server:
	go run ./cmd/auth/main.go

run-grpc-gateway:
	go run .\cmd\grpc-gateway\main.go auth.ed.pub 

# api
token := eyJhbGciOiJFZERTQSIsInR5cCI6IkpXVCJ9.eyJhdWQiOiJhcGkiLCJleHAiOjE3MTYxNDY5MDgsImlhdCI6MTcxNjE0MzkwOCwiaXNzIjoiaHR0cDovL2xvY2FsaG9zdDo1MDAxIiwibmJmIjoxNzE2MTQzOTA4LCJyb2xlcyI6ImFkbWluIiwidXNlciI6ImFkbWluIn0.lo1yWRAAYpkTs9RIsrgiFNAcn5foS4I2jLXr5L4v8XH2k66QX58V7Nt7MAxBDvu4E2WiqCBV4qacvYjPELpNDw
# token := eyJhbGciOiJFZERTQSIsInR5cCI6IkpXVCJ9.eyJhdWQiOiJhcGkiLCJleHAiOjE3MTYxNDY5MTQsImlhdCI6MTcxNjE0MzkxNCwiaXNzIjoiaHR0cDovL2xvY2FsaG9zdDo1MDAxIiwibmJmIjoxNzE2MTQzOTE0LCJyb2xlcyI6InVzZXIiLCJ1c2VyIjoicXVhbjEifQ.3tyZLNfIF7bQA5E1zYgZqyORFy6MK8CQAeyWRih2873jWX7CAx_MKnJPsd6RUIB4mwq3Na0T-MOOAPqY9Oj0Cg

# api of auth server
login: 
	curl -H "Content-Type: application/json" -X POST http://localhost:8080/v1/api/login \
	-d "{\"username\":\"admin\", \"password\":\"pass\"}"

login-user:
	curl -H "Content-Type: application/json" -X POST http://localhost:8080/v1/api/login \
	-d "{\"username\":\"quan2\", \"password\":\"2\"}"

ping-auth-server: 
	curl -s -H "Authorization: Bearer ${token}" \
	localhost:8080/v1/api/ping-auth-server

signup:
	curl -H "Content-Type: application/json" -X POST http://localhost:8080/v1/api/signup \
	-d "{\"username\":\"quan\", \"password\":\"0\", \"email\":\"quan0@gmail.com\"}"

logout:
	curl -H "Authorization: Bearer ${token}" -X POST http://localhost:8080/v1/api/logout

delete: 
	curl -s -H "Content-Type: application/json" "Authorization: Bearer ${token}" DELETE http://localhost:8080/v1/api/delete \
	-d "{\"user_id\":\"5\"}"