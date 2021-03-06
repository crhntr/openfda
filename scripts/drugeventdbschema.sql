-- MySQL Script generated by MySQL Workbench
-- Mon May 14 12:13:27 2018
-- Model: New Model    Version: 1.0
-- MySQL Workbench Forward Engineering

SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0;
SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0;
SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='TRADITIONAL,ALLOW_INVALID_DATES';

-- -----------------------------------------------------
-- Schema drugeventdb
-- -----------------------------------------------------

-- -----------------------------------------------------
-- Schema drugeventdb
-- -----------------------------------------------------
CREATE SCHEMA IF NOT EXISTS `drugeventdb` DEFAULT CHARACTER SET utf8 ;
USE `drugeventdb` ;

-- -----------------------------------------------------
-- Table `drugeventdb`.`Country`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `drugeventdb`.`Country` (
  `country_id` INT NULL,
  `country_name` VARCHAR(45) NOT NULL,
  PRIMARY KEY (`country_id`))
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `drugeventdb`.`Route`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `drugeventdb`.`Route` (
  `route_id` INT NOT NULL,
  `route_description` VARCHAR(45) BINARY NULL COMMENT 'The drug’s route of administration.',
  PRIMARY KEY (`route_id`))
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `drugeventdb`.`Manufacturer`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `drugeventdb`.`Manufacturer` (
  `manufacturer_id` INT NOT NULL,
  `manufacturer_name` VARCHAR(45) NULL,
  PRIMARY KEY (`manufacturer_id`))
ENGINE = InnoDB;


-- -- -----------------------------------------------------
-- -- Table `drugeventdb`.`Ingredient`
-- -- -----------------------------------------------------
-- CREATE TABLE IF NOT EXISTS `drugeventdb`.`Ingredient` (
--   `ingredient_id` INT NOT NULL,
--   `ingredient_name` VARCHAR(45) NULL,
--   PRIMARY KEY (`ingredient_id`))
-- ENGINE = InnoDB;

-- -----------------------------------------------------
-- Table `drugeventdb`.`Drug`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `drugeventdb`.`Drug` (
  `drug_splset_id` CHAR(36) NOT NULL,
  PRIMARY KEY (`drug_splset_id`),
  INDEX `drug_splset_id_idx` (`drug_splset_id` ASC))
ENGINE = InnoDB;

-- -----------------------------------------------------
-- Table `drugeventdb`.`DrugName`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `drugeventdb`.`DrugName` (
  `drug_name_id` INT NOT NULL,
  `drug_name_splset_id`  CHAR(36) NOT NULL,
  `drug_name_name` CHAR(127),
  `drug_name_cutoff` TINYINT NOT NULL,
  `drug_name_type` ENUM('Brand', 'Generic', 'Substance') NOT NULL,
  PRIMARY KEY (`drug_name_id`),
  INDEX `drug_name_idx` (`drug_name_name` ASC),
  INDEX `drug_name_splset_id_idx` (`drug_name_splset_id` ASC),
  CONSTRAINT `drug_name_splset_id`
    FOREIGN KEY (`drug_name_splset_id`)
    REFERENCES `drugeventdb`.`Drug` (`drug_splset_id`)
) ENGINE = InnoDB;

