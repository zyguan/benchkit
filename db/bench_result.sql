create table if not exists `bench_result` (
    `id`       varchar(40),
    `name`     varchar(255),
    `cmd`      text,
    `started`  bigint,
    `finished` bigint,
    `exit`     int,
    `error`    text,
    `stdout`   longtext,
    `stderr`   longtext,
    primary key (`id`) /*T![clustered_index] CLUSTERED */,
    key (started)
);
