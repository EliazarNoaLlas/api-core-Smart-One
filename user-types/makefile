documentation:
	swag init --pd && curl -X POST http://192.168.71.200:8080/api/convert -d @docs/swagger.json --header 'Content-Type: application/json' > docs/swagger3.json

create_user_types_mocks:
	mockery --dir=domain --name=UserTypeRepository --filename=user_types_repository_mock.go --output=domain/mocks --outpkg=userTypes
	mockery --dir=domain --name=UserTypeUseCase --filename=user_types_usecase_mock.go --output=domain/mocks --outpkg=userTypes

PROJECT_PATH = ../../../

deploy-local-user-types:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o app .
	docker build -t localhost:32000/user-types.core.smartone:1.0.0 -f "$(PROJECT_PATH)Dockerfile" .
	docker push localhost:32000/user-types.core.smartone:1.0.0
	rm -rf app

