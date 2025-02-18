# PgBouncer Exporter [![Build Status](https://travis-ci.org/Meroje/pgbouncer_exporter.svg)][travis]

[![CircleCI](https://circleci.com/gh/Meroje/pgbouncer_exporter/tree/master.svg?style=shield)][circleci]
[![Docker Repository on Quay](https://quay.io/repository/meroje/pgbouncer-exporter/status)][quay]
[![Docker Pulls](https://img.shields.io/docker/pulls/meroje/pgbouncer-exporter.svg?maxAge=604800)][hub]

Prometheus exporter for PgBouncer.
Exports metrics at `9127/metrics`

## Building and running

    make
    ./pgbouncer_exporter <flags>

To see all available configuration flags:

    ./pgbouncer_exporter -h

## Run with Docker

```
docker run meroje/pgbouncer-exporter --pgBouncer.connectionString=postgres://postgres:@pgbouncer:6432/pgbouncer?sslmode=disable
```

## Metrics

|PgBouncer column|Prometheus Metric|Description|
|----------------|-----------------|-----------|
stats_total_query_count | pgbouncer_stats_queries_pooled_total | Total number of SQL queries pooled
stats.total_query_time | pgbouncer_stats_queries_duration_seconds | Total number of seconds spent by pgbouncer when actively connected to PostgreSQL, executing queries
stats.total_received | pgbouncer_stats_received_bytes_total | Total volume in bytes of network traffic received by pgbouncer, shown as bytes
stats.total_requests | pgbouncer_stats_queries_total | Total number of SQL requests pooled by pgbouncer, shown as requests
stats.total_sent | pgbouncer_stats_sent_bytes_total | Total volume in bytes of network traffic sent by pgbouncer, shown as bytes
stats.total_wait_time | pgbouncer_stats_client_wait_seconds | Time spent by clients waiting for a server in seconds
stats.total_xact_count | pgbouncer_stats_sql_transactions_pooled_total | Total number of SQL transactions pooled
stats.total_xact_time | pgbouncer_stats_server_in_transaction_seconds | Total number of seconds spent by pgbouncer when connected to PostgreSQL in a transaction, either idle in transaction or executing queries
pools.cl_active | pgbouncer_pools_client_active_connections | Client connections linked to server connection and able to process queries, shown as connection
pools.cl_waiting | pgbouncer_pools_client_waiting_connections | Client connections waiting on a server connection, shown as connection
pools.sv_active | pgbouncer_pools_server_active_connections | Server connections linked to a client connection, shown as connection
pools.sv_idle | pgbouncer_pools_server_idle_connections | Server connections idle and ready for a client query, shown as connection
pools.sv_used | pgbouncer_pools_server_used_connections | Server connections idle more than server_check_delay, needing server_check_query, shown as connection
pools.sv_tested | pgbouncer_pools_server_testing_connections | Server connections currently running either server_reset_query or server_check_query, shown as connection
pools.sv_login | pgbouncer_pools_server_login_connections | Server connections currently in the process of logging in, shown as connection
pools.maxwait | pgbouncer_pools_client_maxwait_seconds | Age of oldest unserved client connection, shown as second
databases.pool_size | Maximum number of pool backend connections
databases.reserve_pool | Maximum amount that the pool size can be exceeded temporarily
databases.max_connections | Maximum number of client connections allowed
databases.current_connections | Current number of client connections


[travis]: https://travis-ci.org/Meroje/pgbouncer_exporter
[circleci]: https://circleci.com/gh/Meroje/pgbouncer_exporter
[quay]: https://quay.io/repository/meroje/pgbouncer-exporter
[hub]: https://hub.docker.com/r/meroje/pgbouncer-exporter/
