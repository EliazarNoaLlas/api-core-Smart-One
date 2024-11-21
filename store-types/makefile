documentation:
	swag init --pd && curl -X POST http://192.168.71.200:8080/api/convert -d @docs/swagger.json --header 'Content-Type: application/json' > docs/swagger3.json

create_store_types_mocks:
	mockery --dir=domain --name=StoreTypeRepository --filename=store_types_repository_mock.go --output=domain/mocks --outpkg=store_types
	mockery --dir=domain --name=StoreTypeUseCase --filename=store_types_usecase_mock.go --output=domain/mocks --outpkg=store_types

PROJECT_PATH = ../../../

deploy-local-store-types:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o app .
	docker build -t localhost:32000/store-types.core.smartone:1.0.0 -f "$(PROJECT_PATH)Dockerfile" .
	docker push localhost:32000/store-types.core.smartone:1.0.0
	rm -rf app