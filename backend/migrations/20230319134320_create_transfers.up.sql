CREATE TYPE transfer_status AS ENUM ('pending', 'approved', 'declined');

CREATE TABLE transfers (
    id uuid,
    sender_account_id uuid,
    recipient_account_id uuid,
    amount bigint,
    currency_code varchar(3),
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL,
    status transfer_status,

    PRIMARY KEY (id),
    CONSTRAINT transfers_sender_account_id_fk FOREIGN KEY (sender_account_id) REFERENCES accounts (id),
    CONSTRAINT transfers_recipient_account_id_fk FOREIGN KEY (recipient_account_id) REFERENCES accounts (id)
);
