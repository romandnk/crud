CREATE TABLE IF NOT EXISTS tasks (
    id serial PRIMARY KEY,
    creation_time TIMESTAMP NOT NULL,
    updating_time TIMESTAMP NOT NULL,
    message varchar(255) NOT NULL
);