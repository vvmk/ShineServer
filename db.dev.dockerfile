FROM postgres:11
COPY dev_schema.sql /docker-entrypoint-initdb.d/
