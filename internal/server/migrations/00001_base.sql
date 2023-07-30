-- +goose Up

create table if not exists users (
    id uuid primary key,
    email text not null unique,
    access boolean not null default false
);

create table if not exists profiles (
    id uuid primary key,
    label text not null,
    proxy text  null,
    mmsk_pk bytea not null,
    meta text  null,
    user_id uuid not null references users (id),
    created_at timestamp not null default now()
);

-- +goose Down
drop table if exists users;
drop table if exists profiles;
