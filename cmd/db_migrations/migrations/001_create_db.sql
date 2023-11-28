-- +goose Up
create extension if not exists postgis;
select postgis_full_version();
create extension if not exists btree_gist;
create extension if not exists cube;
create extension if not exists earthdistance;

create function pre_appended_uuid(table_name text) 
    returns text 
    return table_name || '_' || gen_random_uuid()::text;

create table if not exists users (
    user_id text primary key not null default pre_appended_uuid('users'),
    user_name text not null unique,
    phone_number text not null unique,
    created_at timestamp not null default now()
);
create index if not exists user_name_index on users(user_name);

create table if not exists tokens (
    token_id text primary key not null default pre_appended_uuid('tokens'),
    user_id text not null REFERENCES users (user_id),
    created_at timestamp not null default now(),
    expires_at timestamp not null default now() + '7 days'::interval
);

create table if not exists locations (
    location_id text primary key not null default pre_appended_uuid('locations'),
    geom geography(Point, 4326), -- lat / long magic
    created_at timestamp default now()
);
create index if not exists locations_geom_index on locations using spgist(geom);

create type venue as enum (
    'coffee shop',
    'gym',
    'hotel',
    'restaurant',
    'bar',
    'gas station',
    'other'
);


-- really really want to call it "bowls", but alas, 
-- that is not a good naming convention to hold
create table if not exists  bathrooms (
    bathroom_id text primary key not null default pre_appended_uuid('bathrooms'),
    title text not null, 
    location_id text not null REFERENCES locations (location_id),
    created_at timestamp default now(),
    description text not null,
    last_used timestamp,
    last_updated timestamp,
    last_updated_by text REFERENCES users (user_id),
    is_public boolean not null default false,
    is_customer_only boolean not null default false,
    is_pay_to_use boolean not null default false,
    venue_type venue not null,
    flagged boolean default false
);

-- +goose Down
drop extension if exists  postgis;
drop extension if exists btree_gist;
drop extension if exists cube;
drop extension if exists earthdistance;
drop function custom_uuid;
drop table if exists users;
drop table if exists tokens;
drop table if exists locations;
drop type if exists venue;
drop table if exists bowls;
