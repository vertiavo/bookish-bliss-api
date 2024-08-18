# Build binary for Docker image
local_resource(
    'bookish-bliss-api-compile',
    'go build -o bookish-bliss-api ./cmd/bookishbliss/main.go',
    deps=['./cmd', './internal' , './pkg']
)

# Docker build for Go application
docker_build('bookish-bliss-api:latest', '.')
docker_build('bookish-bliss-api-migrator:latest', '.', target='migrator')

# Kubernetes resources using Kustomize for local
k8s_yaml(kustomize('deploy/dev'))

# Forward port for local access
k8s_resource('bookish-bliss-api', port_forwards=3000)
k8s_resource('postgres', port_forwards=5432)
