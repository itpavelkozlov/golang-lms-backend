-- +goose Up
-- +goose StatementBegin
CREATE TABLE Users
(
    ID  varchar (36),
    Email  varchar(320) NOT NULL,
    FirstName varchar(255),
    LastName   varchar(255),
    Password varchar(255) NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ,
    constraint ID primary key (ID),
    constraint EmailUnique unique (Email),
    constraint EmailValid check (Email ~* '^[A-Za-z0-9._%-]+@[A-Za-z0-9.-]+[.][A-Za-z]+$'),
    constraint IDValid check (char_length(ID) = 36)
)
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE Users;
-- +goose StatementEnd
