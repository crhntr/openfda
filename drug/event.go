package drug

import ("time"
"strconv")

type Event struct {
	FileName string `json:"fileName" bson:"fn"`

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

	Drugs     []EventDrug     `json:"drugs,omitempty" bson:"dgs,omitempty"`
	Reactions []EventReaction `json:"reactions,omitempty" bson:"rns,omitempty"`
}

type EventDrug struct {
	Characterization string `json:"characterization,omitempty" bson:"char,omitempty"`
	MedicinalProduct string `json:"medicinalproduct,omitempty" bson:"mp,omitempty"`

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

	TreatmentDurationValue time.Duration `json:"treatmentDuration,omitempty" bson:"dur,omitempty"`
	TreatmentDurationNumber string `json:"drugtreatmentduration,omitempty" bson:"drugtreatmentduration,omitempty"`
	TreatmentDurationUnit string `json:"drugtreatmentdurationunit,omitempty" bson:"drugtreatmentdurationunit,omitempty"`

	OpenFDA OpenFDA `json:"openfda,omitempty" bson:"openfda,omitempty"`
}

type EventReaction struct {
	MedDRA        string `json:"medDRA,omitempty" bson:"dra,omitempty"`
	MedDRAVersion string `json:"medDRAVersion,omitempty" bson:"drav,omitempty"`
	Outcome       string `json:"outcome,omitempty" bson:"oc,omitempty"`
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
	SafetyReportID      string       `json:"safetyreport,omitempty" bson:"safetyreport,omitempty"`
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
	TransmissionDate             YearMonthDay `json:"traqnsmissiondate,omitempty" bson:"traqnsmissiondate,omitempty"`

	Duplicate            string `json:"duplicate,omitempty" bson:"duplicate,omitempty"`
	CompanyNumber        string `json:"companynumb,omitempty" bson:"companynumb,omitempty"`
	OccurCountry         string `json:"occurcountry,omitempty" bson:"occurcountry,omitempty"`
	PrimarySourceCountry string `json:"primarysourcecountry,omitempty" bson:"primarysourcecountry,omitempty"`

	PrimarySource struct {
		Qualification   string `json:"qualification,omitempty" bson:"qualification,omitempty"`
		ReporterCountry string `json:"reportercountry,omitempty" bson:"reportercountry,omitempty"`
	} `json:"primarysource,omitempty" bson:"primarysource,omitempty"`
	ReportDuplicate struct {
		Source string `json:"duplicatesource,omitempty" bson:"duplicatesource,omitempty"`
		Number string `json:"duplicatenumb,omitempty" bson:"duplicatenumb,omitempty"`
	} `json:"reportduplicate,omitempty" bson:"reportduplicate,omitempty"`
	Sender struct {
		Type         string `json:"sendertype,omitempty" bson:"sendertype,omitempty"`
		Organization string `json:"senderorganization,omitempty" bson:"senderorganization,omitempty"`
	} `json:"sender,omitempty" bson:"sender,omitempty"`
	Receiver struct {
		Type         string `json:"receivertype,omitempty" bson:"receivertype,omitempty"`
		Organization string `json:"receiverorganization,omitempty" bson:"receiverorganization,omitempty"`
	} `json:"receiver,omitempty" bson:"receiver,omitempty"`

	Patient struct {
		OnsetAge     string `json:"patientonsetage,omitempty" bson:"patientonsetage,omitempty"`
		OnsetAgeUnit string `json:"patientonsetageunit,omitempty" bson:"patientonsetageunit,omitempty"`
		Sex          string `json:"patientsex,omitempty" bson:"patientsex,omitempty"`
		Weight       string `json:"patientweight,omitempty" bson:"patientweight,omitempty"`
		Death        struct {
			Date YearMonthDay `json:"patientdeathdate,omitempty" bson:"patientdeathdate,omitempty"`
		} `json:"patientdeath,omitempty" bson:"patientdeath,omitempty"`

		Drugs     []RawEventDrug     `json:"drug,omitempty" bson:"drug,omitempty"`
		Reactions []RawEventReaction `json:"reaction,omitempty" bson:"reaction,omitempty"`
	} `json:"patient,omitempty" bson:"patient,omitempty"`
}

