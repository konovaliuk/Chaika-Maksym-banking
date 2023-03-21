CREATE TABLE managers (
    id uuid,
    full_name varchar(255) NOT NULL,
    password_hash varchar(60) NOT NULL,
    created_at timestamp without time zone NOT NULL,

    PRIMARY KEY (id)
);

CREATE TABLE manager_roles (
    manager_id uuid,
    role_id integer,

    CONSTRAINT manager_roles_manager_id_fk FOREIGN KEY (manager_id) REFERENCES managers (id),
    CONSTRAINT manager_roles_role_id_fk FOREIGN KEY (role_id) REFERENCES roles (id)
)
