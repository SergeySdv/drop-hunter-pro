-- +goose Up
alter table if exists profiles
    add if not exists okex_acc_name text default null;
alter table if exists profiles
    add if not exists okex_deposit_addr text default null;

-- +goose Down
alter table if exists profiles
    drop if exists okex_deposit_addr ;
alter table if exists profiles
    drop if exists okex_acc_name ;
