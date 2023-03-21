CREATE TABLE credit_accounts (
    account_id uuid,
    credit_limit bigint,
    debt bigint,
    accrued_interest integer,
    credit_rate integer,

    UNIQUE (account_id),
    CONSTRAINT credit_accounts_account_id_fk FOREIGN KEY (account_id) REFERENCES accounts (id)
);
