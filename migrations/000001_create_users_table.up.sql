CREATE SCHEMA IF NOT EXISTS content;

CREATE TABLE IF NOT EXISTS content.users(
    id serial PRIMARY KEY,
    username VARCHAR (64) UNIQUE NOT NULL,
    password VARCHAR (255) NOT NULL
);