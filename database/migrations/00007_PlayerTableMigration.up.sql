start transaction;

DROP TYPE IF EXISTS player_position_types;
CREATE TYPE player_position_types AS enum ('penyerang', 'gelandang', 'bertahan', 'penjaga_gawang');

create extension if not exists "uuid-ossp";

create table if not exists players (
    id uuid default uuid_generate_v4() not null primary key,
    name varchar(255) not null,
    height int not null,
    weight int not null,
    number int not null,
    position    player_position_types                    NOT NULL DEFAULT 'gelandang',
    team_id uuid constraint fk_player_team references teams,
    created_at timestamp with time zone not null default current_timestamp,
    updated_at timestamp with time zone not null default current_timestamp,
    deleted_at timestamp with time zone null
                             );

create index if not exists idx_players_deleted_at on players (deleted_at);

create unique index if not exists idx_players_number_team ON players (number, team_id) WHERE (deleted_at is NULL);

commit;