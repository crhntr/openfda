package drug

type Event struct {
	SafetyReportID      string       `json:"safetyreport" bson:"safetyreport,omitempty"`
	SafetyReportVersion string       `json:"safetyreportversion,omitempty" bson:"safetyreportversion,omitempty"`
	ReceiveDate         YearMonthDay `json:"receivedate,omitempty" bson:"receivedate,omitempty"`

	ReceiptDate YearMonthDay `json:"receiptdate,omitempty" bson:"receiptdate,omitempty"`

	Serious                      string       `json:"serious,omitempty" bson:"serious,omitempty"`
	SeriousnessCongenitalAnomali string       `json:"seriousnesscongenitalanomali,omitempty" bson:"seriousnesscongenitalanomali,omitempty"`
	SeriousnessDeath             string       `json:"seriousnessdeath,omitempty" bson:"seriousnessdeath,omitempty"`
	SeriousnessDisabling         string       `json:"seriousnessdisabling,omitempty" bson:"seriousnessdisabling,omitempty"`
	SeriousnessHospitalization   string       `json:"seriousnesshospitalization,omitempty" bson:"seriousnesshospitalization,omitempty"`
	SeriousnessLifeThreatening   string       `json:"seriousnesslifethreatening,omitempty" bson:"seriousnesslifethreatening,omitempty"`
	SeriousnessOther             string       `json:"seriousnessother,omitempty" bson:"seriousnessother,omitempty"`
	TransmissionDate             YearMonthDay `json:"traqnsmissiondate,omitempty" bson:"transmissiondate,omitempty"`

	Duplicate            string `json:"duplicate,omitempty" bson:"duplicate,omitempty"`
	CompanyNumber        string `json:"companynumb,omitempty" bson:"companynumb,omitempty"`
	OccurCountry         string `json:"occurcountry,omitempty" bson:"occurcountry,omitempty"`
	PrimarySourceCountry string `json:"primarysourcecountry,omitempty" bson:"primarysourcecountry,omitempty"`
	PrimarySource        struct {
		Qualification   string `json:"qualification,omitempty" bson:"qualification,omitempty"`
		ReporterCountry string `json:"reportercountry,omitempty" bson:"reportercountry,omitempty"`
	} `json:"primarysource,omitempty" bson:"primarysource,omitempty"`
	ReportDuplicate struct {
		DuplicateSource string `json:"duplicatesource,omitempty" bson:"duplicatesource,omitempty"`
		DuplicateNumber string `json:"duplicatenumb,omitempty" bson:"duplicatenumb,omitempty"`
	} `json:"reportduplicate,omitempty" bson:"reportduplicate,omitempty"`
	Sender struct {
		SenderType         string `json:"sendertype,omitempty" bson:"sendertype,omitempty"`
		SenderOrganization string `json:"senderorganization,omitempty" bson:"senderorganization,omitempty"`
	} `json:"sender,omitempty" bson:"sender,omitempty"`
	Receiver struct {
		ReceiverType         string `json:"receivertype,omitempty" bson:"receivertype,omitempty"`
		ReceiverOrganization string `json:"receiverorganization,omitempty" bson:"receiverorganization,omitempty"`
	} `json:"receiver,omitempty" bson:"receiver,omitempty"`

	Patient struct {
		OnsetAge     string `json:"patientonsetage,omitempty" bson:"patientonsetage,omitempty"`
		OnsetAgeUnit string `json:"patientonsetageunit,omitempty" bson:"patientonsetageunit,omitempty"`
		Sex          string `json:"patientsex,omitempty" bson:"patientsex,omitempty"`
		Weight       string `json:"patientweight,omitempty" bson:"patientweight,omitempty"`
		Death        struct {
			Date YearMonthDay `json:"patientdeathdate,omitempty" bson:"patientdeathdate,omitempty"`
		} `json:"patientdeath,omitempty" bson:"patientdeath,omitempty"`

		Drugs []struct {
			ActionDrug               string `json:"actiondrug,omitempty" bson:"actiondrug,omitempty"`
			Additional               string `json:"drugadditional,omitempty" bson:"drugadditional,omitempty"`
			CumulativeDosageNumber   string `json:"drugcumulativedosagenumb,omitempty" bson:"drugcumulativedosagenumb,omitempty"`
			CumulativeDosageUnit     string `json:"drugcumulativedosageunit,omitempty" bson:"drugcumulativedosageunit,omitempty"`
			DosageForm               string `json:"drugdosageform,omitempty" bson:"drugdosageform,omitempty"`
			IntervalDosageDefinition string `json:"drugintervaldosagedefinition,omitempty" bson:"drugintervaldosagedefinition,omitempty"`
			IntervalDosageUnitNumber string `json:"drugintervaldosageunitnumb,omitempty" bson:"drugintervaldosageunitnumb,omitempty"`
			RecurreAdministration    string `json:"drugrecurreadministration,omitempty" bson:"drugrecurreadministration,omitempty"`
			SeparateDosageNumber     string `json:"drugseparatedosagenumb,omitempty" bson:"drugseparatedosagenumb,omitempty"`
			StructureDosageNumber    string `json:"drugstructuredosagenumb,omitempty" bson:"drugstructuredosagenumb,omitempty"`
			StructureDosageUnit      string `json:"drugstructuredosageunit,omitempty" bson:"drugstructuredosageunit,omitempty"`
			AdministrationRoute      string `json:"drugadministrationroute,omitempty" bson:"drugadministrationroute,omitempty"`
			AuthorizationNumber      string `json:"drugauthorizationnumb,omitempty" bson:"drugauthorizationnumb,omitempty"`
			BatchNumber              string `json:"drugbatchnumb,omitempty" bson:"drugbatchnumb,omitempty"`
			Characterization         string `json:"drugcharacterization,omitempty" bson:"drugcharacterization,omitempty"`
			Doseagetext              string `json:"drugdoseagetext,omitempty" bson:"drugdoseagetext,omitempty"`

			StartDate YearMonthDay `json:"drugstartdate,omitempty" bson:"drugstartdate,omitempty"`
			EndDate   YearMonthDay `json:"drugenddate,omitempty" bson:"drugenddate,omitempty"`

			Indication string `json:"drugindication,omitempty" bson:"drugindication,omitempty"`

			TreatmentDuration     string `json:"drugtreatmentduration,omitempty" bson:"drugtreatmentduration,omitempty"`
			TreatmentDurationUnit string `json:"drugtreatmentdurationunit,omitempty" bson:"drugtreatmentdurationunit,omitempty"`

			OpenFDA OpenFDA `json:"openfda,omitempty" bson:"openfda,omitempty"`
		} `json:"drug,omitempty" bson:"drug,omitempty"`
		Reactions []struct {
			MedDRA        string `json:"reactionmeddrapt,omitempty" bson:"reactionmeddrapt,omitempty"`
			MedDRAVersion string `json:"reactionmeddraversionpt,omitempty" bson:"reactionmeddraversionpt,omitempty"`
			Outcome       string `json:"reactionoutcome,omitempty" bson:"reactionoutcome,omitempty"`
		} `json:"reaction,omitempty" bson:"reaction,omitempty"`
	} `json:"patient,omitempty" bson:"patient,omitempty"`
}
