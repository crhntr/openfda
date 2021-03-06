--COMP420 Final Project
--Drug Event DB Views & Stored Procedures

----------VIEWS----------
CREATE VIEW `EventCountByNDC` AS
SELECT ndc, count(ndc) AS 'Total Events', count(event_serious) AS 'Serious', 
count(event_serDeath) AS 'Deaths', count(event_serDisabiling) AS 'Disabiling',
count(event_serLifeThreatening) AS 'Life Threatening', count(event_serHospitalization) AS 'Hospitalization',
count(event_serOther) AS 'Other' FROM Event
GROUP BY ndc ORDER BY count(ndc) DESC;

CREATE VIEW `EventCountByClass` AS
SELECT Class.class_name AS 'Drug Class', count(class_id) AS 'Total Events', count(event_serious) AS 'Serious', 
count(event_serDeath) AS 'Deaths', count(event_serDisabiling) AS 'Disabiling',
count(event_serLifeThreatening) AS 'Life Threatening', count(event_serHospitalization) AS 'Hospitalization',
count(event_serOther) AS 'Other' FROM Event
JOIN Drug ON Event.ndc = Drug.ndc
JOIN DrugClass ON Drug.ndc = DrugClass.ndc
JOIN Class ON DrugClass.class_id = Class.class_id
GROUP BY Class.class_id ORDER BY count(class_id) DESC;

CREATE VIEW `FrequentReactions` AS
SELECT Term.term AS 'Reaction', count(patient_id) AS 'Number of Patients' FROM Reaction
JOIN Term ON Term.term_id = Reaction.term_id
GROUP BY term_id
ORDER BY count(patient_id) DESC;

CREATE VIEW `PatientsDrugs` AS
SELECT Patient.*, Drug.ndc, Drug.drug_brandname, Drug.drug_genericname, PatientDrug.action, PatientDrug.characterization, Route.route_description FROM PatientDrug
JOIN Patient ON PatientDrug.patient_id = Patient.patient_id
JOIN Drug ON PatientDrug.ndc = Drug.ndc
JOIN Route ON PatientDrug.route_id = Route.route_id;

CREATE VIEW `DrugList` AS
SELECT Drug.ndc AS 'NDC', Class.class_name AS 'Class', Manufacturer.manufacturer_name AS 'Drug Mfg', drug_brandname AS 'Brand Name', drug_genericname AS 'Generic Name' FROM Drug
JOIN DrugClass ON Drug.ndc = DrugClass.ndc
JOIN Class ON DrugClass.class_id = Class.class_id
JOIN Manufacturer ON Drug.drug_manufacturer = Manufacturer.manufacturer_id;

----------STORED PROCEDURES----------

DELIMITER //
CREATE PROCEDURE `countryEvents` (IN countryData VARCHAR(45))
BEGIN
SELECT Event.* FROM Event
JOIN Country ON Country.country_id = Event.country_id
WHERE countryData = Event.country_id OR countryData = Country.country_name;
END //
DELIMITER ;

DELIMITER //
CREATE PROCEDURE `eventsByAge` (IN age INT, IN ageUnits VARCHAR(45))
BEGIN
SELECT Patient.*, Event.* FROM Event
JOIN Patient ON Patient.patient_id = Event.event_patient
WHERE Patient.age = age AND Patient.ageUnits = ageUnits;
END //
DELIMITER ;

DELIMITER //
CREATE PROCEDURE `getEventsByAgeRange` (IN ageStart INT, IN ageEnd INT, IN ageUnits VARCHAR(45))
BEGIN
SELECT Patient.*, Event.* FROM Event
JOIN Patient ON Patient.patient_id = Event.event_patient
WHERE Patient.ageUnits = ageUnits AND Patient.age BETWEEN ageStart AND ageEnd;
END //
DELIMITER ;

DELIMITER //
CREATE PROCEDURE `getDrugEventCount` (IN drugName VARCHAR(45))
BEGIN
SELECT EventCountByNDC.* FROM EventCountByNDC
JOIN Drug ON Drug.ndc = EventCountByNDC.ndc
WHERE Drug.drug_genericname = drugName OR Drug.drug_brandname = drugName;
END //
DELIMITER ;

DELIMITER //
CREATE PROCEDURE `getIngredients` (IN drugData VARCHAR(45))
BEGIN
SELECT Ingredient.ingredient_name FROM DrugIngredient
JOIN Ingredient ON Ingredient.ingredient_id = DrugIngredient.ingredient_id; 
END //
DELIMITER ;

DELIMITER //
CREATE PROCEDURE `reactionCount` (IN reaction VARCHAR(45))
BEGIN
    SELECT count(patient_id) AS count FROM Reaction
	JOIN Term ON Term.term_id = Reaction.term_id
	GROUP BY term_id
	WHERE reaction LIKE Term.term;
END //
DELIMITER ;

DELIMITER //
CREATE PROCEDURE `patientReactions` (IN patient INT)
BEGIN
    SELECT Term.term AS 'Reaction', Reaction.reaction_outcome AS 'Outcome' FROM Reaction
	JOIN Term ON Term.term_id = Reaction.term_id
	WHERE patient = Reaction.patient_id;
END //
DELIMITER ;

