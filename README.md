# go-api-boilerplate

golang boilerplate for apis in k8s

## Install

#### Skaffold

https://skaffold.dev/docs/install/

#### Helm

https://helm.sh/docs/intro/install/

#### Kind

https://kind.sigs.k8s.io/docs/user/quick-start#installation


## Debug

Initialize cluster
* `make cluster-init`

Deploy Services (Could be done through cloud code to enable easy debugging in vscode)
* `make debug`

