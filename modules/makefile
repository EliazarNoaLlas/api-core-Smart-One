documentation:
	swag init --pd && curl -X POST http://192.168.71.200:8080/api/convert -d @docs/swagger.json --header 'Content-Type: application/json' > docs/swagger3.json

create_modules_mocks:
	mockery --dir=domain --name=ModuleRepository --filename=modules_repository_mock.go --output=domain/mocks --outpkg=modules
	mockery --dir=domain --name=ModuleUseCase --filename=modules_usecase_mock.go --output=domain/mocks --outpkg=modules

PROJECT_PATH = ../../../

deploy-local-modules:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o app .
	docker build -t localhost:32000/modules.core.smartone:1.0.0 -f "$(PROJECT_PATH)Dockerfile" .
	docker push localhost:32000/modules.core.smartone:1.0.0
	rm -rf app