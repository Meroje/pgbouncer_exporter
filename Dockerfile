FROM quay.io/prometheus/busybox:latest

COPY pgbouncer_exporter /bin/pgbouncer_exporter

ENTRYPOINT ["/bin/pgbouncer_exporter"]
EXPOSE     9127
