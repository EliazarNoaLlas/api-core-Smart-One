UNAME_S := $(shell uname -s)

PROJECT_PATH = ../../../
#PROJECT_PATH = ../../
REGISTRY_URL = "localhost:32000"
DB_HOST = "192.168.71.200"
NAMESPACE = "smartone-local"

ifeq ($(UNAME_S),Darwin)
    REGISTRY_URL := "192.168.64.2:32000"
endif

ifdef REGISTRY_URL_DEV
    REGISTRY_URL := $(REGISTRY_URL_DEV)
endif

deploy-micro:
	@export DB_HOST="$(DB_HOST)" && \
	cd "$(DIR_MICRO)" && CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o app . && \
	docker build -t "$(REGISTRY_URL)/$(IMAGE)" -f "$(PROJECT_PATH)Dockerfile" . && \
	docker push "$(REGISTRY_URL)/$(IMAGE)" && \
	rm -rf app && \
	export REGISTRY_URL="$(REGISTRY_URL)" && export NAMESPACE="$(NAMESPACE)" && \
	envsubst < cicd/k8s/local/deployment-api.yaml | kubectl --kubeconfig="$(HOME)/.kube/config-local" apply -f - && \
	envsubst < cicd/k8s/local/service-api.yaml | kubectl --kubeconfig="$(HOME)/.kube/config-local" apply -f - && \
	kubectl --kubeconfig="$(HOME)/.kube/config-local" delete pods -l app="$(LABEL)" -n "$(NAMESPACE)"

deploy-core-main:
	@export DIR_MICRO="." && export IMAGE="core.smartone:1.0.0" && export LABEL="core-smartone" && $(MAKE) deploy-micro

deploy-core: deploy-local-document-types deploy-local-economic-activities deploy-local-merchant-economic-activities  deploy-local-merchants deploy-local-modules deploy-local-permissions deploy-local-policies deploy-local-policy-permissions deploy-local-role-policies deploy-local-roles deploy-local-stores deploy-local-user-roles deploy-local-user-types deploy-local-users deploy-local-views

deploy-local-document-types:
	@export DIR_MICRO="document-types" && export IMAGE="document-types.core.smartone:1.0.0" && export LABEL="document-types-core-smartone" && $(MAKE) deploy-micro

deploy-local-economic-activities:
	@export DIR_MICRO="economic-activities" && export IMAGE="economic-activities.core.smartone:1.0.0" && export LABEL="economic-activities-core-smartone" && $(MAKE) deploy-micro

deploy-local-merchant-economic-activities:
	@export DIR_MICRO="merchant-economic-activities" && export IMAGE="merchant-economic-activities.core.smartone:1.0.0" && export LABEL="merchant-economic-activities-core-smartone" && $(MAKE) deploy-micro

deploy-local-merchants:
	@export DIR_MICRO="merchants" && export IMAGE="merchants.core.smartone:1.0.0" && export LABEL="merchants-core-smartone" && $(MAKE) deploy-micro

deploy-local-modules:
	@export DIR_MICRO="modules" && export IMAGE="modules.core.smartone:1.0.0" && export LABEL="modules-core-smartone" && $(MAKE) deploy-micro

deploy-local-permissions:
	@export DIR_MICRO="permissions" && export IMAGE="permissions.core.smartone:1.0.0" && export LABEL="permissions-core-smartone" && $(MAKE) deploy-micro

deploy-local-policies:
	@export DIR_MICRO="policies" && export IMAGE="policies.core.smartone:1.0.0" && export LABEL="policies-core-smartone" && $(MAKE) deploy-micro

deploy-local-policy-permissions:
	@export DIR_MICRO="policy-permissions" && export IMAGE="policy-permissions.core.smartone:1.0.0" && export LABEL="policy-permissions-core-smartone" && $(MAKE) deploy-micro

deploy-local-role-policies:
	@export DIR_MICRO="role-policies" && export IMAGE="role-policies.core.smartone:1.0.0" && export LABEL="role-policies-core-smartone" && $(MAKE) deploy-micro

deploy-local-roles:
	@export DIR_MICRO="roles" && export IMAGE="roles.core.smartone:1.0.0" && export LABEL="roles-core-smartone" && $(MAKE) deploy-micro

deploy-local-store-types:
	@export DIR_MICRO="store-types" && export IMAGE="store-types.core.smartone:1.0.0" && export LABEL="store-types-core-smartone" && $(MAKE) deploy-micro

deploy-local-stores:
	@export DIR_MICRO="stores" && export IMAGE="stores.core.smartone:1.0.0" && export LABEL="stores-core-smartone" && $(MAKE) deploy-micro

deploy-local-user-roles:
	@export DIR_MICRO="user-roles" && export IMAGE="user-roles.core.smartone:1.0.0" && export LABEL="user-roles-core-smartone" && $(MAKE) deploy-micro

deploy-local-user-types:
	@export DIR_MICRO="user-types" && export IMAGE="user-types.core.smartone:1.0.0" && export LABEL="user-types-core-smartone" && $(MAKE) deploy-micro

deploy-local-users:
	@export DIR_MICRO="users" && export IMAGE="users.core.smartone:1.0.0" && export LABEL="users-core-smartone" && $(MAKE) deploy-micro

deploy-local-views:
	@export DIR_MICRO="views" && export IMAGE="views.core.smartone:1.0.0" && export LABEL="views-core-smartone" && $(MAKE) deploy-micro
