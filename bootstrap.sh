#!/bin/bash

mkdir -p ./cmd/{api,grpc,cli}
touch ./cmd/{api,grpc,cli}/main.go
mkdir -p ./internal/domain/{metric,alert,resource}
mkdir -p ./internal/ports/{repositories,services,notifiers}
mkdir -p ./internal/adapters/{api,cli,repositories,clients,notifiers}
touch ./internal/domain/{metric,alert,resource}/.gitkeep
touch ./internal/ports/{repositories,services,notifiers}/.gitkeep
touch ./internal/adapters/{api,cli,repositories,clients,notifiers}/.gitkeep
mkdir -p ./pkg/common
touch ./pkg/common/.gitkeep
mkdir -p ./{configs,deployments,docs}
touch ./{configs,deployments,docs}/.gitkeep
mkdir -p ./tests/{integration,e2e}
touch ./tests/{integration,e2e}/.gitkeep
mkdir -p ./.github/workflows
