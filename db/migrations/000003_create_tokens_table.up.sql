begin;

create table if not exists tokens
(
    user_id       int not null ,
    device_id     text not null,
    access_token  text not null,
    refresh_token text not null,
    foreign key (user_id) references users (user_id) on delete cascade
);

commit;