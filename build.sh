#!/usr/bin/env bash

set -ex

PROJECT_DIR=$(dirname $0)
BUILD_DIR="${PROJECT_DIR}/out"

buildBinaries(){
  [[ ! -d "${BUILD_DIR}" ]] && mkdir "${BUILD_DIR}"

  go build -o "${BUILD_DIR}/crypt-linux-amd64" crypt.go

  GOOS=windows go build -o "${BUILD_DIR}/crypt-windows-amd64" crypt.go

  GOOS=darwin go build -o "${BUILD_DIR}/crypt-darwin-amd64" crypt.go
}

##### main
buildBinaries
