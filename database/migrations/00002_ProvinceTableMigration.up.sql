start transaction;

create extension if not exists "uuid-ossp";

create table if not exists area_provinces (
    id uuid default uuid_generate_v4() not null primary key,
    nama_prop varchar(255) not null,
    no_prop int not null,
    created_at timestamp with time zone not null default current_timestamp,
    updated_at timestamp with time zone not null default current_timestamp,
    deleted_at timestamp with time zone null
                             );

create index if not exists idx_area_provinces_deleted_at on area_provinces (deleted_at);

create unique index if not exists idx_area_provinces_no_prop ON area_provinces (no_prop) WHERE (deleted_at is NULL);

INSERT INTO area_provinces (id, nama_prop, no_prop, created_at, updated_at, deleted_at) VALUES ('957a2efb-fa3a-482d-b482-e4901fe69004', 'ACEH', 11, '2025-10-18 13:46:04.949762 +00:00', '2025-10-18 13:46:04.949762 +00:00', null);

commit;