-- -----------------------------------------------------
-- Table `drugeventdb`.`DrugClass`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `drugeventdb`.`DrugManufacturer` (
  `drug_splset_id` CHAR(36) NOT NULL,
  `manufacturer_id` INT NULL,
  PRIMARY KEY (`drug_splset_id`, `manufacturer_id`),
  INDEX `drug_splset_id_idx` (`drug_splset_id` ASC),
  INDEX `manufacturer_id_idx` (`manufacturer_id` ASC),
  CONSTRAINT `drug_splset_id`
    FOREIGN KEY (`drug_splset_id`)
    REFERENCES `drugeventdb`.`Drug` (`drug_splset_id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION,
  CONSTRAINT `manufacturer_id`
    FOREIGN KEY (`manufacturer_id`)
    REFERENCES `drugeventdb`.`Manufacturer` (`manufacturer_id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION)
ENGINE = InnoDB;


-- -- -----------------------------------------------------
-- -- Table `drugeventdb`.`DrugIngredient`
-- -- -----------------------------------------------------
-- CREATE TABLE IF NOT EXISTS `drugeventdb`.`DrugIngredient` (
--   `splset_id` CHAR(12) NOT NULL,
--   `ingredient_id` INT NOT NULL,
--   PRIMARY KEY (`splset_id`, `ingredient_id`),
--   INDEX `di_ingredientId_idx` (`ingredient_id` ASC),
--   CONSTRAINT `di_splset_id`
--     FOREIGN KEY (`splset_id`)
--     REFERENCES `drugeventdb`.`Drug` (`splset_id`)
--     ON DELETE NO ACTION
--     ON UPDATE NO ACTION,
--   CONSTRAINT `di_ingredientId`
--     FOREIGN KEY (`ingredient_id`)
--     REFERENCES `drugeventdb`.`Ingredient` (`ingredient_id`)
--     ON DELETE NO ACTION
--     ON UPDATE NO ACTION)
-- ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `drugeventdb`.`Class`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `drugeventdb`.`Class` (
  `class_id` INT NOT NULL,
  `class_name` VARCHAR(127) NULL,
  `class_type` VARCHAR(45) NULL,
  PRIMARY KEY (`class_id`))
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `drugeventdb`.`DrugClass`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `drugeventdb`.`DrugClass` (
  `drug_splset_id` CHAR(36) NOT NULL,
  `class_id` INT NULL,
  PRIMARY KEY (`drug_splset_id`, `class_id`),
  INDEX `dc_class_id_idx` (`class_id` ASC),
  CONSTRAINT `dc_drug_splset_id`
    FOREIGN KEY (`drug_splset_id`)
    REFERENCES `drugeventdb`.`Drug` (`drug_splset_id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION,
  CONSTRAINT `dc_class_id`
    FOREIGN KEY (`class_id`)
    REFERENCES `drugeventdb`.`Class` (`class_id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION)
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `drugeventdb`.`Label`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `drugeventdb`.`Label` (
  `label_id` VARCHAR(45) NOT NULL,
  `drug_splset_id` CHAR(36) NULL,
  `label_version` INT NULL,
  -- `label_effectiveTime` DATE NULL,
  PRIMARY KEY (`label_id`),
  INDEX `label_splsetID_idx` (`drug_splset_id` ASC),
  CONSTRAINT `label_splsetID`
    FOREIGN KEY (`drug_splset_id`)
    REFERENCES `drugeventdb`.`Drug` (`drug_splset_id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION)
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `drugeventdb`.`Patient`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `drugeventdb`.`Patient` (
  `patient_id` INT NOT NULL,
  `patient_age` VARCHAR(45) NULL,
  `patient_ageUnit` VARCHAR(45) NULL,
  `patient_sex` VARCHAR(45) NULL,
  -- `patient_deathDate` VARCHAR(45) NULL,
  PRIMARY KEY (`patient_id`))
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `drugeventdb`.`PatientDrug`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `drugeventdb`.`PatientDrug` (
  `drug_splset_id` CHAR(36) NULL,
  `patient_id` INT NOT NULL,
  `action` VARCHAR(45) NULL,
  `characterization` VARCHAR(45) NULL,
  `route_id` INT NULL,
  PRIMARY KEY (`drug_splset_id`, `patient_id`),
  INDEX `pd_routeID_idx` (`route_id` ASC),
  INDEX `pd_patientID_idx` (`patient_id` ASC),
  CONSTRAINT `pd_drug_splset_id`
    FOREIGN KEY (`drug_splset_id`)
    REFERENCES `drugeventdb`.`Drug` (`drug_splset_id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION,
  CONSTRAINT `pd_routeID`
    FOREIGN KEY (`route_id`)
    REFERENCES `drugeventdb`.`Route` (`route_id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION,
  CONSTRAINT `pd_patient_id`
    FOREIGN KEY (`patient_id`)
    REFERENCES `drugeventdb`.`Patient` (`patient_id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION)
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `drugeventdb`.`Term`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `drugeventdb`.`Term` (
  `term_id` INT NOT NULL,
  `term` VARCHAR(127) NULL,
  PRIMARY KEY (`term_id`))
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `drugeventdb`.`Reaction`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `drugeventdb`.`Reaction` (
  `patient_id` INT NOT NULL,
  `term_id` INT NOT NULL,
  `reaction_outcome` INT NOT NULL,
  PRIMARY KEY (`patient_id`, `term_id`),
  INDEX `reaction_termid_idx` (`term_id` ASC),
  CONSTRAINT `reaction_patientid`
    FOREIGN KEY (`patient_id`)
    REFERENCES `drugeventdb`.`Patient` (`patient_id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION,
  CONSTRAINT `reaction_termid`
    FOREIGN KEY (`term_id`)
    REFERENCES `drugeventdb`.`Term` (`term_id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION)
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `drugeventdb`.`Event`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `drugeventdb`.`Event` (
  `event_id` CHAR(8) NOT NULL,
  `event_reportVrs` INT NOT NULL,
  -- `event_receiveDate` DATE NULL,
  -- `event_receiptDate` DATE NULL,
  -- `event_transmissionDate` DATE NULL,
  `event_duplicate` TINYINT NULL,
  `event_serious` TINYINT NULL,
  `event_serDeath` TINYINT NULL,
  `event_serDisabiling` TINYINT NULL,
  `event_serHospitalization` TINYINT NULL,
  `event_serLifeThreatening` TINYINT NULL,
  `event_serOther` TINYINT NULL,
  -- `companyNumber` INT NULL,
  `country_id` INT NULL,
  `patient_id` INT NULL,
  PRIMARY KEY (`event_id`, `event_reportVrs`),
  INDEX `country_id_idx` (`country_id` ASC),
  INDEX `patient_id_idx` (`patient_id` ASC),
  CONSTRAINT `country_idID`
    FOREIGN KEY (`country_id`)
    REFERENCES `drugeventdb`.`Country` (`country_id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION,
  CONSTRAINT `patient_idID`
    FOREIGN KEY (`patient_id`)
    REFERENCES `drugeventdb`.`Patient` (`patient_id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION)
ENGINE = InnoDB;


SET SQL_MODE=@OLD_SQL_MODE;
SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS;
SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS;
