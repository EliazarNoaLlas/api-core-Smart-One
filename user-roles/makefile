documentation:
	swag init --pd && curl -X POST http://192.168.71.200:8080/api/convert -d @docs/swagger.json --header 'Content-Type: application/json' > docs/swagger3.json

create_user_roles_mocks:
	mockery --dir=domain --name=UserRoleRepository --filename=user_roles_repository_mock.go --output=domain/mocks --outpkg=user_roles
	mockery --dir=domain --name=UserRoleUseCase --filename=user_roles_usecase_mock.go --output=domain/mocks --outpkg=user_roles

PROJECT_PATH = ../../../

deploy-local-user-roles:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o app .
	docker build -t localhost:32000/user-roles.core.smartone:1.0.0 -f "$(PROJECT_PATH)Dockerfile" .
	docker push localhost:32000/user-roles.core.smartone:1.0.0
	rm -rf app
