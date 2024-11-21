documentation:
	swag init --pd && curl -X POST http://192.168.71.200:8080/api/convert -d @docs/swagger.json --header 'Content-Type: application/json' > docs/swagger3.json

create_economic_activities_mocks:
	mockery --dir=domain --name=EconomicActivityRepository --filename=economic_activities_repository_mock.go --output=domain/mocks --outpkg=economic_activities
	mockery --dir=domain --name=EconomicActivityUseCase --filename=economic_activities_usecase_mock.go --output=domain/mocks --outpkg=economic_activities

PROJECT_PATH = ../../../

deploy-local-economic-activities:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o app .
	docker build -t localhost:32000/economic-activities.core.smartone:1.0.0 -f "$(PROJECT_PATH)Dockerfile" .
	docker push localhost:32000/economic-activities.core.smartone:1.0.0
	rm -rf app