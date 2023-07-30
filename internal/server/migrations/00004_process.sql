-- +goose Up

create table if not exists process (
    id uuid primary key,
    status text not null,
    user_id uuid not null references users (id),
    flow_id uuid not null references flow (id) ON DELETE CASCADE,
    payload text not null,
    updated_at timestamp not null default now(),
    created_at timestamp not null default now(),
    started_at timestamp null,
    finished_at timestamp null
);

create table if not exists process_profiles (
    id uuid primary key,
    weight int not null,
    process_id uuid references process (id),
    profile_id uuid references profiles (id),
    status text not null
);

create table if not exists process_tasks (
    id uuid primary key,
    weight int not null,
    process_id uuid references process (id),
    profile_id uuid references process_profiles (id),
    status text not null,
    payload text not null
);

-- +goose Down

drop table if exists process_tasks;
drop table if exists process;
