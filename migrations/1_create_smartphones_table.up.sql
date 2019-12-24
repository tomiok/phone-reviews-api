CREATE TABLE smartphone
(
    `id`             int(11) not null auto_increment,
    `name`           varchar(150),
    `country_origin` varchar(150),
    `os`             varchar(100),
    `price`          int(10),
    `created_at`     DATETIME DEFAULT CURRENT_TIMESTAMP,
    `updated_at`     DATETIME ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`)
) engine = InnoDB
  DEFAULT charset = utf8;