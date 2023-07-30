-- +goose Up
begin;
alter table if exists profiles
       add if not exists num integer default 0;

update profiles
set num = cast (regexp_replace(label, '[^0-9]+','', 'g') as integer)
where regexp_replace(label, '[^0-9]+','', 'g') != '';

end;

-- +goose Down
alter table if exists profiles
       drop if exists num;