type RawEventDrug struct {
	MedicinalProduct string `json:"medicinalproduct,omitempty" bson:"medicinalproduct,omitempty"`

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

	TreatmentDurationNumber     string `json:"drugtreatmentduration,omitempty" bson:"drugtreatmentduration,omitempty"`
	TreatmentDurationUnit string `json:"drugtreatmentdurationunit,omitempty" bson:"drugtreatmentdurationunit,omitempty"`

	OpenFDA OpenFDA `json:"openfda,omitempty" bson:"openfda,omitempty"`
}

type RawEventReaction struct {
	MedDRA        string `json:"reactionmeddrapt,omitempty" bson:"reactionmeddrapt,omitempty"`
	MedDRAVersion string `json:"reactionmeddraversionpt,omitempty" bson:"reactionmeddraversionpt,omitempty"`
	Outcome       string `json:"reactionoutcome,omitempty" bson:"reactionoutcome,omitempty"`
}

func (rw RawEvent) Event() (Event, []OpenFDA) {
	var drugData []OpenFDA
	var event = Event{
		SafetyReportID:      rw.SafetyReportID,
		SafetyReportVersion: rw.SafetyReportVersion,
		ReceiveDate:         rw.ReceiveDate,
		ReceiptDate:         rw.ReceiptDate,
		TransmissionDate:    rw.TransmissionDate,
		Duplicate:           rw.Duplicate == "1",
		CompanyNumber:       rw.CompanyNumber,
		OccurCountry:        rw.OccurCountry,

		PrimarySourceQualification:   rw.PrimarySource.Qualification,
		PrimarySourceReporterCountry: rw.PrimarySource.ReporterCountry,

		ReportDuplicateSource: rw.ReportDuplicate.Source,
		ReportDuplicateNumber: rw.ReportDuplicate.Number,

		SenderType:         rw.Sender.Type,
		SenderOrganization: rw.Sender.Organization,

		ReceiverType:         rw.Receiver.Type,
		ReceiverOrganization: rw.Receiver.Organization,

		OnsetAge:     rw.Patient.OnsetAge,
		OnsetAgeUnit: rw.Patient.OnsetAgeUnit,
		Sex:          rw.Patient.Sex,
		Weight:       rw.Patient.Weight,
		DeathDate:    rw.Patient.Death.Date,
	}

	event.Seriousness = ParseSeriousness(
		rw.Serious,
		rw.SeriousnessCongenitalAnomali,
		rw.SeriousnessDeath,
		rw.SeriousnessDisabling,
		rw.SeriousnessHospitalization,
		rw.SeriousnessLifeThreatening,
		rw.SeriousnessOther)

	event.Drugs = make([]EventDrug, 0, len(rw.Patient.Drugs))
	event.Reactions = make([]EventReaction, 0, len(rw.Patient.Reactions))

	for _, d := range rw.Patient.Drugs {
		if d.Indication == "PRODUCT USED FOR UNKNOWN INDICATION" {
			d.Indication = ""
		}

		ed := EventDrug{
			MedicinalProduct:         d.MedicinalProduct,
			ActionDrug:               d.ActionDrug,
			Additional:               d.Additional,
			CumulativeDosageNumber:   d.CumulativeDosageNumber,
			CumulativeDosageUnit:     d.CumulativeDosageUnit,
			DosageForm:               d.DosageForm,
			IntervalDosageDefinition: d.IntervalDosageDefinition,
			IntervalDosageUnitNumber: d.IntervalDosageUnitNumber,
			RecurreAdministration:    d.RecurreAdministration,
			SeparateDosageNumber:     d.SeparateDosageNumber,
			StructureDosageNumber:    d.StructureDosageNumber,
			StructureDosageUnit:      d.StructureDosageUnit,
			AdministrationRoute:      d.AdministrationRoute,
			AuthorizationNumber:      d.AuthorizationNumber,
			BatchNumber:              d.BatchNumber,
			Characterization:         d.Characterization,
			Doseagetext:              d.Doseagetext,
			StartDate:                d.StartDate,
			EndDate:                  d.EndDate,
			TreatmentDurationNumber: d.TreatmentDurationNumber,
		TreatmentDurationUnit: d.TreatmentDurationUnit,
			TreatmentDurationValue: ParseDuration(d.TreatmentDurationUnit, d.TreatmentDurationNumber),

			Indication:               d.Indication,
			OpenFDA: d.OpenFDA,
		}
		if len(d.OpenFDA.SPLSetID) > 0 {
			d.OpenFDA.SetID = d.OpenFDA.SPLSetID[0]
		}
		if len(d.OpenFDA.SPLID) > 0 {
			d.OpenFDA.ID = d.OpenFDA.SPLID[0]
		}
		drugData = append(drugData, d.OpenFDA)

		event.Drugs = append(event.Drugs, ed)
	}

	for _, r := range rw.Patient.Reactions {
		event.Reactions = append(event.Reactions, EventReaction{
			MedDRA:        r.MedDRA,
			MedDRAVersion: r.MedDRAVersion,
			Outcome:       r.Outcome,
		})
	}

	return event, drugData
}

