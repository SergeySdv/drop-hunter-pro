-- +goose Up
alter table if exists profiles
       drop constraint if exists check_pk_uniq;

alter table if exists profiles
       drop constraint if exists profiles_mmsk_id_key;

create unique index if not exists profiles_pk_deleted_at
       on profiles (user_id, mmsk_id, (deleted_at is null))
       where deleted_at is null;
-- +goose Down
