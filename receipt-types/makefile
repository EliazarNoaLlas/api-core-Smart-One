documentation:
	swag init --pd && curl -X POST http://192.168.71.200:8080/api/convert -d @docs/swagger.json --header 'Content-Type: application/json' > docs/swagger3.json

create_receipt_types_mocks:
	mockery --dir=domain --name=ReceiptTypesRepository --filename=receipt_types_repository_mock.go --output=domain/mocks --outpkg=receipt_types
	mockery --dir=domain --name=ReceiptTypesUseCase --filename=receipt_types_usecase_mock.go --output=domain/mocks --outpkg=receipt_types

PROJECT_PATH = ../../../

deploy-local-receipt-types:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o app .
	docker build -t localhost:32000/receipt-types.logistic.smartone:1.0.0 -f "$(PROJECT_PATH)Dockerfile" .
	docker push localhost:32000/receipt-types.logistic.smartone:1.0.0
	rm -rf app