documentation:
	swag init --pd && curl -X POST http://192.168.71.200:8080/api/convert -d @docs/swagger.json --header 'Content-Type: application/json' > docs/swagger3.json

create_view_permissions_mocks:
	mockery --dir=domain --name=ViewPermissionsRepository --filename=view_permissions_repository_mock.go --output=domain/mocks --outpkg=view_permissions
	mockery --dir=domain --name=ViewPermissionsUseCase --filename=view_permissions_usecase_mock.go --output=domain/mocks --outpkg=view_permissions

PROJECT_PATH = ../../../

deploy-local-view-permissions:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o app .
	docker build -t localhost:32000/view-permissions.logistic.smartone:1.0.0 -f "$(PROJECT_PATH)Dockerfile" .
	docker push localhost:32000/view-permissions.logistic.smartone:1.0.0
	rm -rf app