UNAME_S := $(shell uname -s)

PROJECT_PATH = ./
REGISTRY_URL = "localhost:32000"
DB_HOST = "192.168.71.200"
NAMESPACE = "smartone-local"

MIGRATION_DIR = "migrations"
MIGRATION_TENANT_DIR = "migrations-tenant"
DB_TENANT_NAME="db_tenant"
DB_NAME="db_smartone"
DB_PORT="3901"
DB_USER="root"
DB_PASSWORD="U7fiLttFrIrdvVkk"

ifeq ($(UNAME_S),Darwin)
    REGISTRY_URL := "192.168.64.2:32000"
endif

ifdef REGISTRY_URL_DEV
    REGISTRY_URL := $(REGISTRY_URL_DEV)
endif

ifdef REGISTRY_URL_PROD
    REGISTRY_URL := $(REGISTRY_URL_PROD)
endif

deploy-micro:
	cd "$(DIR_MICRO)" && CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o app . && \
	docker build -t "$(REGISTRY_URL)/$(IMAGE):$(VERSION)" -f "$(PROJECT_PATH)Dockerfile" . && \
	docker push "$(REGISTRY_URL)/$(IMAGE):$(VERSION)" && \
	rm -rf app

deploy-core-main:
	@export DIR_MICRO="." && export IMAGE="core.smartone" && $(MAKE) deploy-micro

deploy-core-main-local:
	@export DIR_MICRO="." && export IMAGE="core.smartone" && export VERSION="v1.0.0" && $(MAKE) deploy-micro && \
	export REGISTRY_URL="$(REGISTRY_URL)" && export NAMESPACE="$(NAMESPACE)" && export LABEL="core-smartone" && \
	envsubst < cicd/k8s/local/deployment-api.yaml | kubectl --kubeconfig="$(HOME)/.kube/config-local" apply -f - && \
	envsubst < cicd/k8s/local/service-api.yaml | kubectl --kubeconfig="$(HOME)/.kube/config-local" apply -f - && \
	kubectl --kubeconfig="$(HOME)/.kube/config-local" delete pods -l app="$(LABEL)" -n "$(NAMESPACE)"

migrate-up:
	goose -dir $(MIGRATION_DIR) mysql "$(DB_USER):$(DB_PASSWORD)@tcp($(DB_HOST):$(DB_PORT))/$(DB_NAME)?parseTime=true" up

migrate-down:
	goose -dir $(MIGRATION_DIR) mysql "$(DB_USER):$(DB_PASSWORD)@tcp($(DB_HOST):$(DB_PORT))/$(DB_NAME)?parseTime=true" down

migrate-tenant-up:
	goose -dir $(MIGRATION_TENANT_DIR) mysql "$(DB_USER):$(DB_PASSWORD)@tcp($(DB_HOST):$(DB_PORT))/$(DB_TENANT_NAME)?parseTime=true" up

migrate-tenant-down:
	goose -dir $(MIGRATION_TENANT_DIR) mysql "$(DB_USER):$(DB_PASSWORD)@tcp($(DB_HOST):$(DB_PORT))/$(DB_TENANT_NAME)?parseTime=true" down
