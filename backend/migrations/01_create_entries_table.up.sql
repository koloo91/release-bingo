create table entries
(
    id      uuid                    not null,
    text    varchar                 not null,
    created timestamp default now() not null,
    updated timestamp default now() not null
);

create unique index entries_id_uindex
    on entries (id);

alter table entries
    add constraint entries_pk
        primary key (id);

