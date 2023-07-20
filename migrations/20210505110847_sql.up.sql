CREATE TABLE IF NOT EXISTS todos
(
    id          SERIAL PRIMARY KEY,
    title       VARCHAR(255),
    description TEXT
);