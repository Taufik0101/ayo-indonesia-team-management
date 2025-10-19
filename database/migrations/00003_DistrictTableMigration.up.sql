start transaction;

create extension if not exists "uuid-ossp";

create table if not exists area_districts (
    id uuid default uuid_generate_v4() not null primary key,
    nama_kab varchar(255) not null,
    no_prop int not null,
    no_kab int not null,
    created_at timestamp with time zone not null default current_timestamp,
    updated_at timestamp with time zone not null default current_timestamp,
    deleted_at timestamp with time zone null
                             );

create index if not exists idx_area_districts_deleted_at on area_districts (deleted_at);

create unique index if not exists idx_area_districts_no_prop_no_kab ON area_districts (no_prop, no_kab) WHERE (deleted_at is NULL);

INSERT INTO area_districts (id, nama_kab, no_prop, no_kab, created_at, updated_at, deleted_at) VALUES ('280f1804-0cff-4053-8138-879feccfa8e7', 'ACEH SELATAN', 11, 1, '2025-10-18 13:50:23.648147 +00:00', '2025-10-18 13:50:23.648147 +00:00', null);

commit;