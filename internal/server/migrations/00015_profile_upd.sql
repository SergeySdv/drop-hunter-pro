-- +goose Up
alter table if exists profiles
       add if not exists user_agent text default '';

alter table if exists users
       drop if exists user_agent_header;

-- +goose Down
alter table if exists profiles
       drop if exists user_agent;

alter table if exists users
       add if not exists user_agent_header text default '';