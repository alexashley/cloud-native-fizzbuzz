#!/usr/bin/env bash

set -eu

cd "$(dirname "$0")"

source common.sh

export KUBECONFIG="$(k3d get-kubeconfig --name "${CLUSTER_NAME}")"

helm template ../k8s/cn-fizzbuzz -n cn-fizzbuzz | kubectl apply -f -