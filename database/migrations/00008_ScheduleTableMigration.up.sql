start transaction;

create extension if not exists "uuid-ossp";

create table if not exists schedules (
    id uuid default uuid_generate_v4() not null primary key,
    date timestamp with time zone not null,
    home_team_id uuid constraint fk_schedule_team_home references teams,
    away_team_id uuid constraint fk_schedule_team_away references teams,
    winner_team_id uuid constraint fk_schedule_team_winner references teams,
    is_finished boolean not null,
    created_at timestamp with time zone not null default current_timestamp,
    updated_at timestamp with time zone not null default current_timestamp,
    deleted_at timestamp with time zone null
                             );

create index if not exists idx_schedules_deleted_at on schedules (deleted_at);

commit;