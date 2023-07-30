-- +goose Up
alter table if exists flow
    add if not exists deleted_at timestamp null default null;

alter table if exists process
    add if not exists deleted_at timestamp null default null;

-- +goose Down
alter table if exists flow
    drop if exists deleted_at;

alter table if exists process
    drop if exists deleted_at;

