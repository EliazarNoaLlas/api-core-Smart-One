documentation:
	swag init --pd && curl -X POST http://192.168.71.200:8080/api/convert -d @docs/swagger.json --header 'Content-Type: application/json' > docs/swagger3.json

documentation-win:
	swag init --pd
	powershell -Command "Invoke-WebRequest -Method Post -Uri 'http://192.168.71.200:8080/api/convert' -InFile 'docs\swagger.json' -ContentType 'application/json' -OutFile 'docs\swagger3.json'"

create_users_mocks:
	mockery --dir=domain --name=UserRepository --filename=users_repository_mock.go --output=domain/mocks --outpkg=users
	mockery --dir=domain --name=UserUseCase --filename=users_usecase_mock.go --output=domain/mocks --outpkg=users

PROJECT_PATH = ../../../

deploy-local-users:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o app .
	docker build -t localhost:32000/users.core.smartone:1.0.0 -f "$(PROJECT_PATH)Dockerfile" .
	docker push localhost:32000/users.core.smartone:1.0.0
	rm -rf app
