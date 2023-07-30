-- +goose Up
alter table if exists users
       add if not exists user_agent_header text default '';

-- +goose Down
alter table if exists users
       drop if exists user_agent_header;