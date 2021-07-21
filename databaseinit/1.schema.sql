create table if not exists users(
    id serial primary key,
    user_name text not null,
    password_hash text not null,
    created_at timestamp with time zone not null default now(),
    created_by int,
    updated_at timestamp with time zone not null default now(),
    updated_by int
);

alter table users add constraint fk_created_by foreign key (created_by) references users(id);
alter table users add constraint fk_updated_by foreign key (updated_by) references users(id);
--
-- drop table users