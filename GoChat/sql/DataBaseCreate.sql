create database GoChat;

create table gochat.user
(
    id             bigint unsigned auto_increment primary key,
    created_at     datetime(3) null,
    updated_at     datetime(3) null,
    deleted_at     datetime(3) null,
    name           longtext    null,
    password       longtext    null,
    telephone      longtext    null,
    email          longtext    null,
    client_ip      longtext    null,
    client_port    longtext    null,
    is_login_out   tinyint(1)  null,
    login_time     datetime(3) null,
    heartbeat_time datetime(3) null,
    loginout_time  datetime(3) null,
    device_info    longtext    null
);

create index idx_user_deleted_at
    on gochat.user (deleted_at);

