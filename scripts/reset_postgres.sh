#!/bin/bash

pg_ctl -D ./data/toc-pd-captain -l ./data/toc-pd-captain/logfile stop

rm -rf ./data/toc-pd-captain
mkdir -p ./data/toc-pd-captain

initdb ./data/toc-pd-captain

gsed -i "$ a host    all    all    0.0.0.0/0    trust" ./data/toc-pd-captain/pg_hba.conf
gsed -i "$ a listen_addresses = '*'" ./data/toc-pd-captain/postgresql.conf

pg_ctl -D ./data/toc-pd-captain -l ./data/toc-pd-captain/logfile start

echo "\du
CREATE ROLE postgres WITH LOGIN PASSWORD 'asdf0000';
ALTER USER postgres WITH SUPERUSER;
\du" > sql_script

psql postgres -f sql_script
rm -rf sql_script

# pg_ctl -D ./data/toc-pd-captain -l ./data/toc-pd-captain/logfile stop
