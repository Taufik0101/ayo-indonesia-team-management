start transaction;

create extension if not exists "uuid-ossp";

create table if not exists results (
    id uuid default uuid_generate_v4() not null primary key,
    schedule_id uuid constraint fk_result_schedule references schedules,
    score_home int not null,
    score_away int not null,
    winner_team_id uuid constraint fk_schedule_team_winner references teams,
    created_at timestamp with time zone not null default current_timestamp,
    updated_at timestamp with time zone not null default current_timestamp,
    deleted_at timestamp with time zone null
                             );

create index if not exists idx_results_deleted_at on results (deleted_at);

commit;