#!/usr/bin/env bash

set -eu

cd "$(dirname "$0")"

source common.sh

export KUBECONFIG="$(k3d get-kubeconfig --name "${CLUSTER_NAME}")"

k3d import-images "cloud-native-fizzbuzz/mod-3:28a4f52" --name "${CLUSTER_NAME}"
