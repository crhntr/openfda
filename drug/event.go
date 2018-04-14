package drug

import "time"

type Event struct {
	FileName string `json:"fileName" bson:"fn"`
	FileHash string `json:"fileHash" bson:"fh"`

	SafetyReportID      string `json:"safetyreport,omitempty" bson:"rid,omitempty"`
	SafetyReportVersion string `json:"safetyreportversion,omitempty" bson:"rv,omitempty"`

	ReceiveDate      YearMonthDay `json:"receivedate,omitempty" bson:"rvd,omitempty"`
	ReceiptDate      YearMonthDay `json:"receiptdate,omitempty" bson:"rpd,omitempty"`
	TransmissionDate YearMonthDay `json:"transmissiondate,omitempty" bson:"td,omitempty"`

	Seriousness Seriousness `json:"seriousness,omitempty" bson:"s,omitempty"`

	Duplicate bool `json:"duplicate,omitempty" bson:"dup,omitempty"`

	CompanyNumber string `json:"companynumb,omitempty" bson:"cnm,omitempty"`
	OccurCountry  string `json:"occurcountry,omitempty" bson:"occ,omitempty"`

	PrimarySourceQualification   string `json:"primarySourceQualification,omitempty" bson:"pcsql,omitempty"`
	PrimarySourceReporterCountry string `json:"primarySourceReporterCountry,omitempty" bson:"psrc,omitempty"`

	ReportDuplicateSource string `json:"duplicateSource,omitempty" bson:"rds,omitempty"`
	ReportDuplicateNumber string `json:"duplicateNumber,omitempty" bson:"rdn,omitempty"`

	SenderType         string `json:"type,omitempty" bson:"st,omitempty"`
	SenderOrganization string `json:"org,omitempty" bson:"so,omitempty"`

	ReceiverType         string `json:"type,omitempty" bson:"rt,omitempty"`
	ReceiverOrganization string `json:"org,omitempty" bson:"ro,omitempty"`

	OnsetAge     string       `json:"patientOnsetAge,omitempty" bson:"oage,omitempty"`
	OnsetAgeUnit string       `json:"patientOnsetAgeUnit,omitempty" bson:"ount,omitempty"`
	Sex          string       `json:"patientSex,omitempty" bson:"sex,omitempty"`
	Weight       string       `json:"patientWeight,omitempty" bson:"wt,omitempty"`
	DeathDate    YearMonthDay `json:"patientDeathDate,omitempty" bson:"dthd,omitempty"`

	Drugs []struct {
		Characterization string `json:"characterization,omitempty" bson:"char,omitempty"`

		ActionDrug               string `json:"action,omitempty" bson:"act,omitempty"`
		Additional               string `json:"additional,omitempty" bson:"add,omitempty"`
		CumulativeDosageNumber   string `json:"cumulativeDosageNumber,omitempty" bson:"cn,omitempty"`
		CumulativeDosageUnit     string `json:"cumulativeDosageUnit,omitempty" bson:"cu,omitempty"`
		DosageForm               string `json:"dosageForm,omitempty" bson:"dfm,omitempty"`
		IntervalDosageDefinition string `json:"intervalDosageDefinition,omitempty" bson:"idd,omitempty"`
		IntervalDosageUnitNumber string `json:"intervalDosageUnitNumber,omitempty" bson:"idn,omitempty"`
		RecurreAdministration    string `json:"recurReadministration,omitempty" bson:"rea,omitempty"`
		SeparateDosageNumber     string `json:"separateDosageNumber,omitempty" bson:"sdn,omitempty"`
		StructureDosageNumber    string `json:"structureDosageNumb,omitempty" bson:"sedn,omitempty"`
		Doseagetext              string `json:"doseageText,omitempty" bson:"dtxt,omitempty"`
		StructureDosageUnit      string `json:"structureDosageUnit,omitempty" bson:"stdu,omitempty"`
		AdministrationRoute      string `json:"administrationRoute,omitempty" bson:"adrt,omitempty"`
		AuthorizationNumber      string `json:"authorizationNumber,omitempty" bson:"anum,omitempty"`
		BatchNumber              string `json:"batchNumber,omitempty" bson:"bnum,omitempty"`

		StartDate YearMonthDay `json:"startdate,omitempty" bson:"strt,omitempty"`
		EndDate   YearMonthDay `json:"enddate,omitempty" bson:"end,omitempty"`

		Indication string `json:"indication,omitempty" bson:"ind,omitempty"`

		TreatmentDuration time.Duration `json:"treatmentDuration,omitempty" bson:"dur,omitempty"`

		OpenFDA OpenFDA `json:"openfda,omitempty" bson:"openfda,omitempty"`
	} `json:"drugs,omitempty" bson:"dgs,omitempty"`

	Reactions []struct {
		MedDRA        string `json:"medDRA,omitempty" bson:"dra,omitempty"`
		MedDRAVersion string `json:"medDRAVersion,omitempty" bson:"drav,omitempty"`
		Outcome       string `json:"outcome,omitempty" bson:"oc,omitempty"`
	} `json:"reactions,omitempty" bson:"rns,omitempty"`
}


