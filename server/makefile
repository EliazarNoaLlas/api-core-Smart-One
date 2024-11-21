documentation:
	swag init --pd --instanceName server  && curl -X POST http://192.168.71.200:8080/api/convert -d @docs/server_swagger.json --header 'Content-Type: application/json' > docs/swagger3.json

documentation-win:
	swag init --pd --instanceName server
	powershell -Command "Invoke-WebRequest -Method Post -Uri 'http://192.168.71.200:8080/api/convert' -InFile 'docs\server_swagger.json' -ContentType 'application/json' -OutFile 'docs\swagger3.json'"

create_server_mocks:
	mockery --dir=domain --name=ServerUseCase --filename=server_usecase_mock.go --output=domain/mocks --outpkg=server

PROJECT_PATH = ../../../

deploy-local-server:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o app .
	docker build -t localhost:32000/server.core.smartone:1.0.0 -f "$(PROJECT_PATH)Dockerfile" .
	docker push localhost:32000/server.core.smartone:1.0.0
	rm -rf app