type EventOpenFDA struct {
	ApplicationNumber []string `json:"application_number,omitempty" bson:"application_number,omitempty"`
	BrandName         []string `json:"brand_name,omitempty" bson:"brand_name,omitempty"`
	GenericName       []string `json:"generic_name,omitempty" bson:"generic_name,omitempty"`
	ManufacturerName  []string `json:"manufacturer_name,omitempty" bson:"manufacturer_name,omitempty"`
	NUI               []string `json:"nui,omitempty" bson:"nui,omitempty"`
	PackageNDC        []string `json:"package_ndc,omitempty" bson:"package_ndc,omitempty"`
	PharmClassCS      []string `json:"pharm_class_cs,omitempty" bson:"pharm_class_cs,omitempty"`
	PharmClassEPC     []string `json:"pharm_class_epc,omitempty" bson:"pharm_class_epc,omitempty"`
	PharmClassMOA     []string `json:"pharm_class_moa,omitempty" bson:"pharm_class_moa,omitempty"`
	PharmClassPE      []string `json:"pharm_class_pe,omitempty" bson:"pharm_class_pe,omitempty"`
	ProductNDC        []string `json:"product_ndc,omitempty" bson:"product_ndc,omitempty"`
	ProductType       []string `json:"product_type,omitempty" bson:"product_type,omitempty"`
	Route             []string `json:"route,omitempty" bson:"route,omitempty"`
	RxCUI             []string `json:"rxcui,omitempty" bson:"rxcui,omitempty"`
	SPLID             []string `json:"spl_id,omitempty" bson:"spl_id,omitempty"`
	SPLSetID          []string `json:"spl_set_id,omitempty" bson:"spl_set_id,omitempty"`
	SubstanceName     []string `json:"substance_name,omitempty" bson:"substance_name,omitempty"`
	UNII              []string `json:"unii,omitempty" bson:"unii,omitempty"`
	UPC               []string `json:"upc,omitempty" bson:"upc,omitempty"`

	// Undocumented Fields Found In Responses
	IsOriginalPackager []bool `json:"is_original_packager,omitempty" bson:"is_original_packager,omitempty"`
}

func ParseDuration(unit, number string) time.Duration {
	n, err := strconv.Atoi(number)
	if err != nil {
		return time.Duration(0)
	}
	nd := time.Duration(n)
	switch unit {
	case "801": // Year
		return time.Duration(364.25 * 24 * time.Hour * nd )
	case "802": // Month
		return time.Duration(30 * 24 * time.Hour * nd)
	case "803": // Week
		return time.Duration(30 * 24 * time.Hour * nd)
	case "804": // Day
		return time.Duration(24 * time.Hour * nd)
	case "805": // Hour
		return time.Duration(time.Hour * nd)
	case "806": // Minute
		return time.Duration(time.Minute * nd)
	default:
		return time.Duration(0)
	}
}
