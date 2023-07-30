-- +goose Up
alter table if exists process
    add if not exists auto_retry boolean default false;

-- +goose Down
alter table if exists process
    drop if exists auto_retry;

