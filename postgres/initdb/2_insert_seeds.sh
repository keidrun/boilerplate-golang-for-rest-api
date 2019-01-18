#!/bin/bash

psql -U postgres -d testdb << EOSQL
  insert into users (email, password, name, age, gender) values (
    'user0001@example.com', '$1a$10$AjaEt2x2CuVPECskm0UJXujQIS/zCh/s1BaE0t2MNJ52rrKbfl17e', 'Tom', 20, 'male'
  );
  insert into users (email, password, name, age, gender) values (
    'user0002@example.com', '$2a$10$AjaEt2x2CuVPECskm0UJXujQIS/zCh/s1BaE0t2MNJ52rrKbfl17e', 'John', 31, 'male'
  );
  insert into users (email, password, name, age, gender) values (
    'user0003@example.com', '$3a$10$AjaEt2x2CuVPECskm0UJXujQIS/zCh/s1BaE0t2MNJ52rrKbfl17e', 'Sophia', 18, 'female'
  );
  insert into users (email, password, name, age, gender) values (
    'user0004@example.com', '$4a$10$AjaEt2x2CuVPECskm0UJXujQIS/zCh/s1BaE0t2MNJ52rrKbfl17e', 'Olivia', 24, 'female'
  );
  insert into users (email, password, name) values (
    'user0005@example.com', '$5a$10$AjaEt2x2CuVPECskm0UJXujQIS/zCh/s1BaE0t2MNJ52rrKbfl17e', 'Steve'
  );
EOSQL