const (
	BSONKeyEventFileName = "fn"
	BSONKeyEventFileHash = "fh"

	BSONKeyEventSafetyReportID               = "rid"
	BSONKeyEventSafetyReportVersion          = "rv"
	BSONKeyEventReceiveDate                  = "rvd"
	BSONKeyEventReceiptDate                  = "rpd"
	BSONKeyEventTransmissionDate             = "td"
	BSONKeyEventSeriousness                  = "s"
	BSONKeyEventDuplicate                    = "dup"
	BSONKeyEventCompanyNumber                = "cnm"
	BSONKeyEventOccurCountry                 = "occ"
	BSONKeyEventPrimarySourceQualification   = "pcsql"
	BSONKeyEventPrimarySourceReporterCountry = "psrc"
	BSONKeyEventReportDuplicateSource        = "rds"
	BSONKeyEventReportDuplicateNumber        = "rdn"
	BSONKeyEventSenderType                   = "st"
	BSONKeyEventSenderOrganization           = "so"
	BSONKeyEventReceiverType                 = "rt"
	BSONKeyEventReceiverOrganization         = "ro"
	BSONKeyEventOnsetAge                     = "oage"
	BSONKeyEventOnsetAgeUnit                 = "ount"
	BSONKeyEventSex                          = "sex"
	BSONKeyEventWeight                       = "wt"
	BSONKeyEventDeathDate                    = "dthd"

	BSONKeyEventDrugs                    = "dgs"
	BSONKeyEventCharacterization         = "char"
	BSONKeyEventActionDrug               = "act"
	BSONKeyEventAdditional               = "add"
	BSONKeyEventCumulativeDosageNumber   = "cn"
	BSONKeyEventCumulativeDosageUnit     = "cu"
	BSONKeyEventDosageForm               = "dfm"
	BSONKeyEventIntervalDosageDefinition = "idd"
	BSONKeyEventIntervalDosageUnitNumber = "idn"
	BSONKeyEventRecurreAdministration    = "rea"
	BSONKeyEventSeparateDosageNumber     = "sdn"
	BSONKeyEventStructureDosageNumber    = "sedn"
	BSONKeyEventDoseagetext              = "dtxt"
	BSONKeyEventStructureDosageUnit      = "stdu"
	BSONKeyEventAdministrationRoute      = "adrt"
	BSONKeyEventAuthorizationNumber      = "anum"
	BSONKeyEventBatchNumber              = "bnum"
	BSONKeyEventStartDate                = "strt"
	BSONKeyEventEndDate                  = "end"
	BSONKeyEventIndication               = "ind"
	BSONKeyEventTreatmentDuration        = "dur"
	BSONKeyEventOpenFDA                  = "openfda"

	BSONKeyEventReactions     = "rns"
	BSONKeyEventMedDRA        = "dra"
	BSONKeyEventMedDRAVersion = "drav"
	BSONKeyEventOutcome       = "oc"
)
