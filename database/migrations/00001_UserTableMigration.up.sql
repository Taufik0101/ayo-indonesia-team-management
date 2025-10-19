start transaction;

DROP TYPE IF EXISTS role_types;
CREATE TYPE role_types AS enum ('admin', 'user');

create extension if not exists "uuid-ossp";

create table if not exists users (
    id uuid default uuid_generate_v4() not null primary key,
    name varchar(255) not null,
    email varchar(255) not null,
    password text null,
    role    role_types                    NOT NULL DEFAULT 'user',
    created_at timestamp with time zone not null default current_timestamp,
    updated_at timestamp with time zone not null default current_timestamp,
    deleted_at timestamp with time zone null
                             );

create index if not exists idx_users_deleted_at on users (deleted_at);

create unique index if not exists idx_users_email ON users (email) WHERE (deleted_at is NULL);

INSERT INTO users (id, name, email, password, role, created_at, updated_at, deleted_at) VALUES ('538a0a16-c940-4e1e-b7c9-6525a1688c87', 'Test Admin', 'testadmin@gmail.com', '$2a$10$uYOMQaYf0pExkpEuwyRx9uK1dnM3CTjFFwNQg56XOzlEz0bpObQaO', 'admin', '2025-10-18 13:37:26.270847 +00:00', '2025-10-18 13:37:26.270847 +00:00', null);

commit;