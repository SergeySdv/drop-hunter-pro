-- +goose Up

create table if not exists withdrawers (
    id uuid primary key,
    exchange_type text not null,
    label text not null,
    proxy text  null,
    secret_key bytea not null,
    api_key bytea not null,
    user_id uuid not null references users (id),
    created_at timestamp not null default now()
);

-- +goose Down

drop table if exists withdrawers;
