initialize-cluster:
	kind create cluster --name boilerplate

set-context:
	kubectl cluster-info --context kind-boilerplate

test:
	go test ./... -short