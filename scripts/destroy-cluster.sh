#!/usr/bin/env bash

set -eu

cd "$(dirname "$0")"

source common.sh

k3d delete --name "${CLUSTER_NAME}"
