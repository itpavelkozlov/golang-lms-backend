-- +goose Up
-- +goose StatementBegin
CREATE TABLE Users
(
    id         varchar(36),
    email      varchar(320) NOT NULL,
    first_name varchar(255),
    last_name  varchar(255),
    password   varchar(255) NOT NULL,
    created_at TIMESTAMPTZ  NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ,
    deleted_at TIMESTAMPTZ,
    constraint User_id_pk primary key (id),
    constraint User_email_Unique unique (email),
    constraint User_email_Valid check (email ~* '^[A-Za-z0-9._%-]+@[A-Za-z0-9.-]+[.][A-Za-z]+$'),
    constraint User_id_Valid check (char_length(ID) = 36)
)
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS Users;
-- +goose StatementEnd
