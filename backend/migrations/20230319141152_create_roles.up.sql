CREATE TABLE roles (
    id integer PRIMARY KEY,
    name varchar(30) UNIQUE NOT NULL
);

INSERT INTO roles (id, name) VALUES (1, 'admin');
INSERT INTO roles (id, name) VALUES (2, 'moderator');
