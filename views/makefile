documentation:
	swag init --pd && curl -X POST http://192.168.71.200:8080/api/convert -d @docs/swagger.json --header 'Content-Type: application/json' > docs/swagger3.json

create_views_mocks:
	mockery --dir=domain --name=ViewRepository --filename=views_repository_mock.go --output=domain/mocks --outpkg=views
	mockery --dir=domain --name=ViewUseCase --filename=views_usecase_mock.go --output=domain/mocks --outpkg=views

PROJECT_PATH = ../../../

deploy-local-views:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o app .
	docker build -t localhost:32000/views.core.smartone:1.0.0 -f "$(PROJECT_PATH)Dockerfile" .
	docker push localhost:32000/views.core.smartone:1.0.0
	rm -rf app