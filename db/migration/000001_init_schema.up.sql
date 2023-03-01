create table users
(
    id            bigserial not null primary key,
    email         varchar   not null unique,
    password_hash varchar   not null
);

create table tokens
(
    id            int,
    device_id     text,
    access_token  text,
    refresh_token text,
    foreign key (id) references users (id)
);