begin;

create table if not exists passwords
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

commit;