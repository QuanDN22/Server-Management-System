gen:	
	protoc -I ./proto \
   	--go_out ./proto --go_opt paths=source_relative \
   	--go-grpc_out ./proto --go-grpc_opt paths=source_relative \
	--grpc-gateway_out ./proto --grpc-gateway_opt paths=source_relative \
	--openapiv2_out ./static/openapiv2 --openapiv2_opt use_go_templates=true \
   	./proto/auth/auth.proto ./proto/management-system/server.proto

gen-key:
	openssl genpkey -algorithm ED25519 -outform pem -out auth.ed
	openssl pkey -in auth.ed -pubout > auth.ed.pub

# run server:
run-auth-server:
	go run ./cmd/auth/main.go

run-grpc-gateway:
	go run .\cmd\grpc-gateway\main.go auth.ed.pub 

run-management-system:
	go run .\cmd\management-system\main.go auth.ed.pub

# api
# token := eyJhbGciOiJFZERTQSIsInR5cCI6IkpXVCJ9.eyJhdWQiOiJhcGkiLCJleHAiOjE3MTYxOTg1MTIsImlhdCI6MTcxNjE5NTUxMiwiaXNzIjoiaHR0cDovL2xvY2FsaG9zdDo1MDAxIiwibmJmIjoxNzE2MTk1NTEyLCJyb2xlcyI6InVzZXIiLCJ1c2VyIjoicXVhbjIifQ.9XpJGhGferHWPeVq7TD6XhfaxtzSkhMdbEyVrkf3jOG_2HjmP7pAWfupoeHU5JlDm7XtHX2832XbgRIK0tjOAA
token := eyJhbGciOiJFZERTQSIsInR5cCI6IkpXVCJ9.eyJhdWQiOiJhcGkiLCJleHAiOjE3MTYyMDQ5MzksImlhdCI6MTcxNjIwMTkzOSwiaXNzIjoiaHR0cDovL2xvY2FsaG9zdDo1MDAxIiwibmJmIjoxNzE2MjAxOTM5LCJyb2xlcyI6ImFkbWluIiwidXNlciI6ImFkbWluIn0.FlMQFy0uWZG7Dhw9_AHufbY874qfu8hlp9pM_eFccSE5ztQGSjR5SDFw-a0aHb49awxbwl3HAUUOYMHJ_DeSBg
# api of auth server
login: 
	curl -H "Content-Type: application/json" -X POST http://localhost:8080/v1/api/auth/login \
	-d "{"username":"admin", "password":"pass"}"

# -d "{\"username\":\"admin\", \"password\":\"pass\"}"

login-user:
	curl -H "Content-Type: application/json" -X POST http://localhost:8080/v1/api/auth/login \
	-d "{\"username\":\"quan2\", \"password\":\"2\"}"

ping-auth-server: 
	curl -s -H "Authorization: Bearer ${token}" \
	localhost:8080/v1/api/auth/ping

signup:
	curl -H "Content-Type: application/json" -X POST http://localhost:8080/v1/api/auth/signup \
	-d "{\"username\":\"quan\", \"password\":\"0\", \"email\":\"quan0@gmail.com\"}"

logout:
	curl -H "Authorization: Bearer ${token}" -X POST http://localhost:8080/v1/api/logout

delete: 
	curl -s \
  	-H "Authorization: Bearer ${token}" \
  	-H "Content-Type: application/json" \
  	DELETE http://localhost:8080/v1/api/delete-user \
  	-d "{\"userId\":\"2\"}"

# curl -s -H "Content-Type: application/json" "Authorization: Bearer ${token}" DELETE http://localhost:8080/v1/api/delete \
	# -d "{\"user_id\":\"1\"}"