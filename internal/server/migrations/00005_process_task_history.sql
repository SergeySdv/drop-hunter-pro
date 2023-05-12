-- +goose Up
create table if not exists process_task_history (
    id uuid primary key,
    task_id uuid references process_tasks (id),
    started_at timestamp not null,
    finished_at timestamp  null,
    start_status text not null,
    finish_status text  null,
    msg text null
);

-- +goose Down
drop table if exists process_task_history;
