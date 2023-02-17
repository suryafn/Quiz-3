-- +migrate Up
-- +migrate StatementBegin

create table category
(
    id         serial primary key not null,
    name       varchar(256)       not null,
    created_at timestamp,
    updated_at timestamp
)

-- +migrate StatementEnd