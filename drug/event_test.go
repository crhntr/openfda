package drug_test

import (
	"encoding/json"
	"testing"

	"github.com/crhntr/openfda/drug"
	"github.com/globalsign/mgo/bson"
)

func TestRawEvent_Encoding(t *testing.T) {
	var rawEvents []drug.RawEvent
	if err := json.Unmarshal([]byte(testBSONTestData), &rawEvents); err != nil {
		t.Error(err)
	}

	buf, err := bson.Marshal(rawEvents)
	t.Log(len(buf), err)
}

func TestEvent_Encoding(t *testing.T) {
	var (
		rawEvents []drug.RawEvent
		events    []drug.Event
	)
	if err := json.Unmarshal([]byte(testBSONTestData), &rawEvents); err != nil {
		t.Error(err)
	}

	for _, re := range rawEvents {
		events = append(events, re.Event())
	}

	buf, err := bson.Marshal(events)
	t.Log(len(buf), err)
}

const testBSONTestData = `[
  {
      "receivedate": "20151116",
      "seriousnessother": "1",
      "patient": {
        "reaction": [
          {
            "reactionmeddrapt": "Spinal fracture",
            "reactionmeddraversionpt": "19.0",
            "reactionoutcome": "6"
          },
          {
            "reactionmeddrapt": "Fall",
            "reactionmeddraversionpt": "19.0",
            "reactionoutcome": "6"
          }
        ],
        "patientsex": "2",
        "drug": [
          {
            "drugcharacterization": "1",
            "drugauthorizationnumb": "103960",
            "medicinalproduct": "HUMATE",
            "drugindication": "PRODUCT USED FOR UNKNOWN INDICATION",
            "activesubstance": {
              "activesubstancename": "HUMAN COAGULATION FACTOR VIII/VON WILLEBRAND FACTOR COMPLEX"
            }
          }
        ]
      },
      "sender": {
        "senderorganization": "FDA-Public Use",
        "sendertype": "2"
      },
      "primarysourcecountry": "US",
      "transmissiondate": "20160525",
      "companynumb": "US-BEH-2015055685",
      "occurcountry": "US",
      "receiptdate": "20160120",
      "transmissiondateformat": "102",
      "receiptdateformat": "102",
      "primarysource": {
        "reportercountry": "US",
        "qualification": "5"
      },
      "duplicate": "1",
      "reporttype": "2",
      "reportduplicate": {
        "duplicatenumb": "US-BEH-2015055685",
        "duplicatesource": "CSL BEHRING"
      },
      "receiver": {
        "receiverorganization": "FDA",
        "receivertype": "6"
      },
      "serious": "1",
      "safetyreportversion": "2",
      "fulfillexpeditecriteria": "1",
      "safetyreportid": "11741988",
      "receivedateformat": "102"
    },
    {
      "receivedate": "20151123",
      "seriousnessother": "1",
      "patient": {
        "reaction": [
          {
            "reactionmeddrapt": "Muscle tightness",
            "reactionmeddraversionpt": "19.0"
          },
          {
            "reactionmeddrapt": "Screaming",
            "reactionmeddraversionpt": "19.0"
          },
          {
            "reactionmeddrapt": "Therapeutic response decreased",
            "reactionmeddraversionpt": "19.0"
          },
          {
            "reactionmeddrapt": "Muscle spasms",
            "reactionmeddraversionpt": "19.0"
          },
          {
            "reactionmeddrapt": "Hypoaesthesia",
            "reactionmeddraversionpt": "19.0"
          },
          {
            "reactionmeddrapt": "Skin abrasion",
            "reactionmeddraversionpt": "19.0"
          },
          {
            "reactionmeddrapt": "Pain",
            "reactionmeddraversionpt": "19.0"
          },
          {
            "reactionmeddrapt": "Pain in extremity",
            "reactionmeddraversionpt": "19.0"
          },
          {
            "reactionmeddrapt": "Insomnia",
            "reactionmeddraversionpt": "19.0"
          },
          {
            "reactionmeddrapt": "Limb discomfort",
            "reactionmeddraversionpt": "19.0"
          },
          {
            "reactionmeddrapt": "Performance status decreased",
            "reactionmeddraversionpt": "19.0"
          },
          {
            "reactionmeddrapt": "Urinary tract infection",
            "reactionmeddraversionpt": "19.0"
          },
          {
            "reactionmeddrapt": "Influenza",
            "reactionmeddraversionpt": "19.0"
          }
        ],
        "patientonsetage": "68",
        "patientsex": "1",
        "patientonsetageunit": "801",
        "drug": [
          {
            "medicinalproduct": "BACLOFEN INTRATHECAL 2,000MCG/ML",
            "drugindication": "MUSCLE SPASTICITY",
            "drugdosagetext": "157.3 MCG/DAY",
            "drugcharacterization": "1",
            "activesubstance": {
              "activesubstancename": "BACLOFEN"
            },
            "drugrecurreadministration": "3",
            "drugauthorizationnumb": "020075"
          }
        ]
      },
      "sender": {
        "senderorganization": "FDA-Public Use",
        "sendertype": "2"
      },
      "primarysourcecountry": "US",
      "transmissiondate": "20160525",
      "fulfillexpeditecriteria": "1",
      "receiptdate": "20160304",
      "transmissiondateformat": "102",
      "receiptdateformat": "102",
      "primarysource": {
        "reportercountry": "US",
        "qualification": "5"
      },
      "duplicate": "1",
      "reporttype": "1",
      "reportduplicate": {
        "duplicatenumb": "US-MDT-ADR-2015-02208",
        "duplicatesource": "MEDTRONIC"
      },
      "receiver": {
        "receiverorganization": "FDA",
        "receivertype": "6"
      },
      "serious": "1",
      "safetyreportversion": "4",
      "companynumb": "US-MDT-ADR-2015-02208",
      "safetyreportid": "11768583",
      "receivedateformat": "102"
    },
    {
      "receivedate": "20151125",
      "receivedateformat": "102",
      "patient": {
        "reaction": [
          {
            "reactionmeddrapt": "Hyperglycaemia",
            "reactionmeddraversionpt": "19.0",
            "reactionoutcome": "6"
          }
        ],
        "patientonsetage": null,
        "drug": [
          {
            "drugstartdateformat": "102",
            "medicinalproduct": "DEURSIL",
            "drugindication": "PRODUCT USED FOR UNKNOWN INDICATION",
            "drugadministrationroute": "065",
            "drugdosagetext": "300 MG, UNK",
            "drugstructuredosageunit": "003",
            "drugcharacterization": "2",
            "drugstartdate": "20150310",
            "activesubstance": {
              "activesubstancename": "URSODIOL"
            },
            "actiondrug": "4",
            "drugstructuredosagenumb": "300"
          },
          {
            "drugstartdateformat": "102",
            "medicinalproduct": "SOM230",
            "drugindication": "PITUITARY-DEPENDENT CUSHING^S SYNDROME",
            "drugadministrationroute": "030",
            "drugenddateformat": "102",
            "drugdosagetext": "DOUBLE BLIND",
            "drugcharacterization": "1",
            "drugstartdate": "20141209",
            "actiondrug": "1",
            "activesubstance": {
              "activesubstancename": "PASIREOTIDE"
            },
            "drugauthorizationnumb": "200677",
            "drugdosageform": "VIAL",
            "drugenddate": "20151110"
          }
        ],
        "summary": {
          "narrativeincludeclinical": "CASE EVENT DATE: 20151019"
        },
        "patientweight": "71.5",
        "patientsex": "2",
        "patientonsetageunit": "801"
      },
      "sender": {
        "senderorganization": "FDA-Public Use",
        "sendertype": "2"
      },
      "transmissiondate": "20160525",
      "companynumb": "PHHO2015IT019433",
      "occurcountry": "IT",
      "receiptdate": "20160216",
      "transmissiondateformat": "102",
      "receiptdateformat": "102",
      "primarysource": {
        "reportercountry": "IT",
        "qualification": "1"
      },
      "duplicate": "1",
      "reporttype": "2",
      "seriousnesshospitalization": "1",
      "reportduplicate": {
        "duplicatenumb": "PHHO2015IT019433",
        "duplicatesource": "NOVARTIS"
      },
      "receiver": {
        "receiverorganization": "FDA",
        "receivertype": "6"
      },
      "serious": "1",
      "safetyreportversion": "4",
      "fulfillexpeditecriteria": "1",
      "safetyreportid": "11778722",
      "primarysourcecountry": "IT"
    },
    {
      "receivedate": "20151208",
      "seriousnessother": "1",
      "patient": {
        "reaction": [
          {
            "reactionmeddrapt": "Orbital oedema",
            "reactionmeddraversionpt": "19.0",
            "reactionoutcome": "1"
          },
          {
            "reactionmeddrapt": "Overdose",
            "reactionmeddraversionpt": "19.0",
            "reactionoutcome": "6"
          },
          {
            "reactionmeddrapt": "Face oedema",
            "reactionmeddraversionpt": "19.0",
            "reactionoutcome": "1"
          }
        ],
        "patientsex": "1",
        "drug": [
          {
            "medicinalproduct": "FLIXONASE",
            "actiondrug": "1",
            "drugdosagetext": "UNK",
            "drugcharacterization": "1",
            "drugcumulativedosageunit": "032",
            "activesubstance": {
              "activesubstancename": "FLUTICASONE PROPIONATE"
            },
            "drugrecurreadministration": "2",
            "drugcumulativedosagenumb": "3",
            "drugauthorizationnumb": "020121"
          },
          {
            "drugstartdateformat": "102",
            "medicinalproduct": "FLIXONASE",
            "drugindication": "RHINITIS ALLERGIC",
            "drugintervaldosageunitnumb": "1",
            "drugenddateformat": "102",
            "drugseparatedosagenumb": "3",
            "drugstructuredosageunit": "032",
            "drugdosagetext": "1 DF, TID  (SPRAY) (MORNING, NOON AND NIGHT)",
            "drugcumulativedosagenumb": "3",
            "drugcharacterization": "1",
            "drugcumulativedosageunit": "032",
            "drugstartdate": "20151130",
            "actiondrug": "1",
            "activesubstance": {
              "activesubstancename": "FLUTICASONE PROPIONATE"
            },
            "drugintervaldosagedefinition": "804",
            "drugauthorizationnumb": "020121",
            "drugrecurreadministration": "2",
            "drugbatchnumb": "3C5B",
            "drugenddate": "20151201",
            "drugstructuredosagenumb": "1"
          }
        ],
        "summary": {
          "narrativeincludeclinical": "CASE EVENT DATE: 20151130"
        }
      },
      "sender": {
        "senderorganization": "FDA-Public Use",
        "sendertype": "2"
      },
      "primarysourcecountry": "CN",
      "transmissiondate": "20160525",
      "companynumb": "CN-GLAXOSMITHKLINE-CN2015GSK171795",
      "occurcountry": "CN",
      "receiptdate": "20160118",
      "transmissiondateformat": "102",
      "receiptdateformat": "102",
      "primarysource": {
        "reportercountry": "CN",
        "qualification": "5"
      },
      "duplicate": "1",
      "reporttype": "1",
      "reportduplicate": {
        "duplicatenumb": "CN-GLAXOSMITHKLINE-CN2015GSK171795",
        "duplicatesource": "GLAXOSMITHKLINE"
      },
      "receiver": {
        "receiverorganization": "FDA",
        "receivertype": "6"
      },
      "serious": "1",
      "safetyreportversion": "3",
      "fulfillexpeditecriteria": "1",
      "safetyreportid": "11812020",
      "receivedateformat": "102"
    },
    {
      "receivedate": "20151210",
      "receivedateformat": "102",
      "patient": {
        "reaction": [
          {
            "reactionmeddrapt": "Myocardial infarction",
            "reactionmeddraversionpt": "19.0",
            "reactionoutcome": "6"
          }
        ],
        "patientonsetage": "68",
        "drug": [
          {
            "medicinalproduct": "DIANEAL LOW CALCIUM PERITONEAL DIALYSIS SOLUTION WITH DEXTROSE",
            "drugindication": "PERITONEAL DIALYSIS",
            "drugadministrationroute": "033",
            "actiondrug": "4",
            "drugcharacterization": "1",
            "activesubstance": {
              "activesubstancename": "CALCIUM CHLORIDE\\DEXTROSE\\MAGNESIUM CHLORIDE\\SODIUM CHLORIDE\\SODIUM LACTATE"
            },
            "drugauthorizationnumb": "020183",
            "drugdosageform": "SOLUTION FOR PERITONEAL DIALYSIS"
          },
          {
            "medicinalproduct": "DIANEAL LOW CALCIUM PERITONEAL DIALYSIS SOLUTION WITH DEXTROSE",
            "drugindication": "PERITONEAL DIALYSIS",
            "drugadministrationroute": "033",
            "actiondrug": "4",
            "drugcharacterization": "1",
            "activesubstance": {
              "activesubstancename": "CALCIUM CHLORIDE\\DEXTROSE\\MAGNESIUM CHLORIDE\\SODIUM CHLORIDE\\SODIUM LACTATE"
            },
            "drugauthorizationnumb": "017512",
            "drugdosageform": "SOLUTION FOR PERITONEAL DIALYSIS"
          }
        ],
        "summary": {
          "narrativeincludeclinical": "CASE EVENT DATE: 20151129"
        },
        "patientsex": "1",
        "patientonsetageunit": "801"
      },
      "sender": {
        "senderorganization": "FDA-Public Use",
        "sendertype": "2"
      },
      "transmissiondate": "20160525",
      "companynumb": "US-BAXTER-2015BAX067020",
      "occurcountry": "US",
      "receiptdate": "20160304",
      "transmissiondateformat": "102",
      "receiptdateformat": "102",
      "primarysource": {
        "reportercountry": "US",
        "qualification": "5"
      },
      "duplicate": "1",
      "reporttype": "1",
      "seriousnesshospitalization": "1",
      "reportduplicate": {
        "duplicatenumb": "US-BAXTER-2015BAX067020",
        "duplicatesource": "BAXTER"
      },
      "receiver": {
        "receiverorganization": "FDA",
        "receivertype": "6"
      },
      "serious": "1",
      "safetyreportversion": "3",
      "fulfillexpeditecriteria": "1",
      "safetyreportid": "11823178",
      "primarysourcecountry": "US"
    },
    {
      "reportduplicate": {
        "duplicatenumb": "US-BAXTER-2015BAX067448",
        "duplicatesource": "BAXTER"
      },
      "safetyreportversion": "3",
      "receiptdate": "20160125",
      "duplicate": "1",
      "seriousnessdeath": "1",
      "seriousnesshospitalization": "1",
      "receivedate": "20151215",
      "patient": {
        "reaction": [
          {
            "reactionmeddrapt": "Myocardial infarction",
            "reactionmeddraversionpt": "19.0",
            "reactionoutcome": "2"
          },
          {
            "reactionmeddrapt": "Death",
        "reactionmeddraversionpt": "19.0",
        "reactionoutcome": "5"
      },
      {
        "reactionmeddrapt": "Device related infection",
        "reactionmeddraversionpt": "19.0",
        "reactionoutcome": "6"
      }
    ],
    "patientonsetage": "62",
    "drug": [
      {
        "medicinalproduct": "DIANEAL LOW CALCIUM PERITONEAL DIALYSIS SOLUTION WITH DEXTROSE",
        "drugindication": "PERITONEAL DIALYSIS",
        "drugadministrationroute": "033",
        "drugcharacterization": "1",
        "activesubstance": {
          "activesubstancename": "CALCIUM CHLORIDE\\DEXTROSE\\MAGNESIUM CHLORIDE\\SODIUM CHLORIDE\\SODIUM LACTATE"
        },
        "drugauthorizationnumb": "017512",
        "drugdosageform": "SOLUTION FOR PERITONEAL DIALYSIS"
      }
    ],
    "summary": {
      "narrativeincludeclinical": "CASE EVENT DATE: 20151130"
    },
    "patientsex": "1",
    "patientonsetageunit": "801"
  },
  "occurcountry": "US",
  "reporttype": "1",
  "companynumb": "US-BAXTER-2015BAX067448",
  "safetyreportid": "11836400",
  "sender": {
    "senderorganization": "FDA-Public Use",
    "sendertype": "2"
  },
  "transmissiondate": "20160525",
  "fulfillexpeditecriteria": "1",
  "transmissiondateformat": "102",
  "receiptdateformat": "102",
  "receiver": {
    "receiverorganization": "FDA",
    "receivertype": "6"
  },
  "serious": "1",
  "receivedateformat": "102",
  "primarysource": {
    "reportercountry": "US",
    "qualification": "5"
  },
  "primarysourcecountry": "US"
},
{
  "receivedate": "20151228",
  "receivedateformat": "102",
  "patient": {
    "reaction": [
      {
        "reactionmeddrapt": "Hyperglycaemia",
        "reactionmeddraversionpt": "19.0",
        "reactionoutcome": "6"
      },
      {
        "reactionmeddrapt": "Insulin resistance syndrome",
        "reactionmeddraversionpt": "19.0",
        "reactionoutcome": "1"
      }
    ],
    "drug": [
      {
        "medicinalproduct": "PREDNISOLONE MYLAN",
        "drugindication": "HYPOGLYCAEMIA",
        "drugadministrationroute": "065",
        "drugdosagetext": "30 MG FOR 2 WEEKS",
        "drugcharacterization": "1",
        "actiondrug": "2",
        "activesubstance": {
          "activesubstancename": "PREDNISOLONE"
        },
        "drugrecurreadministration": "3",
        "drugauthorizationnumb": "202179",
        "drugdosageform": "TABLET"
      },
      {
        "medicinalproduct": "INSULIN",
        "drugindication": "PRODUCT USED FOR UNKNOWN INDICATION",
        "drugadministrationroute": "042",
        "drugdosagetext": "0.1 U/KG",
        "drugcharacterization": "1",
        "activesubstance": {
          "activesubstancename": "INSULIN NOS"
        },
        "drugrecurreadministration": "3",
        "actiondrug": "1"
      }
    ]
  },
  "sender": {
    "senderorganization": "FDA-Public Use",
    "sendertype": "2"
  },
  "transmissiondate": "20160526",
  "companynumb": "JP-MYLANLABS-2015M1046478",
  "occurcountry": "JP",
  "receiptdate": "20160129",
  "transmissiondateformat": "102",
  "receiptdateformat": "102",
  "primarysource": {
    "literaturereference": "YAMADA H, ASANO T, KUSAKA I, KAKEI M, ISHIKAWA S-E. TYPE B INSULIN RESISTANCE SYNDROME WITH FASTING HYPOGLYCEMIA AND POSTPRANDIAL HYPERGLYCEMIA. DIABET-INT 2015?6(2):144-148.",
    "reportercountry": "JP",
    "qualification": "3"
  },
  "duplicate": "1",
  "reporttype": "1",
  "seriousnesshospitalization": "1",
  "reportduplicate": {
    "duplicatenumb": "JP-MYLANLABS-2015M1046478",
    "duplicatesource": "MYLAN"
  },
  "receiver": {
    "receiverorganization": "FDA",
    "receivertype": "6"
  },
  "serious": "1",
  "safetyreportversion": "3",
  "fulfillexpeditecriteria": "1",
  "safetyreportid": "11871285",
  "primarysourcecountry": "JP"
},
{
  "receivedate": "20151230",
  "receivedateformat": "102",
  "patient": {
    "reaction": [
      {
        "reactionmeddrapt": "Arrhythmia",
        "reactionmeddraversionpt": "19.0",
        "reactionoutcome": "5"
      }
    ],
    "patientonsetage": "92",
    "drug": [
      {
        "drugstartdateformat": "102",
        "medicinalproduct": "CINACALCET HCL",
        "drugindication": "HYPERPARATHYROIDISM PRIMARY",
        "drugcharacterization": "1",
        "drugadministrationroute": "048",
        "drugseparatedosagenumb": "2",
        "drugstructuredosageunit": "003",
        "drugdosagetext": "25 MG, BID",
        "drugintervaldosageunitnumb": "1",
        "drugstartdate": "20141204",
        "activesubstance": {
          "activesubstancename": "CINACALCET HYDROCHLORIDE"
        },
        "drugintervaldosagedefinition": "804",
        "drugauthorizationnumb": "021688",
        "drugdosageform": "TABLET",
        "drugstructuredosagenumb": "25"
      },
      {
        "drugstartdateformat": "102",
        "medicinalproduct": "SENNOSIDE                          /00571901/",
        "drugindication": "CONSTIPATION",
        "drugadministrationroute": "048",
        "drugdosagetext": "12 MG, UNK",
        "drugstructuredosageunit": "003",
        "drugcharacterization": "2",
        "drugstartdate": "20141204",
        "activesubstance": {
          "activesubstancename": "SENNOSIDES A AND B"
        },
        "drugstructuredosagenumb": "12"
      },
      {
        "medicinalproduct": "SODIUM RABEPRAZOLE",
        "drugindication": "GASTROOESOPHAGEAL REFLUX DISEASE",
        "drugadministrationroute": "048",
        "drugdosagetext": "10 MG, UNK",
        "drugstructuredosageunit": "003",
        "drugcharacterization": "2",
        "activesubstance": {
          "activesubstancename": "RABEPRAZOLE SODIUM"
        },
        "drugdosageform": "TABLET",
        "drugstructuredosagenumb": "10"
      },
      {
        "drugstartdateformat": "102",
        "medicinalproduct": "LECICARBON                         /00739901/",
        "drugindication": "CONSTIPATION",
        "drugadministrationroute": "054",
        "drugdosagetext": "1 DF, UNK",
        "drugstructuredosageunit": "032",
        "drugcharacterization": "2",
        "drugstartdate": "20141212",
        "activesubstance": {
          "activesubstancename": "LECITHIN\\SODIUM BICARBONATE\\SODIUM PHOSPHATE, MONOBASIC"
        },
        "drugstructuredosagenumb": "1"
      },
      {
        "medicinalproduct": "CIBENOL",
        "drugindication": "TACHYARRHYTHMIA",
        "drugadministrationroute": "048",
        "drugdosagetext": "100 MG, UNK",
        "drugstructuredosageunit": "003",
        "drugcharacterization": "1",
        "activesubstance": {
          "activesubstancename": "CIFENLINE"
        },
        "drugdosageform": "TABLET",
        "drugstructuredosagenumb": "100"
      }
    ],
    "summary": {
      "narrativeincludeclinical": "CASE EVENT DATE: 20141213"
    },
    "patientsex": "2",
    "patientagegroup": "6",
    "patientonsetageunit": "801"
  },
  "sender": {
    "senderorganization": "FDA-Public Use",
    "sendertype": "2"
  },
  "transmissiondate": "20160526",
  "companynumb": "JP-AMGEN-JPNCT2015139588",
  "occurcountry": "JP",
  "receiptdate": "20160229",
  "transmissiondateformat": "102",
  "receiptdateformat": "102",
  "primarysource": {
    "reportercountry": "JP",
    "qualification": "1"
  },
  "duplicate": "1",
  "seriousnessdeath": "1",
  "reporttype": "2",
  "reportduplicate": {
    "duplicatenumb": "JP-AMGEN-JPNCT2015139588",
    "duplicatesource": "AMGEN"
  },
  "receiver": {
    "receiverorganization": "FDA",
    "receivertype": "6"
  },
  "serious": "1",
  "safetyreportversion": "2",
  "fulfillexpeditecriteria": "1",
  "safetyreportid": "11877810",
  "primarysourcecountry": "JP"
},
{
  "receivedate": "20160108",
  "receivedateformat": "102",
  "patient": {
    "reaction": [
      {
        "reactionmeddrapt": "Procedural pain",
        "reactionmeddraversionpt": "19.0",
        "reactionoutcome": "6"
      }
    ],
    "drug": [
      {
        "medicinalproduct": "DIANEAL LOW CALCIUM PERITONEAL DIALYSIS SOLUTION WITH DEXTROSE",
        "drugindication": "PERITONEAL DIALYSIS",
        "drugadministrationroute": "033",
        "actiondrug": "4",
        "drugcharacterization": "1",
        "activesubstance": {
          "activesubstancename": "CALCIUM CHLORIDE\\DEXTROSE\\MAGNESIUM CHLORIDE\\SODIUM CHLORIDE\\SODIUM LACTATE"
        },
        "drugauthorizationnumb": "017512",
        "drugdosageform": "SOLUTION FOR PERITONEAL DIALYSIS"
      }
    ]
  },
  "sender": {
    "senderorganization": "FDA-Public Use",
    "sendertype": "2"
  },
  "transmissiondate": "20160526",
  "fulfillexpeditecriteria": "2",
  "occurcountry": "US",
  "receiptdate": "20160108",
  "transmissiondateformat": "102",
  "receiptdateformat": "102",
  "primarysource": {
    "reportercountry": "US",
    "qualification": "5"
  },
  "duplicate": "1",
  "reporttype": "1",
  "reportduplicate": {
    "duplicatenumb": "US-BAXTER-2016BAX000667",
    "duplicatesource": "BAXTER"
  },
  "receiver": {
    "receiverorganization": "FDA",
    "receivertype": "6"
  },
  "serious": "2",
  "safetyreportversion": "1",
  "companynumb": "US-BAXTER-2016BAX000667",
  "safetyreportid": "11899287",
  "primarysourcecountry": "US"
},
{
  "reportduplicate": {
    "duplicatenumb": "JP-CELGENE-JPN-2015115782",
    "duplicatesource": "CELGENE"
  },
  "safetyreportversion": "1",
  "receiptdate": "20160107",
  "seriousnesslifethreatening": "1",
  "duplicate": "1",
  "seriousnesshospitalization": "1",
  "receivedate": "20160107",
  "patient": {
    "reaction": [
      {
        "reactionmeddrapt": "Interstitial lung disease",
        "reactionmeddraversionpt": "19.0",
        "reactionoutcome": "1"
      }
    ],
    "patientonsetage": "80",
    "drug": [
      {
        "drugstartdateformat": "102",
        "medicinalproduct": "ABRAXANE",
        "drugindication": "METASTATIC GASTRIC CANCER",
        "drugcharacterization": "1",
        "drugadministrationroute": "041",
        "drugenddateformat": "102",
        "drugseparatedosagenumb": "1",
        "drugstructuredosageunit": "003",
        "drugdosagetext": "380 MILLIGRAM",
        "drugintervaldosageunitnumb": "1",
        "drugstartdate": "20150701",
        "actiondrug": "1",
        "activesubstance": {
          "activesubstancename": "PACLITAXEL"
        },
        "drugintervaldosagedefinition": "804",
        "drugauthorizationnumb": "021660",
        "drugrecurreadministration": "3",
        "drugdosageform": "INJECTION FOR INFUSION",
        "drugenddate": "20150701",
        "drugstructuredosagenumb": "380"
      }
    ],
    "summary": {
      "narrativeincludeclinical": "CASE EVENT DATE: 20150708"
    },
    "patientweight": "49",
    "patientsex": "1",
    "patientonsetageunit": "801"
  },
  "occurcountry": "JP",
  "reporttype": "1",
  "companynumb": "JP-CELGENE-JPN-2015115782",
  "safetyreportid": "11895285",
  "sender": {
    "senderorganization": "FDA-Public Use",
    "sendertype": "2"
  },
  "transmissiondate": "20160526",
  "fulfillexpeditecriteria": "1",
  "transmissiondateformat": "102",
  "receiptdateformat": "102",
  "receiver": {
    "receiverorganization": "FDA",
    "receivertype": "6"
  },
  "serious": "1",
  "receivedateformat": "102",
  "primarysource": {
    "reportercountry": "JP",
    "qualification": "1"
  },
  "primarysourcecountry": "JP"
},
{
  "seriousnessother": "1",
  "reportduplicate": {
    "duplicatenumb": "JP-GLAXOSMITHKLINE-JP2015JPN169095",
    "duplicatesource": "GLAXOSMITHKLINE"
  },
  "safetyreportversion": "2",
  "receiptdate": "20160225",
  "duplicate": "1",
  "seriousnesshospitalization": "1",
  "receivedate": "20151202",
  "patient": {
    "reaction": [
      {
        "reactionmeddrapt": "Prostate cancer",
        "reactionmeddraversionpt": "19.0",
        "reactionoutcome": "1"
      }
    ],
    "patientonsetage": null,
    "drug": [
      {
        "drugstartdateformat": "102",
        "medicinalproduct": "AVOLVE",
        "drugindication": "BENIGN PROSTATIC HYPERPLASIA",
        "drugcharacterization": "1",
        "drugadministrationroute": "048",
        "drugenddateformat": "102",
        "drugseparatedosagenumb": "1",
        "drugstructuredosageunit": "003",
        "drugdosagetext": "0.5 MG, QD",
        "drugcumulativedosagenumb": "537",
        "drugintervaldosageunitnumb": "1",
        "drugcumulativedosageunit": "003",
        "drugstartdate": "20100304",
        "activesubstance": {
          "activesubstancename": "DUTASTERIDE"
        },
        "drugintervaldosagedefinition": "804",
        "drugauthorizationnumb": "021319",
        "drugdosageform": "CAPSULE",
        "drugenddate": "20130210",
        "drugstructuredosagenumb": ".5"
      },
      {
        "drugcharacterization": "2",
        "drugdosagetext": "75 MG, UNK",
        "drugstructuredosageunit": "003",
        "medicinalproduct": "FLIVAS OD",
        "drugstructuredosagenumb": "75"
      }
    ],
    "summary": {
      "narrativeincludeclinical": "CASE EVENT DATE: 20130218"
    },
    "patientweight": "66",
    "patientsex": "1",
    "patientonsetageunit": "801"
  },
  "occurcountry": "JP",
  "reporttype": "1",
  "companynumb": "JP-GLAXOSMITHKLINE-JP2015JPN169095",
  "safetyreportid": "11793793",
  "sender": {
    "senderorganization": "FDA-Public Use",
    "sendertype": "2"
  },
  "transmissiondate": "20160525",
  "fulfillexpeditecriteria": "1",
  "transmissiondateformat": "102",
  "receiptdateformat": "102",
  "receiver": {
    "receiverorganization": "FDA",
    "receivertype": "6"
  },
  "serious": "1",
  "receivedateformat": "102",
  "primarysource": {
    "reportercountry": "JP",
    "qualification": "5"
  },
  "primarysourcecountry": "JP"
}]`
