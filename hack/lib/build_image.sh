#!/usr/bin/env bash

set -o errexit
set -o nounset
set -o pipefail

DEFAULT_TAG=${DEFAULT_TAG:-latest}
#INPUT=[0]
#TAR=${TAG:-${INPUT}}
# support other container tools. e.g. podman
CONTAINER_CLI=${CONTAINER_CLI:-docker}
CONTAINER_BUILDER=${CONTAINER_BUILDER:-build}

# run the service
# judge where the certs is exist.
${CONTAINER_CLI} "${CONTAINER_BUILDER}" \
  -f build/Dockerfile \
  -t db-init:"${DEFAULT_TAG}" .
