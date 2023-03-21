CREATE TABLE deposit_accounts (
    account_id uuid,
    deposit_amount bigint,
    annual_rate integer,

    UNIQUE (account_id),
    CONSTRAINT credit_accounts_account_id_fk FOREIGN KEY (account_id) REFERENCES accounts (id)
);
