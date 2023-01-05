#!/bin/bash

APP="toc-pd-captain"
BUILD_DIR="../"

rm -rf ${BUILD_DIR}/"${APP}"

go build -o ${BUILD_DIR}/"${APP}" ../cmd/app
