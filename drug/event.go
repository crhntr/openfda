package drug

import "time"

type Event struct {
	SafetyReportID      string `json:"_id,omitempty" bson:"_id,omitempty"`
	SafetyReportVersion string `json:"version,omitempty" bson:"version,omitempty"`

	ReceiveDate time.Time `json:"receiveDate,omitempty" bson:"receiveDate,omitempty"` // always yyyymmdd
	ReceiptDate time.Time `json:"receiptDate,omitempty" bson:"receiptDate,omitempty"` // always yyyymmdd

	Seriousness Seriousness `json:"seriousness,omitempty" bson:"seriousness,omitempty"`

	TransmissionDate time.Time `json:"traqnsmissionDate,omitempty" bson:"transmissionDate,omitempty"`

	Duplicate bool `json:"dup,omitempty" bson:"dup,omitempty"`

	CompanyNumber        string `json:"cNumb,omitempty" bson:"cNumb,omitempty"`
	OccurCountry         string `json:"occurcountry,omitempty" bson:"occurcountry,omitempty"`
	PrimarySourceCountry string `json:"primarysourcecountry,omitempty" bson:"primarysourcecountry,omitempty"`

	PrimarySource struct {
		Qualification   int8   `json:"qualification,omitempty" bson:"qualification,omitempty"`
		ReporterCountry string `json:"reportercountry,omitempty" bson:"reportercountry,omitempty"`
	} `json:"primarysource,omitempty" bson:"primarysource,omitempty"`

	Sender struct {
		SenderType         string `json:"type,omitempty" bson:"type,omitempty"`
		SenderOrganization string `json:"organization,omitempty" bson:"organization,omitempty"`
	} `json:"sender,omitempty" bson:"sender,omitempty"`

	ReportDuplicate struct {
		DuplicateSource string `json:"duplicatesource,omitempty" bson:"duplicatesource,omitempty"`
		DuplicateNumber string `json:"duplicateNumb,omitempty" bson:"duplicateNumb,omitempty"`
	} `json:"reportduplicate,omitempty" bson:"reportduplicate,omitempty"`

	Receiver struct {
		ReceiverType         string `json:"type,omitempty" bson:"type,omitempty"`
		ReceiverOrganization string `json:"org,omitempty" bson:"org,omitempty"`
	} `json:"receiver,omitempty" bson:"receiver,omitempty"`

	Patient struct {
		OnsetAge     string `json:"onsetAge,omitempty" bson:"onsetAge,omitempty"`
		OnsetAgeUnit string `json:"onsetAgeUnit,omitempty" bson:"onsetAgeUnit,omitempty"`
		Sex          string `json:"sex,omitempty" bson:"sex,omitempty"`
		Weight       string `json:"weight,omitempty" bson:"weight,omitempty"`

		Death time.Time `json:"deathDate,omitempty" bson:"deathDate,omitempty"`

		Drugs []struct {
			Characterization string `json:"characterization,omitempty" bson:"characterization,omitempty"`

			ActionDrug            string `json:"actionDrug,omitempty" bson:"actionDrug,omitempty"`
			RecurreAdministration string `json:"recurreAdministration,omitempty" bson:"recurreAdministration,omitempty"`

			AdministrationRoute string `json:"AdministrationRoute,omitempty" bson:"AdministrationRoute,omitempty"`

			Indication string `json:"indication,omitempty" bson:"indication,omitempty"`
			Additional string `json:"additional,omitempty" bson:"additional,omitempty"`

			DosageForm string `json:"dosageForm,omitempty" bson:"dosageForm,omitempty"`

			CumulativeDosageNumber string `json:"cumulativeDosageNumb,omitempty" bson:"cumulativeDosageNumb,omitempty"`
			CumulativeDosageUnit   string `json:"cumulativeDosageUnit,omitempty" bson:"cumulativeDosageUnit,omitempty"`

			IntervalDosageDefinition string `json:"intervalDosageDefinition,omitempty" bson:"intervalDosageDefinition,omitempty"`
			IntervalDosageUnitNumber string `json:"intervalDosageUnitNumb,omitempty" bson:"intervalDosageUnitNumb,omitempty"`

			SeparateDosageNumber  string `json:"separateDosageNumb,omitempty" bson:"separateDosageNumb,omitempty"`
			StructureDosageNumber string `json:"structureDosageNumb,omitempty" bson:"structureDosageNumb,omitempty"`
			StructureDosageUnit   string `json:"structureDosageUnit,omitempty" bson:"structureDosageUnit,omitempty"`

			AuthorizationNumber string `json:"authorizationNumb,omitempty" bson:"authorizationNumb,omitempty"`
			BatchNumber         string `json:"batchNumb,omitempty" bson:"batchNumb,omitempty"`
			DoseageText         string `json:"doseageText,omitempty" bson:"doseageText,omitempty"`

			EndDate       string `json:"endDate,omitempty" bson:"endDate,omitempty"`
			EndDateFormat string `json:"endDateFormat,omitempty" bson:"endDateFormat,omitempty"`

			StartDate       string `json:"startDate,omitempty" bson:"startDate,omitempty"`
			StartDateFormat string `json:"startDateFormat,omitempty" bson:"startDateFormat,omitempty"`

			TreatmentDuration     string `json:"treatmentDuration,omitempty" bson:"treatmentDuration,omitempty"`
			TreatmentDurationUnit string `json:"treatmentDurationUnit,omitempty" bson:"treatmentDurationUnit,omitempty"`

			OpenFDA string `json:"openfda,omitempty" bson:"openfda,omitempty"`
		} `json:"drug,omitempty" bson:"drug,omitempty"`

		Reactions []struct {
			MedDRA        string `json:"meddrapt,omitempty" bson:"meddrapt,omitempty"`
			MedDRAVersion string `json:"meddraversionpt,omitempty" bson:"meddraversionpt,omitempty"`
			Outcome       string `json:"outcome,omitempty" bson:"outcome,omitempty"`
		} `json:"reaction,omitempty" bson:"reaction,omitempty"`
	} `json:"patient,omitempty" bson:"patient,omitempty"`
}
