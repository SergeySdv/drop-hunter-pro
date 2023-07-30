-- +goose Up

create table if not exists flow (
    id uuid primary key,
    next_id uuid null references flow (id),
    label text not null,
    payload text not null,
    user_id uuid not null references users (id),
    created_at timestamp not null default now()
);

-- +goose Down

drop table if exists flow;
