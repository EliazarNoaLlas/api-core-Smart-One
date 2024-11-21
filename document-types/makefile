documentation:
	swag init --pd && curl -X POST http://192.168.71.200:8080/api/convert -d @docs/swagger.json --header 'Content-Type: application/json' > docs/swagger3.json

create_document_types_mocks:
	mockery --dir=domain --name=DocumentTypeRepository --filename=document_types_repository_mock.go --output=domain/mocks --outpkg=document_types
	mockery --dir=domain --name=DocumentTypeUseCase --filename=document_types_usecase_mock.go --output=domain/mocks --outpkg=document_types

PROJECT_PATH = ../../../

deploy-local-document-types:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o app .
	docker build -t localhost:32000/document-types.core.smartone:1.0.0 -f "$(PROJECT_PATH)Dockerfile" .
	docker push localhost:32000/document-types.core.smartone:1.0.0
	rm -rf app