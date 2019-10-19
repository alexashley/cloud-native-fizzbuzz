#!/usr/bin/env bash

set -eu

cd "$(dirname "$0")"

source common.sh

k3d create --api-port 6550 --publish 8081:8081 --workers 2 --name "${CLUSTER_NAME}"
