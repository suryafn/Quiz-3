-- +migrate Up
-- +migrate StatementBegin

create table book
(
    id           serial primary key not null,
    title        varchar(256)       not null,
    description  varchar(256)       not null,
    image_url    varchar(256)       not null,
    release_year int                not null,
    price        int                not null,
    total_page   int                not null,
    thickness    varchar(256)       not null,
    created_at   timestamp,
    updated_at   timestamp,
    category_id  int                not null,
    foreign key (category_id) references category (id)
)

-- +migrate StatementEnd