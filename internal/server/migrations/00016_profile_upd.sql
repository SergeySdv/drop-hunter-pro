-- +goose Up
ALTER TABLE profiles DROP CONSTRAINT IF EXISTS check_num_uniq;
alter table if exists profiles
       add constraint check_num_uniq
              unique (user_id, num);

ALTER TABLE profiles DROP CONSTRAINT IF EXISTS check_pk_uniq;
alter table if exists profiles
       add constraint check_pk_uniq
              unique (user_id, mmsk_pk);

alter table if exists profiles
       add if not exists mmsk_id bytea not null unique default cast(floor(random() * 10000) as text)::bytea;

-- +goose Down
alter table if exists profiles
       drop constraint if exists check_num_uniq;

alter table if exists profiles
       drop constraint if exists check_pk_uniq;

alter table if exists profiles
       drop if exists mmsk_id;