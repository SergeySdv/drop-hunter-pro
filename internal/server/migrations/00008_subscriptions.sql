-- +goose Up
create table if not exists alert_subscriptions (
    user_id uuid primary key references users (id),
    telegram_chat_id text not null
);

-- +goose Down
drop table if exists alert_subscriptions;
