begin;

create table if not exists users
(
    user_id       bigserial    not null primary key,
    email         varchar(180) not null unique,
    password_hash varchar(255) not null
);

commit;