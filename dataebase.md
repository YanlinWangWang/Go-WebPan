数据库创建语句
CREATE TABLE `fileserver`.`tbl_file` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `file_sha1` CHAR(40) NOT NULL DEFAULT '',
  `file_name` VARCHAR(256) NOT NULL DEFAULT '',
  `file_size` VARCHAR(45) NOT NULL DEFAULT 0,
  `file_addr` VARCHAR(1024) NOT NULL DEFAULT '',
  `create_at` DATETIME NOT NULL DEFAULT NOW(),
  `update_at` DATETIME NOT NULL DEFAULT NOW(),
  `status` INT NULL,
  PRIMARY KEY (`id`));