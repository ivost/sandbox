#!/bin/bash

set -o errexit
set -o pipefail

if [ -z "$AMBASSADOR_LICENSE_KEY" ]; then
    echo "Error: AMBASSADOR_LICENSE_KEY is not set"
    exit 1
fi

set -o nounset

if ! [ -x "$(command -v kubeapply)" ]; then
    echo "Error: kubeapply is not installed"
    exit 1
fi

echo "Docker registry is $DOCKER_REGISTRY"
kubeapply -f k8s
