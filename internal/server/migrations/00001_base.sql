-- +goose Up

create table if not exists profiles (
    id uuid primary key,
    label text not null,
    proxy text  null,
    mmsk_pk text not null,
    meta text  null,
    user_id uuid not null,
    created_at timestamp not null default now()
);

-- +goose Down

drop table if exists profiles;
