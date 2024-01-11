-- +goose Up
create table "user"
(
    uuid       uuid
        constraint user_pk
            primary key,
    created_at timestamp not null
);


-- +goose Down
drop table "user";
