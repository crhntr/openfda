package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"
	"time"

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

	countriesMap := make(map[string]bool)
	for _, event := range drugEvents {
		if event.OccurCountry != "" {
			countriesMap[event.OccurCountry] = true
		}
	}

	for key, _ := range countriesMap {
		countries = append(countries, Country{
			ID: len(countries), Name: key,
		})
	}

	type Route struct {
		ID          int
		Description string
	}
	var routes []Route

	routesMap := make(map[string]bool)
	for _, label := range drugLabels {
		for _, route := range label.OpenFDA.Route {
			routesMap[route] = true
		}
	}

	for key, _ := range routesMap {
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

	{
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
	}

	type Ingredient struct {
		ID   int
		Name string
	}

	var ingredients []Ingredient

	{
		ingredientsMap := make(map[string]bool)
		for _, label := range drugLabels {
			for _, mName := range label.OpenFDA.SubstanceName {
				ingredientsMap[mName] = true
			}
		}

		for key, _ := range ingredientsMap {
			ingredients = append(ingredients, Ingredient{
				ID: len(ingredients), Name: key,
			})
		}
	}

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
			for _, cls := range label.OpenFDA.PharmClassCS {
				classesMap["cs"][cls[:len(cls)-len(" [Chemical/Ingredient]")]] = -1
			}

			for _, cls := range label.OpenFDA.PharmClassEPC {
				classesMap["epc"][cls[:len(cls)-len(" [EPC]")]] = -1
			}

			for _, cls := range label.OpenFDA.PharmClassMOA {
				classesMap["moa"][cls[:len(cls)-len(" [MoA]")]] = -1
			}

			for _, cls := range label.OpenFDA.PharmClassPE {
				classesMap["pe"][cls[:len(cls)-len(" [PE]")]] = -1
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

	type DrugIngredient struct {
		NDC          string
		IngredientID int
	}

	type Patient struct {
		ID, Age, AgeUnit int
		Sex, DeathDate   string
	}

	type PatientDrug struct {
		NDC              string
		PatientID        int
		Action           string
		Characterization string
		RouteID          int
	}

	type Term struct {
		ID          int
		Term        string
		TermVersion int
	}

	type Reaction struct {
		PatientID int
		TermID    int
		Outcome   int
	}

	type Event struct {
		ID            string
		ReportVersion int

		ReceiveDatem, ReceiptDatem, TransmissionDate time.Time

		Duplicate, Serious, SerDeath, SerDisabiling, SerHospitalization, SerLifeThreatening, SerOther bool

		CompanyID, Country, Patient int
	}

	fmt.Print("\ncountries\n\n")
	for i, item := range countries {
		fmt.Printf("%4d: %T %v\n", i, item, item)
	}
	fmt.Print("\nroutes\n\n")
	for i, item := range routes {
		fmt.Printf("%4d: %T %v\n", i, item, item)
	}
	fmt.Print("\nmanufacturers\n\n")
	for i, item := range manufacturers {
		fmt.Printf("%4d: %T %v\n", i, item, item)
	}
	fmt.Print("\ningredients\n\n")
	for i, item := range ingredients {
		fmt.Printf("%4d: %T %v\n", i, item, item)
	}
	fmt.Print("\nclasses\n\n")
	for i, item := range classes {
		fmt.Printf("%4d: %T %v\n", i, item, item)
	}
	fmt.Print("\ndrugs\n\n")
	for i, item := range drugs {
		fmt.Printf("%4d: %T %v\n", i, item, item)
	}
	fmt.Print("\ndrugNames\n\n")
	for i, item := range drugNames {
		fmt.Printf("%4d: %T %v\n", i, item, item)
	}
	fmt.Print("\ndrugManufacturers\n\n")
	for i, item := range drugManufacturers {
		fmt.Printf("%4d: %T %v\n", i, item, item)
	}
	fmt.Print("\ndrugClasses\n\n")
	for i, item := range drugClasses {
		fmt.Printf("%4d: %T %v\n", i, item, item)
	}
	fmt.Print("\nlabels\n\n")
	for i, item := range labels {
		fmt.Printf("%4d: %T %v\n", i, item, item)
	}
	//
	// db, err := getMySQLSession()
	// if err != nil {
	// 	log.Fatalf("error could not get db Seesion: %s", err)
	// }
}
