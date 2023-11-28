-- +goose Up
create table if not exists users (
    user_id text primary key not null default pre_appended_uuid('users'),
    user_name text not null unique,
    phone_number text not null unique,
    created_at timestamp not null default now()
);
create index if not exists user_name_index on users(user_name);

-- +goose Down
drop table if exists users;
