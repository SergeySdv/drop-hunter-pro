-- +goose Up

create table if not exists withdrawers (
    id uuid primary key,
    exchange_type text not null,
    label text not null,
    proxy text  null,
    secret_key text not null,
    api_key text not null,
    user_id uuid not null,
    created_at timestamp not null default now()
);

-- +goose Down

drop table if exists withdrawers;
