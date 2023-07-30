-- +goose Up
alter table if exists withdrawers
    add if not exists prev_id uuid null references withdrawers (id);


create table if not exists okex_deposit_addr_profile
(
    withdrawer_id uuid not null references withdrawers (id),
    okex_deposit_addr text not null,
    profile_id uuid not null references profiles (id),
    user_id uuid references users (id),
    primary key (withdrawer_id, okex_deposit_addr, profile_id, user_id)
);

-- +goose Down
alter table if exists withdrawers
    drop if exists prev_id;

drop table if exists okex_deposit_addr_profile;
