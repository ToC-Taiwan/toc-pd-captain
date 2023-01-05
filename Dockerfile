# build-stage
FROM golang:1.19.4-bullseye as build-stage
USER root

ENV TZ=Asia/Taipei

WORKDIR /
RUN mkdir build_space
WORKDIR /build_space
COPY . .
RUN make build

# production-stage
FROM debian:bullseye as production-stage
USER root

ENV TZ=Asia/Taipei

WORKDIR /
RUN apt update -y && \
    apt install -y tzdata && \
    apt autoremove -y && \
    apt clean && \
    mkdir toc-pd-captain && \
    mkdir toc-pd-captain/migrations && \
    mkdir toc-pd-captain/configs && \
    mkdir toc-pd-captain/logs && \
    mkdir toc-pd-captain/scripts && \
    rm -rf /var/lib/apt/lists/*

WORKDIR /toc-pd-captain

COPY --from=build-stage /build_space/toc-pd-captain ./toc-pd-captain
COPY --from=build-stage /build_space/migrations ./migrations/
COPY --from=build-stage /build_space/scripts/docker-entrypoint.sh ./scripts/docker-entrypoint.sh

ENTRYPOINT ["/toc-pd-captain/scripts/docker-entrypoint.sh"]
