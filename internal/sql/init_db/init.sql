CREATE DATABASE greenlight;
\c greenlight;

CREATE ROLE greenlight WITH LOGIN PASSWORD 'password';

CREATE EXTENSION IF NOT EXISTS citext;
