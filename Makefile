cluster-init:
	kind create cluster --name boilerplate

cluster-clean:
	kind delete cluster --name boilerplate

set-context:
	kubectl cluster-info --context kind-boilerplate

test:
	go test ./... -short