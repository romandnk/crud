CREATE TABLE IF NOT EXISTS task (
    id serial PRIMARY KEY,
    creation_time varchar(255),
    updating_time varchar(255),
    message varchar(255)
);