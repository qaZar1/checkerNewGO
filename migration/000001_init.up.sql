CREATE SCHEMA IF NOT EXISTS main;

CREATE TABLE IF NOT EXISTS main.users (
    chat_id     BIGINT NOT NULL UNIQUE PRIMARY KEY,
    username    VARCHAR(32) NOT NULL,
    name        VARCHAR(64)
);

CREATE TABLE IF NOT EXISTS main.versions (
    version         TEXT NOT NULL UNIQUE,
    description     TEXT NOT NULL,
    description_ru  TEXT,
    release_date    TEXT
);