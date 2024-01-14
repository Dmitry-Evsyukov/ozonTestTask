create extension if not exists "uuid-ossp";

create table if not exists url (
    id uuid primary key default uuid_generate_v4(),
    original_url text not null unique,
    short_url text not null unique,
    creation_time pg_catalog.timestamptz not null default now()
);
