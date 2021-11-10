export PATH := $(PWD)/bin:$(PWD):$(PATH)
export GOBIN := $(PWD)/bin


tools-install:
	go install github.com/spf13/cobra/cobra
	go install go.mozilla.org/sops/v3/cmd/sops

cluster-init:
	@kind create cluster --name boilerplate --wait 40s

cluster-rm:
	@kind delete cluster --name boilerplate

cluster-reset: cluster-rm cluster-init
	@echo "Successfully Completed Cluster Reset"

local-debug:
	@skaffold debug

debug: local-debug

test:
	go test ./...
