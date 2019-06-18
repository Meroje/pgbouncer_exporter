ARG ARCH="amd64"
ARG OS="linux"
FROM quay.io/prometheus/busybox-${OS}-${ARCH}:latest

ARG ARCH="amd64"
ARG OS="linux"
COPY .build/${OS}-${ARCH}/pgbouncer_exporter /bin/pgbouncer_exporter

ENTRYPOINT ["/bin/pgbouncer_exporter"]
EXPOSE     9127
