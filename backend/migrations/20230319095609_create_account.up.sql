CREATE TYPE account_status AS ENUM ('pending', 'approved', 'closed');

CREATE TABLE accounts (
    id uuid,
    holder_id uuid,
    currency_code varchar(3),
    amount bigint,
    opened_at timestamp without time zone,
    updated_at timestamp without time zone NOT NULL,
    expiry_data date,
    status account_status,

    PRIMARY KEY (id),
    CONSTRAINT accounts_holder_id_fk FOREIGN KEY (holder_id) REFERENCES customers (id)
);
