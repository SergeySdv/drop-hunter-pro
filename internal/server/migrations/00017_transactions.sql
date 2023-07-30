-- +goose Up
create table if not exists transactions (
    id text primary key,
    code text not null,
    task_id uuid not null references process_tasks (id),
    profile_id uuid not null references profiles (id),
    process_id uuid not null references process (id),
    created_at timestamp default now(),
    user_id uuid references users (id),
    url text not null,
    network text not null
);

-- +goose Down
drop table if exists transactions;
