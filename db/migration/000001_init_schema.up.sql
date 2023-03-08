create table users
(
    user_id       bigserial not null primary key,
    email         varchar   not null unique,
    password_hash varchar   not null
);

create table tokens
(
    user_id       int,
    device_id     text,
    access_token  text,
    refresh_token text,
    foreign key (user_id) references users (user_id)
);

create table passwords
(
    id           bigserial primary key not null,
    user_id      int                   not null,
    service_name varchar,
    link         varchar,
    email        varchar,
    login        varchar,
    password     varchar               not null,
    foreign key (user_id) references users (user_id)
);