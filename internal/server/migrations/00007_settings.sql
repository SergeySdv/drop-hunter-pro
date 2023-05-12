-- +goose Up
create table if not exists settings (
    user_id uuid primary key,
    payload text not null
);

-- +goose Down
drop table if exists settings;
