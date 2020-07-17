#!/usr/bin/env bash

baseimage="alpine"
dockerarch=("amd64" "arm32v6" "arm64v8" "ppc64le" "s390x")
goarch=("amd64" "arm" "arm64" "ppc64le" "s390x")

for i in "${!dockerarch[@]}"; do
    fromimage="${dockerarch[$i]}/${baseimage}"
    make -B DFROM=${fromimage} GOARCH=${goarch[$i]} all
done

rm -rf ~/.docker/manifests/docker.io_staranto_simple-http-counter-latest
rm -rf ~/.docker/manifests/docker.io_staranto_simple-http-pounder-latest
make -B manifest
