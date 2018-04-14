package drug

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


type RawEvent struct {
	SafetyReportID      string       `json:"safetyreport,omitempty"`
	SafetyReportVersion string       `json:"safetyreportversion,omitempty"`
	ReceiveDate         YearMonthDay `json:"receivedate,omitempty"`

	ReceiptDate YearMonthDay `json:"receiptdate,omitempty"`

	Serious                      string       `json:"serious,omitempty"`
	SeriousnessCongenitalAnomali string       `json:"seriousnesscongenitalanomali,omitempty"`
	SeriousnessDeath             string       `json:"seriousnessdeath,omitempty"`
	SeriousnessDisabling         string       `json:"seriousnessdisabling,omitempty"`
	SeriousnessHospitalization   string       `json:"seriousnesshospitalization,omitempty"`
	SeriousnessLifeThreatening   string       `json:"seriousnesslifethreatening,omitempty"`
	SeriousnessOther             string       `json:"seriousnessother,omitempty"`
	TransmissionDate             YearMonthDay `json:"traqnsmissiondate,omitempty"`

	Duplicate            string `json:"duplicate,omitempty"`
	CompanyNumber        string `json:"companynumb,omitempty"`
	OccurCountry         string `json:"occurcountry,omitempty"`
	PrimarySourceCountry string `json:"primarysourcecountry,omitempty"`
	PrimarySource        struct {
		Qualification   string `json:"qualification,omitempty"`
		ReporterCountry string `json:"reportercountry,omitempty"`
	} `json:"primarysource,omitempty"`
	ReportDuplicate struct {
		DuplicateSource string `json:"duplicatesource,omitempty"`
		DuplicateNumber string `json:"duplicatenumb,omitempty"`
	} `json:"reportduplicate,omitempty"`
	Sender struct {
		SenderType         string `json:"sendertype,omitempty"`
		SenderOrganization string `json:"senderorganization,omitempty"`
	} `json:"sender,omitempty"`
	Receiver struct {
		ReceiverType         string `json:"receivertype,omitempty"`
		ReceiverOrganization string `json:"receiverorganization,omitempty"`
	} `json:"receiver,omitempty"`

	Patient struct {
		OnsetAge     string `json:"patientonsetage,omitempty"`
		OnsetAgeUnit string `json:"patientonsetageunit,omitempty"`
		Sex          string `json:"patientsex,omitempty"`
		Weight       string `json:"patientweight,omitempty"`
		Death        struct {
			Date YearMonthDay `json:"patientdeathdate,omitempty"`
		} `json:"patientdeath,omitempty"`

		Drugs []struct {
			ActionDrug               string `json:"actiondrug,omitempty"`
			Additional               string `json:"drugadditional,omitempty"`
			CumulativeDosageNumber   string `json:"drugcumulativedosagenumb,omitempty"`
			CumulativeDosageUnit     string `json:"drugcumulativedosageunit,omitempty"`
			DosageForm               string `json:"drugdosageform,omitempty"`
			IntervalDosageDefinition string `json:"drugintervaldosagedefinition,omitempty"`
			IntervalDosageUnitNumber string `json:"drugintervaldosageunitnumb,omitempty"`
			RecurreAdministration    string `json:"drugrecurreadministration,omitempty"`
			SeparateDosageNumber     string `json:"drugseparatedosagenumb,omitempty"`
			StructureDosageNumber    string `json:"drugstructuredosagenumb,omitempty"`
			StructureDosageUnit      string `json:"drugstructuredosageunit,omitempty"`
			AdministrationRoute      string `json:"drugadministrationroute,omitempty"`
			AuthorizationNumber      string `json:"drugauthorizationnumb,omitempty"`
			BatchNumber              string `json:"drugbatchnumb,omitempty"`
			Characterization         string `json:"drugcharacterization,omitempty"`
			Doseagetext              string `json:"drugdoseagetext,omitempty"`

			StartDate YearMonthDay `json:"drugstartdate,omitempty"`
			EndDate   YearMonthDay `json:"drugenddate,omitempty"`

			Indication string `json:"drugindication,omitempty"`

			TreatmentDuration     string `json:"drugtreatmentduration,omitempty"`
			TreatmentDurationUnit string `json:"drugtreatmentdurationunit,omitempty"`

			OpenFDA OpenFDA `json:"openfda,omitempty"`
		} `json:"drug,omitempty"`
		Reactions []struct {
			MedDRA        string `json:"reactionmeddrapt,omitempty"`
			MedDRAVersion string `json:"reactionmeddraversionpt,omitempty"`
			Outcome       string `json:"reactionoutcome,omitempty"`
		} `json:"reaction,omitempty"`
	} `json:"patient,omitempty"`
}
