create extension if not exists "pgcrypto";

create table if not exists stories (
    id uuid primary key default gen_random_uuid(),
    title varchar(100) not null,
    body varchar(10000) not null,
    createdat timestamp without time zone default (now() at time zone 'utc'),
    updatedat timestamp without time zone default (now() at time zone 'utc'),
    CHECK (title <> ''),
    CHECK (body <> '')
)