CREATE DATABASE crawler;
GRANT ALL PRIVILEGES ON DATABASE crawler TO postgres;

\connect crawler;

CREATE TABLE links
(
  id SERIAL PRIMARY KEY,
  parent VARCHAR(255) NOT NULL,
  link VARCHAR(255) NOT NULL,
  CONSTRAINT parent_link UNIQUE (parent, link)
);
