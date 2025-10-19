start transaction;

create extension if not exists "uuid-ossp";

create table if not exists area_sub_districts (
    id uuid default uuid_generate_v4() not null primary key,
    nama_kec varchar(255) not null,
    no_prop int not null,
    no_kab int not null,
    no_kec int not null,
    created_at timestamp with time zone not null default current_timestamp,
    updated_at timestamp with time zone not null default current_timestamp,
    deleted_at timestamp with time zone null
                             );

create index if not exists idx_area_sub_districts_deleted_at on area_sub_districts (deleted_at);

create unique index if not exists idx_area_sub_districts_no_prop_no_kab_no_kec ON area_sub_districts (no_prop, no_kab, no_kec) WHERE (deleted_at is NULL);

INSERT INTO area_sub_districts (id, nama_kec, no_prop, no_kab, no_kec, created_at, updated_at, deleted_at) VALUES ('3eb0c123-6aea-4c59-8ac3-bb63e8df0b1f', 'MEUKEK', 11, 1, 5, '2025-10-18 14:29:16.913772 +00:00', '2025-10-18 14:29:16.913772 +00:00', null);

commit;