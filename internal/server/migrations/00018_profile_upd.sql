-- +goose Up
alter table if exists profiles
       add if not exists deleted_at  timestamp null;
-- +goose Down
alter table if exists profiles
       drop if exists deleted_at;