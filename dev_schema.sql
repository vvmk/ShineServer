CREATE EXTENSION citext;                                              │~
CREATE EXTENSION                                                               │~
CREATE DOMAIN email AS citext                                         │~
CHECK ( value ~ '^[a-zA-Z0-9.!#$%&''*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{│~
0,61}[a-zA-Z0-9])?(?:\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$' );    │~
CREATE DOMAIN                                                                  │~
CREATE TABLE users(                                                   │~
user_id serial primary key,                                                    │~
email email UNIQUE NOT NULL,                                                   │~
tag text NOT NULL,                                                             │~
bio text,                                                                      │~
main text);                                                                    │~
CREATE TABLE                                                                   │~
CREATE TABLE routines(                                                │~
routine_id serial primary key,                                                 │~
title text NOT NULL,                                                           │~
total_duration SMALLINT NOT NULL CHECK (total_duration > 0),                   │~
character text NOT NULL,                                                       │~
creator_id INT NOT NULL REFERENCES users(user_id),                             │~
created timestamp default current_timestamp,                                   │~
popularity INT default 0,                                                      │~
drills jsonb);                                                                 │~
CREATE TABLE                                                                   │~
ALTER TABLE routines ADD COLUMN original_creator_id INT NOT NULL REFER│~
ENCES users(user_id);                                                          │~
ALTER TABLE
