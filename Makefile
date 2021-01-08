# Copyright SecureKey Technologies Inc.
#
# SPDX-License-Identifier: Apache-2.0

VC_REST_PATH=cmd/vc-rest
DID_REST_PATH=cmd/did-rest

# Namespace for the agent images
DOCKER_OUTPUT_NS   ?= ghcr.io
VC_REST_IMAGE_NAME   ?= trustbloc/vc-server
DID_REST_IMAGE_NAME   ?= trustbloc/did-resolver
DID_ELEMENT_SIDETREE_REQUEST_URL ?= https://element-did.com/api/v1/sidetree/requests

# OpenAPI spec
OPENAPI_DOCKER_IMG=quay.io/goswagger/swagger
OPENAPI_SPEC_PATH=build/rest/openapi/spec
OPENAPI_DOCKER_IMG_VERSION=v0.23.0

# Tool commands (overridable)
ALPINE_VER ?= 3.12
GO_VER ?= 1.15

.PHONY: all
all: checks unit-test bdd-test

.PHONY: checks
checks: license lint generate-openapi-spec

.PHONY: lint
lint:
	@scripts/check_lint.sh

.PHONY: license
license:
	@scripts/check_license.sh

.PHONY: generate-test-config
generate-test-config:
	@scripts/generate_test_config.sh

.PHONY: vc-rest
vc-rest:
	@echo "Building vc-rest"
	@mkdir -p ./.build/bin
	@cd ${VC_REST_PATH} && go build -o ../../.build/bin/vc-rest main.go

.PHONY: vc-server-docker
vc-server-docker:
	@echo "Building vc rest docker image"
	@docker build -f ./images/vc-rest/Dockerfile --no-cache -t $(DOCKER_OUTPUT_NS)/$(VC_REST_IMAGE_NAME):latest \
	--build-arg GO_VER=$(GO_VER) \
	--build-arg ALPINE_VER=$(ALPINE_VER) .

.PHONY: did-rest
did-rest:
	@echo "Building did-rest"
	@mkdir -p ./.build/bin
	@cd ${DID_REST_PATH} && go build -o ../../.build/bin/did-rest main.go

.PHONY: did-resolver-docker
did-resolver-docker:
	@echo "Building did rest docker image"
	@docker build -f ./images/did-rest/Dockerfile --no-cache -t $(DOCKER_OUTPUT_NS)/$(DID_REST_IMAGE_NAME):latest \
	--build-arg GO_VER=$(GO_VER) \
	--build-arg ALPINE_VER=$(ALPINE_VER) .

.PHONY: bdd-test
bdd-test: clean vc-server-docker did-resolver-docker generate-test-keys generate-test-config
	@scripts/check_integration.sh

.PHONY: bdd-interop-test
bdd-interop-test:clean vc-server-docker did-resolver-docker generate-test-keys generate-test-config
	@scripts/check_integration_interop.sh

unit-test:
	@scripts/check_unit.sh

.PHONY: generate-test-keys
generate-test-keys: clean
	@mkdir -p -p test/bdd/fixtures/keys/tls
	@docker run -i --rm \
		-v $(abspath .):/opt/workspace/edge-service \
		--entrypoint "/opt/workspace/edge-service/scripts/generate_test_keys.sh" \
		frapsoft/openssl

create-element-did: clean
	@mkdir -p .build
	@cp scripts/create-element-did.js .build/
	@REQUEST_URL=$(DID_ELEMENT_SIDETREE_REQUEST_URL) scripts/create_element_did.sh

# this target creates VCs and VPs from other systems for interop tests
prepare-test-verifiables: clean
	@mkdir -p .build
	@cp scripts/prepare-test-verifiables.js .build/
	@scripts/prepare_test_verifiables.sh

.PHONY: generate-openapi-spec
generate-openapi-spec: clean
	@echo "Generating and validating controller API specifications using Open API"
	@mkdir -p build/rest/openapi/spec
	@SPEC_META=$(VC_REST_PATH) SPEC_LOC=${OPENAPI_SPEC_PATH}  \
	DOCKER_IMAGE=$(OPENAPI_DOCKER_IMG) DOCKER_IMAGE_VERSION=$(OPENAPI_DOCKER_IMG_VERSION)  \
	scripts/generate-openapi-spec.sh

.PHONY: generate-openapi-demo-specs
generate-openapi-demo-specs: clean generate-openapi-spec vc-server-docker did-resolver-docker
	@echo "Generate demo agent rest controller API specifications using Open API"
	@SPEC_PATH=${OPENAPI_SPEC_PATH} OPENAPI_DEMO_PATH=test/bdd/fixtures/openapi-demo \
    	DOCKER_IMAGE=$(OPENAPI_DOCKER_IMG) DOCKER_IMAGE_VERSION=$(OPENAPI_DOCKER_IMG_VERSION)  \
    	scripts/generate-openapi-demo-specs.sh

.PHONY: run-openapi-demo
run-openapi-demo: generate-test-keys generate-test-config  generate-openapi-demo-specs
	@echo "Starting demo vc rest containers ..."
	@FIXTURES_PATH=test/bdd/fixtures  \
        scripts/run-openapi-demo.sh

.PHONY: clean
clean: clean-build

.PHONY: clean-build
clean-build:
	@rm -Rf ./.build
	@rm -Rf ./test/bdd/fixtures/keys/tls
	@rm -Rf ./test/bdd/docker-compose.log
	@rm -Rf ./test/bdd/fixtures/wellknown/jws
	@rm -Rf ./test/bdd/fixtures/discovery-server/config
	@rm -Rf ./test/bdd/fixtures/stakeholder-server/config
