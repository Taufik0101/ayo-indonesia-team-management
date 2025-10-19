start transaction;

create extension if not exists "uuid-ossp";

create table if not exists detail_results (
    id uuid default uuid_generate_v4() not null primary key,
    result_id uuid constraint fk_result_detail_result references results,
    player_id uuid constraint fk_result_detail_player references players,
    goal_time varchar(50) not null,
    is_penalty boolean not null,
    created_at timestamp with time zone not null default current_timestamp,
    updated_at timestamp with time zone not null default current_timestamp,
    deleted_at timestamp with time zone null
                             );

create index if not exists idx_detail_results_deleted_at on detail_results (deleted_at);

commit;