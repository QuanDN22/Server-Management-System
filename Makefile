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
token := eyJhbGciOiJFZERTQSIsInR5cCI6IkpXVCJ9.eyJhdWQiOiJhcGkiLCJleHAiOjE3MTYxOTg1MTIsImlhdCI6MTcxNjE5NTUxMiwiaXNzIjoiaHR0cDovL2xvY2FsaG9zdDo1MDAxIiwibmJmIjoxNzE2MTk1NTEyLCJyb2xlcyI6InVzZXIiLCJ1c2VyIjoicXVhbjIifQ.9XpJGhGferHWPeVq7TD6XhfaxtzSkhMdbEyVrkf3jOG_2HjmP7pAWfupoeHU5JlDm7XtHX2832XbgRIK0tjOAA
# token := eyJhbGciOiJFZERTQSIsInR5cCI6IkpXVCJ9.eyJhdWQiOiJhcGkiLCJleHAiOjE3MTYxOTgwNzgsImlhdCI6MTcxNjE5NTA3OCwiaXNzIjoiaHR0cDovL2xvY2FsaG9zdDo1MDAxIiwibmJmIjoxNzE2MTk1MDc4LCJyb2xlcyI6ImFkbWluIiwidXNlciI6ImFkbWluIn0.55IockvUHsxnIW1IXc18OGwCarwwTGWOaw2pmlrq_Wqr_zZ6Jc3kBqE1lWOELcYI7qzKXojAkEGvd8Bth7ghCA

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