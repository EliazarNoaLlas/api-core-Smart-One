documentation:
	swag init --pd && curl -X POST http://192.168.71.200:8080/api/convert -d @docs/swagger.json --header 'Content-Type: application/json' > docs/swagger3.json

create_role_policies_mocks:
	mockery --dir=domain --name=RolePolicyRepository --filename=role_policies_repository_mock.go --output=domain/mocks --outpkg=rolePolicies
	mockery --dir=domain --name=RolePolicyUseCase --filename=role_policies_usecase_mock.go --output=domain/mocks --outpkg=rolePolicies

PROJECT_PATH = ../../../

deploy-local-role-policies:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o app .
	docker build -t localhost:32000/role-policies.core.smartone:1.0.0 -f "$(PROJECT_PATH)Dockerfile" .
	docker push localhost:32000/role-policies.core.smartone:1.0.0
	rm -rf app
