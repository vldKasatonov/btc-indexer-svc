-- +migrate Up

create table users (
    id bigserial primary key,
    username text unique not null,
    password_hash text not null,
    created_at timestamptz default now()
);

-- +migrate Down

drop table users;
