documentation:
	swag init --pd && curl -X POST http://192.168.71.200:8080/api/convert -d @docs/swagger.json --header 'Content-Type: application/json' > docs/swagger3.json

create_stores_mocks:
	mockery --dir=domain --name=StoreRepository --filename=stores_repository_mock.go --output=domain/mocks --outpkg=stores
	mockery --dir=domain --name=StoreUseCase --filename=stores_usecase_mock.go --output=domain/mocks --outpkg=stores

PROJECT_PATH = ../../../

deploy-local-stores:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o app .
	docker build -t localhost:32000/stores.core.smartone:1.0.0 -f "$(PROJECT_PATH)Dockerfile" .
	docker push localhost:32000/stores.core.smartone:1.0.0
	rm -rf app
