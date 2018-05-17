package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"
	"strconv"

	"github.com/crhntr/openfda/drug"
)

func insertDataToMySQL() {

	var (
		drugLabels []drug.Label
		drugEvents []drug.Event
	)

	fileInfos, err := ioutil.ReadDir(path.Join(*outPath, "drug", "event"))
	if err != nil {
		log.Fatalf("could not read drug/event directory: %q", err)
	}
	for _, info := range fileInfos {
		if info.IsDir() {
			continue
		}

		f, err := os.Open(path.Join(*outPath, "drug", "event", info.Name()))
		if err != nil {
			log.Printf("could not open file event %q: %q", info.Name(), err)
			continue
		}

		var event drug.Event
		if err := json.NewDecoder(f).Decode(&event); err != nil {
			log.Printf("error decoding event file %q: %q", info.Name(), err)
			continue
		}
		drugEvents = append(drugEvents, event)
	}

	fileInfos, err = ioutil.ReadDir(path.Join(*outPath, "drug", "label"))
	if err != nil {
		log.Fatalf("could not read drug/label directory: %q", err)
	}
	for _, info := range fileInfos {
		if info.IsDir() {
			continue
		}

		f, err := os.Open(path.Join(*outPath, "drug", "label", info.Name()))
		if err != nil {
			log.Printf("could not open file label %q: %q", info.Name(), err)
			continue
		}

		var label drug.Label
		if err := json.NewDecoder(f).Decode(&label); err != nil {
			log.Printf("error decoding label file %q: %q", info.Name(), err)
			continue
		}
		drugLabels = append(drugLabels, label)
	}

	fmt.Println(len(drugLabels), len(drugEvents))

	type Country struct {
		ID   int
		Name string
	}
	var countries []Country

	countriesMap := make(map[string]int)
	for _, event := range drugEvents {
		countriesMap[event.OccurCountry] = -1
	}

	for key, _ := range countriesMap {
		countriesMap[key] = len(countries)
		countries = append(countries, Country{
			ID: len(countries), Name: key,
		})
	}

	type Route struct {
		ID          int
		Description string
	}
	var routes []Route

	routesMap := make(map[string]int)
	for _, label := range drugLabels {
		for _, route := range label.OpenFDA.Route {
			routesMap[route] = -1
		}
	}

	for _, event := range drugEvents {
		for _, ed := range event.Drugs {
			routesMap[ed.AdministrationRoute] = -1
		}
	}

	for key, _ := range routesMap {
		routesMap[key] = len(routes)
		routes = append(routes, Route{
			ID: len(routes), Description: key,
		})
	}

	type Manufacturer struct {
		ID   int
		Name string
	}

	var manufacturers []Manufacturer

	const MaxManufacturerLen = 45

	manufacturersMap := make(map[string]bool)
	for _, label := range drugLabels {
		for _, mName := range label.OpenFDA.ManufacturerName {
			if len(mName) > MaxManufacturerLen {
				mName = mName[:MaxManufacturerLen]
			}
			manufacturersMap[mName] = true
		}
	}

	for key, _ := range manufacturersMap {
		manufacturers = append(manufacturers, Manufacturer{
			ID: len(manufacturers), Name: key,
		})
	}

	// type Ingredient struct {
	// 	ID   int
	// 	Name string
	// }
	//
	// var ingredients []Ingredient
	//
	// {
	// 	ingredientsMap := make(map[string]bool)
	// 	for _, label := range drugLabels {
	// 		for _, mName := range label.OpenFDA.SubstanceName {
	// 			ingredientsMap[mName] = true
	// 		}
	// 	}
	//
	// 	for key, _ := range ingredientsMap {
	// 		ingredients = append(ingredients, Ingredient{
	// 			ID: len(ingredients), Name: key,
	// 		})
	// 	}
	// }

	type Class struct {
		ID         int
		Name, Type string
	}

	var classes []Class

	classesMap := make(map[string]map[string]int)
	{
		classesMap["cs"] = make(map[string]int)
		classesMap["epc"] = make(map[string]int)
		classesMap["moa"] = make(map[string]int)
		classesMap["pe"] = make(map[string]int)

		for _, label := range drugLabels {
			for i, cls := range label.OpenFDA.PharmClassCS {
				cls = cls[:len(cls)-len(" [Chemical/Ingredient]")]
				label.OpenFDA.PharmClassCS[i] = cls
				classesMap["cs"][cls] = -1
			}

			for i, cls := range label.OpenFDA.PharmClassEPC {
				cls = cls[:len(cls)-len(" [EPC]")]
				label.OpenFDA.PharmClassEPC[i] = cls
				classesMap["epc"][cls] = -1
			}

			for i, cls := range label.OpenFDA.PharmClassMOA {
				cls = cls[:len(cls)-len(" [MoA]")]
				label.OpenFDA.PharmClassMOA[i] = cls
				classesMap["moa"][cls] = -1
			}

			for i, cls := range label.OpenFDA.PharmClassPE {
				cls = cls[:len(cls)-len(" [PE]")]
				label.OpenFDA.PharmClassPE[i] = cls
				classesMap["pe"][cls] = -1
			}
		}

		for classType, _ := range classesMap {
			for className, _ := range classesMap[classType] {
				classesMap[classType][className] = len(classes)
				classes = append(classes, Class{
					ID: len(classes), Name: className,
					Type: classType,
				})
			}
		}
	}

	type Drug struct {
		SPLSetID string
	}

	type NameType string
	const (
		NameTypeBrand     NameType = "Brand"
		NameTypeGeneric   NameType = "Generic"
		NameTypeSubstance NameType = "Substance"

		MaxNameLen = 127
	)

	type DrugName struct {
		ID             int
		SPLSetID, Name string
		Cuttoff        bool
		Type           NameType
	}

	type DrugManufacturer struct {
		SetID          string
		ManufacturerID int
	}

	type DrugClass struct {
		SetID   string
		ClassID int
	}

	type Label struct {
		LabelID      string
		SetID        string
		LabelVersion int
		// EffectiveTime time.Time
	}

	var (
		drugs             []Drug
		drugNames         []DrugName
		drugManufacturers []DrugManufacturer
		drugClasses       []DrugClass

		labels []Label
	)

	{
		drugsMap := make(map[string]bool)

		for _, label := range drugLabels {

			// skip labels already done
			if _, ok := drugsMap[label.SetID]; ok {
				continue
			}

			// Create DrugNames
			for _, name := range label.OpenFDA.BrandName {
				cutoff := len(name) > MaxNameLen
				if cutoff {
					name = name[:MaxNameLen]
				}

				drugNames = append(drugNames, DrugName{
					ID:       len(drugNames),
					SPLSetID: label.SetID, Name: name, Cuttoff: cutoff, Type: NameTypeBrand,
				})
			}
			for _, name := range label.OpenFDA.GenericName {
				cutoff := len(name) > MaxNameLen
				if cutoff {
					name = name[:MaxNameLen]
				}

				drugNames = append(drugNames, DrugName{
					ID:       len(drugNames),
					SPLSetID: label.SetID, Name: name, Cuttoff: cutoff, Type: NameTypeGeneric,
				})
			}
			for _, name := range label.OpenFDA.SubstanceName {
				cutoff := len(name) > MaxNameLen
				if cutoff {
					name = name[:MaxNameLen]
				}

				drugNames = append(drugNames, DrugName{
					ID:       len(drugNames),
					SPLSetID: label.SetID, Name: name, Cuttoff: cutoff, Type: NameTypeSubstance,
				})
			}

			drugsMap[label.SetID] = true

			for _, manName := range label.OpenFDA.ManufacturerName {
				for _, man := range manufacturers {
					if man.Name == manName {
						drugManufacturers = append(drugManufacturers, DrugManufacturer{
							SetID:          label.SetID,
							ManufacturerID: man.ID,
						})
						break
					}
				}
			}

			for _, cls := range label.OpenFDA.PharmClassCS {
				drugClasses = append(drugClasses, DrugClass{ClassID: classesMap["cs"][cls], SetID: label.SetID})
			}

			for _, cls := range label.OpenFDA.PharmClassEPC {
				drugClasses = append(drugClasses, DrugClass{ClassID: classesMap["epc"][cls], SetID: label.SetID})
			}

			for _, cls := range label.OpenFDA.PharmClassMOA {
				drugClasses = append(drugClasses, DrugClass{ClassID: classesMap["moa"][cls], SetID: label.SetID})
			}

			for _, cls := range label.OpenFDA.PharmClassPE {
				drugClasses = append(drugClasses, DrugClass{ClassID: classesMap["pe"][cls], SetID: label.SetID})
			}

			drugs = append(drugs, Drug{
				SPLSetID: label.SetID,
			})

			labels = append(labels, Label{
				SetID:        label.SetID,
				LabelID:      label.ID,
				LabelVersion: label.Version,
				// EffectiveTime: label.EffectiveTime
			})
		}
	}

	type Term struct {
		ID   int
		Term string
		// TermVersion int
	}
	var terms []Term

	termsMap := make(map[string]int)
	for _, e := range drugEvents {
		for _, r := range e.Reactions {
			termsMap[r.MedDRA] = -1
		}
	}

	for term, _ := range termsMap {
		termsMap[term] = len(terms)
		terms = append(terms, Term{
			ID:   len(terms),
			Term: term,
		})
	}

	type Patient struct {
		ID                           int
		Age, AgeUnit, Sex, DeathDate string
	}

	type PatientDrug struct {
		SetID            string
		PatientID        int
		Action           string
		Characterization string
		RouteID          int
	}

	type Reaction struct {
		PatientID int
		TermID    int
		Outcome   int
	}

	type Event struct {
		ID            string
		ReportVersion int

		// ReceiveDate, ReceiptDate, TransmissionDate time.Time

		Duplicate, Serious, SerDeath, SerDisabiling, SerHospitalization, SerLifeThreatening, SerOther bool

		// CompanyID int
		Country, Patient int
	}

	var (
		patients     []Patient
		patientDrugs []PatientDrug
		reactions    []Reaction
		events       []Event
	)

	for _, dEvent := range drugEvents {

		for _, r := range dEvent.Reactions {
			outcome, err := strconv.Atoi(r.Outcome)
			if err != nil {
				outcome = 6
			}

			reactions = append(reactions, Reaction{
				PatientID: len(patients),
				TermID:    termsMap[r.MedDRA],
				Outcome:   outcome,
			})
		}
		version, err := strconv.Atoi(dEvent.SafetyReportVersion)
		if err != nil {
			version = 1
		}

		for _, deDrug := range dEvent.Drugs {
			var drugID string

			for _, drug := range drugs {
				for _, id := range deDrug.OpenFDA.SPLSetID {
					if id == drug.SPLSetID {
						drugID = drug.SPLSetID
						break
					}
				}
			}
			if drugID == "" {
				continue
			}

			patientDrugs = append(patientDrugs, PatientDrug{
				SetID:            drugID,
				PatientID:        len(patients),
				Action:           deDrug.ActionDrug,
				Characterization: deDrug.Characterization,
				RouteID:          routesMap[deDrug.AdministrationRoute],
			})
		}

		events = append(events, Event{
			ID:                 dEvent.SafetyReportID,
			ReportVersion:      version,
			Duplicate:          dEvent.Duplicate,
			Serious:            (dEvent.Seriousness & drug.Serious) != 0,
			SerDeath:           (dEvent.Seriousness & drug.Death) != 0,
			SerDisabiling:      (dEvent.Seriousness & drug.Disabling) != 0,
			SerHospitalization: (dEvent.Seriousness & drug.Hospitalization) != 0,
			SerLifeThreatening: (dEvent.Seriousness & drug.LifeThreatening) != 0,
			SerOther:           (dEvent.Seriousness & drug.Other) != 0,
			Country:            countriesMap[dEvent.OccurCountry],
			Patient:            len(patients),
		})

		patients = append(patients, Patient{
			ID:  len(patients),
			Age: dEvent.OnsetAge, AgeUnit: dEvent.OnsetAgeUnit,
			Sex: dEvent.Sex,
		})

	}

	{
		dedupMap := make([]map[string]bool, len(patients))
		for i := range dedupMap {
			dedupMap[i] = make(map[string]bool)
		}

		pds := patientDrugs[:0]

		for _, pd := range patientDrugs {
			if _, ok := dedupMap[pd.PatientID][pd.SetID]; !ok {
				pds = append(pds, pd)
			}
			dedupMap[pd.PatientID][pd.SetID] = true
		}

		patientDrugs = pds
	}

	endValue := func(i, l int) {
		if i != l-1 {
			fmt.Print(",\n")
		} else {
			fmt.Print(";\n\n\n")
		}
	}

	fmt.Print("\n-- countries\n\n")
	fmt.Println("INSERT INTO Country (country_id, country_name)\nVALUES ")
	for i, item := range countries {
		fmt.Printf("\t(%d, '%s')", item.ID, item.Name)
		endValue(i, len(countries))
	}

	fmt.Print("\n-- routes\n\n")
	fmt.Println("INSERT INTO Route (route_id, route_description)\nVALUES ")
	for i, item := range routes {
		fmt.Printf("\t(%d, '%s')", item.ID, item.Description)
		endValue(i, len(routes))
	}

	fmt.Print("\n-- manufacturers\n\n")
	fmt.Println("INSERT INTO Manufacturer (manufacturer_id, manufacturer_name)\nVALUES ")
	for i, item := range manufacturers {
		fmt.Printf("\t(%d, '%s')", item.ID, item.Name)
		endValue(i, len(manufacturers))
	}

	fmt.Print("\n-- classes\n\n")
	fmt.Println("INSERT INTO Class (class_id, class_name, class_type)\nVALUES ")
	for i, item := range classes {
		fmt.Printf("\t(%d, '%s', '%s')", item.ID, item.Name, item.Type)
		endValue(i, len(classes))
	}

	fmt.Print("\n-- drugs\n\n")
	fmt.Println("INSERT INTO Drug (drug_splset_id)\nVALUES ")
	for i, item := range drugs {
		fmt.Printf("\t('%s')", item.SPLSetID)
		endValue(i, len(drugs))
	}

	fmt.Print("\n-- drugNames\n\n")
	fmt.Println("INSERT INTO DrugName (drug_name_id, drug_name_splset_id, drug_name_name, drug_name_cutoff, drug_name_type)\nVALUES ")
	for i, item := range drugNames {
		fmt.Printf("\t(%d, '%s', '%s', %d, '%s')", item.ID, item.SPLSetID, item.Name, ToTINYINT(item.Cuttoff), item.Type)
		endValue(i, len(drugNames))
	}

	fmt.Print("\n-- drugManufacturers\n\n")
	fmt.Println("INSERT INTO DrugManufacturer (drug_splset_id, manufacturer_id)\nVALUES ")
	for i, item := range drugManufacturers {
		fmt.Printf("\t('%s', %d)", item.SetID, item.ManufacturerID)
		endValue(i, len(drugManufacturers))
	}

	fmt.Print("\n-- drugClasses\n\n")
	fmt.Println("INSERT INTO DrugClass (drug_splset_id, class_id)\nVALUES ")
	for i, item := range drugClasses {
		fmt.Printf("\t('%s', %d)", item.SetID, item.ClassID)
		endValue(i, len(drugClasses))
	}

	fmt.Print("\n-- labels\n\n")
	fmt.Println("INSERT INTO Label (label_id, drug_splset_id, label_version)\nVALUES ")
	for i, item := range labels {
		fmt.Printf("\t('%s', '%s', %d)", item.LabelID, item.SetID, item.LabelVersion)
		endValue(i, len(labels))
	}

	fmt.Print("\n-- patients\n\n")
	fmt.Println("INSERT INTO Patient (patient_id, patient_age, patient_ageUnit, patient_sex)\nVALUES ")
	for i, item := range patients {
		fmt.Printf("\t(%d, '%s', '%s', '%s')", item.ID, item.Age, item.AgeUnit, item.Sex)
		endValue(i, len(patients))
	}

	fmt.Print("\n-- patientDrugs\n\n")
	fmt.Println("INSERT INTO PatientDrug (drug_splset_id, patient_id, action, characterization, route_id)\nVALUES ")
	for i, item := range patientDrugs {
		fmt.Printf("\t('%s', %d, '%s', '%s', %d)", item.SetID, item.PatientID, item.Action, item.Characterization, item.RouteID)
		endValue(i, len(patientDrugs))
	}

	fmt.Print("\n-- terms\n\n")
	fmt.Println("INSERT INTO Term (term_id, term)\nVALUES ")
	for i, item := range terms {
		fmt.Printf("\t(%d, '%s')", item.ID, item.Term)
		endValue(i, len(terms))
	}

	fmt.Print("\n-- reactions\n\n")
	fmt.Println("INSERT INTO Reaction (patient_id, term_id, reaction_outcome)\nVALUES ")
	for i, item := range reactions {
		fmt.Printf("\t(%d, %d, %d)", item.PatientID, item.TermID, item.Outcome)
		endValue(i, len(reactions))
	}

	fmt.Print("\n-- events\n\n")
	fmt.Println("INSERT INTO Event (event_id, event_reportVrs, event_duplicate, event_serious, event_serDeath, event_serDisabiling, event_serHospitalization, event_serLifeThreatening, event_serOther, country_id, patient_id)\nVALUES ")
	for i, item := range events {
		fmt.Printf("\t('%s', %d, %d, %d, %d, %d, %d, %d, %d, %d, %d)",
			item.ID,
			item.ReportVersion,
			ToTINYINT(item.Duplicate),
			ToTINYINT(item.Serious),
			ToTINYINT(item.SerDeath),
			ToTINYINT(item.SerDisabiling),
			ToTINYINT(item.SerHospitalization),
			ToTINYINT(item.SerLifeThreatening),
			ToTINYINT(item.SerOther),
			item.Country,
			item.Patient,
		)
		endValue(i, len(events))
	}

	//
	// db, err := getMySQLSession()
	// if err != nil {
	// 	log.Fatalf("error could not get db Seesion: %s", err)
	// }
}
