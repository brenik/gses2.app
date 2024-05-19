CREATE DATABASE IF NOT EXISTS gses;

use gses;

CREATE TABLE IF NOT EXISTS gses.subscribers
(
    id         int(10) unsigned NOT NULL AUTO_INCREMENT,
    email      varchar(255) NOT NULL,
    created_at datetime     NOT NULL DEFAULT current_timestamp(),
    PRIMARY KEY (id)
);

INSERT INTO subscribers (email, created_at)
VALUES ('vitaliy.brenyk@gmail.com', '2024-05-18 00:59:51'),
       ('vitaliy_brenik@outlook.com', '2024-05-18 02:32:15');
