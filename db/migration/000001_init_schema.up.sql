create table users
(
    id       bigserial not null primary key,
    email    varchar   not null unique,
    password varchar   not null
)