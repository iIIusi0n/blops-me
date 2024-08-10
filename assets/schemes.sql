create table storage
(
    id           int auto_increment
        primary key,
    storage_name varchar(32) not null,
    user_id      varchar(64) not null
);

create table file
(
    id            int auto_increment
        primary key,
    name          varchar(128) not null,
    type          varchar(16)  not null,
    last_modified date         not null,
    size          int          not null,
    path          varchar(256) not null,
    storage_id    int          not null,
    parent_id     int          null,
    constraint file_file_id_fk
        foreign key (parent_id) references file (id),
    constraint file_storage_id_fk
        foreign key (storage_id) references storage (id)
);
