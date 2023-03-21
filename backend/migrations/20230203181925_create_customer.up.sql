CREATE TABLE customers (
    id uuid,
    email varchar(255) NOT NULL,
    password_hash varchar(60) NOT NULL,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL,
    PRIMARY KEY (id),
    UNIQUE (email)
);
