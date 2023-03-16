#!/usr/bin/env bash

set -o errexit
set -o nounset
set -o pipefail

BINARY="db-init"
OUTPUT="_output"
BIN_PATH="bin"

if [ -f ${OUTPUT}/${BIN_PATH}/${BINARY} ] ; then
  rm -rf ${OUTPUT}/${BIN_PATH}
fi