#!/bin/bash
psql -U postgres -d testdb << "EOSQL"
insert into users (email, password) values ('user0001@example.com', '123451');
insert into users (email, password) values ('user0002@example.com', '123452');
insert into users (email, password) values ('user0003@example.com', '123453');
insert into users (email, password) values ('user0004@example.com', '123453');
insert into users (email, password) values ('user0005@example.com', '123455');
EOSQL
