# Polaris Provider Helm Chart

Helm Chart for deploying Polaris Provider service to Kubernetes.

## Prerequisites

- Kubernetes 1.19+
- Helm 3.0+

## Installation

### Using default values

```bash
helm install polaris-provider ./deploy/helm
```

### Using custom values

```bash
helm install polaris-provider ./deploy/helm -f custom-values.yaml
```

### Using --set flags

```bash
helm install polaris-provider ./deploy/helm \
  --set deployment.replicas=3 \
  --set container.image=my-registry/polaris-go-bin:v1.0 \
  --set polaris.config.global.serverConnector.addresses[0]=192.168.1.5:8091
```

## Upgrade

```bash
helm upgrade polaris-provider ./deploy/helm
```

## Uninstall

```bash
helm uninstall polaris-provider
```

## Configuration

See `values.yaml` for all available configuration options.

### Key Configuration Parameters

- `deployment.replicas`: Number of provider replicas (default: 2)
- `container.image`: Docker image to use (default: polaris-go-bin:latest)
- `container.resources`: CPU and memory limits/requests
- `service.type`: Kubernetes service type (default: ClusterIP)
- `polaris.config`: Polaris service configuration

## Equivalence with kubectl

The original deployment using `kubectl apply -f provider.yaml` can now be replaced with:

```bash
helm install polaris-provider ./deploy/helm
```

Both methods achieve the same result, but Helm provides additional benefits:
- Easy parameter customization
- Version control for deployments
- Rollback capabilities
- Template reusability
