# TOC PD CAPTAIN

[![Workflow](https://github.com/ToC-Taiwan/toc-pd-captain/actions/workflows/main.yml/badge.svg)](https://github.com/ToC-Taiwan/toc-pd-captain/actions/workflows/main.yml)
[![Maintained](https://img.shields.io/badge/Maintained-yes-green)](https://github.com/ToC-Taiwan/toc-pd-captain)
[![Go](https://img.shields.io/badge/Go-1.19.4-blue?logo=go&logoColor=blue)](https://golang.org)
[![OS](https://img.shields.io/badge/OS-Linux-orange?logo=linux&logoColor=orange)](https://www.linux.org/)
[![Container](https://img.shields.io/badge/Container-Docker-blue?logo=docker&logoColor=blue)](https://www.docker.com/)

## Config

```sh
cp ./configs/default.config.yml ./configs/config.yml
```

## Env

```sh
cp .env.template .env
```

## VSCode

```sh
rm -rf .vscode
mkdir .vscode
cp .launch.template.json .vscode/launch.json
```

## Build

```sh
make build
```

## Run

```sh
make
```

## Authors

- [**Tim Hsu**](https://github.com/Chindada)
