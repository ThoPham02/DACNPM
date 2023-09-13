CREATE TABLE `user_tbl`
(
    `id`           bigint auto_increment,
    `full_name`    varchar(255) NOT NULL,
    `user_name`    varchar(255) NOT NULL,
    `password`     varchar(255) NOT NULL,
    `email`        varchar(255) NOT NULL,
    `role`         int          NOT NULL,
    `created_at`   TIMESTAMP NOT NULL,
    `updated_at`   TIMESTAMP NOT NULL,
    PRIMARY KEY (id)
);