-- MySQL Script generated by MySQL Workbench
-- Wed Aug  9 18:09:05 2023
-- Model: New Model    Version: 1.0
-- MySQL Workbench Forward Engineering

SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0;
SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0;
SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='ONLY_FULL_GROUP_BY,STRICT_TRANS_TABLES,NO_ZERO_IN_DATE,NO_ZERO_DATE,ERROR_FOR_DIVISION_BY_ZERO,NO_ENGINE_SUBSTITUTION';

-- -----------------------------------------------------
-- Schema ISPS_plates
-- -----------------------------------------------------

-- -----------------------------------------------------
-- Schema ISPS_plates
-- -----------------------------------------------------
CREATE SCHEMA IF NOT EXISTS `ISPS_plates` DEFAULT CHARACTER SET utf8 ;
USE `ISPS_plates` ;

-- -----------------------------------------------------
-- Table `ISPS_plates`.`plates`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `ISPS_plates`.`plates` (
  `plate_id` INT NOT NULL AUTO_INCREMENT,
  `plate_name` VARCHAR(45) NULL,
  `views` INT NULL,
  `plate_type` VARCHAR(45) NULL,
  `posts_number` INT NULL,
  `plate_time` VARCHAR(45) NULL,
  `moderator_id` INT NULL,
  PRIMARY KEY (`plate_id`),
  INDEX `fk_plate_idx` (`plate_id` ASC) VISIBLE)
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `ISPS_plates`.`posts`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `ISPS_plates`.`posts` (
  `post_id` INT NOT NULL AUTO_INCREMENT,
  `post_name` VARCHAR(45) NULL,
  `author` VARCHAR(45) NULL,
  `views` VARCHAR(45) NULL,
  `post_type` VARCHAR(45) NULL,
  `like_number` INT NULL,
  `post_time` VARCHAR(45) NULL,
  `plate_id` INT NOT NULL,
  PRIMARY KEY (`post_id`),
  INDEX `fk_posts_plates_idx` (`plate_id` ASC) VISIBLE,
  CONSTRAINT `fk_posts_plates`
    FOREIGN KEY (`plate_id`)
    REFERENCES `ISPS_plates`.`plates` (`plate_id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION)
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `ISPS_plates`.`announcements`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `ISPS_plates`.`announcements` (
  `announcement_id` INT NOT NULL AUTO_INCREMENT,
  `announcement_name` VARCHAR(45) NULL,
  `announcement_time` VARCHAR(45) NULL,
  `announcement_information` VARCHAR(10000) NULL,
  `plate_id` INT NOT NULL,
  PRIMARY KEY (`announcement_id`),
  INDEX `fk_announcements_plates1_idx` (`plate_id` ASC) VISIBLE,
  CONSTRAINT `fk_announcements_plates1`
    FOREIGN KEY (`plate_id`)
    REFERENCES `ISPS_plates`.`plates` (`plate_id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION)
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `ISPS_plates`.`manage`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `ISPS_plates`.`manage` (
  `plate_id` INT NOT NULL,
  `manage_number` INT NULL,
  `manage_id` VARCHAR(16830) NULL,
  PRIMARY KEY (`plate_id`),
  UNIQUE INDEX `plate_id_UNIQUE` (`plate_id` ASC) VISIBLE,
  INDEX `fk_ma_idx` (`plate_id` ASC) VISIBLE,
  CONSTRAINT `fk_manage_plates1`
    FOREIGN KEY (`plate_id`)
    REFERENCES `ISPS_plates`.`plates` (`plate_id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION)
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `ISPS_plates`.`fans`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `ISPS_plates`.`fans` (
  `plate_id` INT NOT NULL,
  `fans_number` INT NULL,
  `fans_id` VARCHAR(16380) NULL,
  PRIMARY KEY (`plate_id`),
  UNIQUE INDEX `plate_id_UNIQUE` (`plate_id` ASC) VISIBLE,
  INDEX `fk_fn_idx` (`plate_id` ASC) VISIBLE,
  CONSTRAINT `fk_fans_plates1`
    FOREIGN KEY (`plate_id`)
    REFERENCES `ISPS_plates`.`plates` (`plate_id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION)
ENGINE = InnoDB;


SET SQL_MODE=@OLD_SQL_MODE;
SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS;
SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS;