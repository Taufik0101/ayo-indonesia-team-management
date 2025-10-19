start transaction;

create extension if not exists "uuid-ossp";

create table if not exists area_villages (
    id uuid default uuid_generate_v4() not null primary key,
    nama_kel varchar(255) not null,
    no_prop int not null,
    no_kab int not null,
    no_kec int not null,
    no_kel int not null,
    created_at timestamp with time zone not null default current_timestamp,
    updated_at timestamp with time zone not null default current_timestamp,
    deleted_at timestamp with time zone null
                             );

create index if not exists idx_area_villages_deleted_at on area_villages (deleted_at);

create unique index if not exists idx_area_villages_no_prop_no_kab_no_kec ON area_villages (no_prop, no_kab, no_kec, no_kel) WHERE (deleted_at is NULL);

INSERT INTO area_villages (id, nama_kel, no_prop, no_kab, no_kec, no_kel, created_at, updated_at, deleted_at) VALUES ('00e057ba-586c-4cf0-bc08-fd59186f299e', 'KEUDE MEUKEK', 11, 1, 5, 2016, '2025-10-18 23:15:41.410654 +00:00', '2025-10-18 23:15:41.410654 +00:00', null);

commit;