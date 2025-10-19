start transaction;

create extension if not exists "uuid-ossp";

create table if not exists teams (
    id uuid default uuid_generate_v4() not null primary key,
    name varchar(255) not null,
    logo text not null,
    year int not null,
    address text not null,
    province_id uuid constraint fk_teams_province references area_provinces,
    district_id uuid constraint fk_teams_district references area_districts,
    sub_district_id uuid constraint fk_teams_sub_district references area_sub_districts,
    village_id uuid constraint fk_teams_village references area_villages,
    created_at timestamp with time zone not null default current_timestamp,
    updated_at timestamp with time zone not null default current_timestamp,
    deleted_at timestamp with time zone null
                             );

create index if not exists idx_teams_deleted_at on teams (deleted_at);

commit;