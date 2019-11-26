#!/bin/bash

set -o errexit
set -o pipefail

if [ -z "$1" ]; then
    echo Please specify the name of the service.
    echo There should be a Dockerfile in services/\$name.
    exit 1
fi

case $2 in
    intercept)
        template=service_intercept.yaml.template
        ;;
    plain)
        template=service_plain.yaml.template
        ;;
    *)
        echo Please specify intercept or plain for the type of service.
        exit 1
        ;;
esac

set -o nounset

name=$1

if [ \! -f "services/${name}/Dockerfile" ]; then
    echo "Did not find services/${name}/Dockerfile"
    exit 1
fi

sed "s/__NAME__/${name}/g" < "k8s/${template}" >> k8s/20-services.yaml
