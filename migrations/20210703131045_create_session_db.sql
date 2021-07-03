-- +goose Up
-- +goose StatementBegin
CREATE TABLE Session
(
    id            varchar(36),
    user_id       varchar(36) NOT NULL,
    refresh_token varchar(54) NOT NULL,
    created_at    TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at    TIMESTAMPTZ,
    deleted_at    TIMESTAMPTZ,
    constraint Session_id_pk primary key (id),
    constraint Session_userId_Unique unique (user_id),
    constraint Session_id_Valid check (char_length(id) = 36),
    constraint Session_refreshToken_Valid check (char_length(refresh_token) = 54),
    constraint Session_userId_Valid check (char_length(user_id) = 36),
    constraint Session_userId_fk FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE CASCADE
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS Session;
-- +goose StatementEnd
