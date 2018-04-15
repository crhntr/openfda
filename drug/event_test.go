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
		e := re.Event()
		events = append(events, e)
	}

	buf, err := bson.Marshal(events)
	t.Log(len(buf), err)
}

const testBSONTestData = `[{
      "receivedate": "20170929",
      "seriousnessother": "1",
      "patient": {
        "reaction": [
          {
            "reactionmeddrapt": "Full blood count decreased",
            "reactionmeddraversionpt": "20.1",
            "reactionoutcome": "1"
          }
        ],
        "patientonsetage": "88",
        "drug": [
          {
            "drugstartdateformat": "102",
            "medicinalproduct": "THALOMID",
            "drugcharacterization": "1",
            "drugadministrationroute": "048",
            "drugseparatedosagenumb": "1",
            "drugstructuredosageunit": "003",
            "openfda": {
              "manufacturer_name": [
                "Celgene Corporation"
              ],
              "unii": [
                "4Z8R6ORS6L"
              ],
              "product_type": [
                "HUMAN PRESCRIPTION DRUG"
              ],
              "rxcui": [
                "200390",
                "213360",
                "700418",
                "700416",
                "404449",
                "404450",
                "360177",
                "360176"
              ],
              "spl_set_id": [
                "2eda833b-1357-4ed4-a093-194524fcb061"
              ],
              "route": [
                "ORAL"
              ],
              "generic_name": [
                "THALIDOMIDE"
              ],
              "brand_name": [
                "THALOMID"
              ],
              "pharm_class_pe": [
                "Decreased Immunologically Active Molecule Activity [PE]"
              ],
              "product_ndc": [
                "59572-215",
                "59572-220",
                "59572-205",
                "59572-210"
              ],
              "substance_name": [
                "THALIDOMIDE"
              ],
              "spl_id": [
                "0ef5e838-f546-4568-96d1-5d1459b849ec"
              ],
              "application_number": [
                "NDA020785"
              ],
              "nui": [
                "N0000008663"
              ],
              "package_ndc": [
                "59572-215-93",
                "59572-205-17",
                "59572-205-14",
                "59572-215-13",
                "59572-220-16",
                "59572-205-97",
                "59572-210-95",
                "59572-210-15",
                "59572-205-94",
                "59572-220-96"
              ]
            },
            "drugstructuredosagenumb": "50",
            "drugbatchnumb": "UNKNOWN",
            "drugintervaldosageunitnumb": "1",
            "drugstartdate": "20170609",
            "actiondrug": "1",
            "drugdosagetext": "50 MILLIGRAM",
            "drugintervaldosagedefinition": "804",
            "drugauthorizationnumb": "020785",
            "drugrecurreadministration": "3",
            "activesubstance": {
              "activesubstancename": "THALIDOMIDE"
            },
            "drugdosageform": "CAPSULES",
            "drugadditional": "1"
          },
          {
            "drugstartdateformat": "610",
            "medicinalproduct": "THALOMID",
            "drugindication": "PLASMA CELL MYELOMA",
            "drugcharacterization": "1",
            "drugadministrationroute": "048",
            "drugseparatedosagenumb": "1",
            "drugstructuredosageunit": "003",
            "openfda": {
              "manufacturer_name": [
                "Celgene Corporation"
              ],
              "unii": [
                "4Z8R6ORS6L"
              ],
              "product_type": [
                "HUMAN PRESCRIPTION DRUG"
              ],
              "rxcui": [
                "200390",
                "213360",
                "700418",
                "700416",
                "404449",
                "404450",
                "360177",
                "360176"
              ],
              "spl_set_id": [
                "2eda833b-1357-4ed4-a093-194524fcb061"
              ],
              "route": [
                "ORAL"
              ],
              "generic_name": [
                "THALIDOMIDE"
              ],
              "brand_name": [
                "THALOMID"
              ],
              "pharm_class_pe": [
                "Decreased Immunologically Active Molecule Activity [PE]"
              ],
              "product_ndc": [
                "59572-215",
                "59572-220",
                "59572-205",
                "59572-210"
              ],
              "substance_name": [
                "THALIDOMIDE"
              ],
              "spl_id": [
                "0ef5e838-f546-4568-96d1-5d1459b849ec"
              ],
              "application_number": [
                "NDA020785"
              ],
              "nui": [
                "N0000008663"
              ],
              "package_ndc": [
                "59572-215-93",
                "59572-205-17",
                "59572-205-14",
                "59572-215-13",
                "59572-220-16",
                "59572-205-97",
                "59572-210-95",
                "59572-210-15",
                "59572-205-94",
                "59572-220-96"
              ]
            },
            "drugstructuredosagenumb": "50",
            "drugbatchnumb": "UNKNOWN",
            "drugintervaldosageunitnumb": "1",
            "drugstartdate": "201602",
            "actiondrug": "1",
            "drugdosagetext": "50 MILLIGRAM",
            "drugintervaldosagedefinition": "804",
            "drugauthorizationnumb": "020785",
            "drugrecurreadministration": "3",
            "activesubstance": {
              "activesubstancename": "THALIDOMIDE"
            },
            "drugdosageform": "CAPSULES",
            "drugadditional": "1"
          }
        ],
        "summary": {
          "narrativeincludeclinical": "CASE EVENT DATE: 20170901"
        },
        "patientsex": "1",
        "patientonsetageunit": "801"
      },
      "sender": {
        "senderorganization": "FDA-Public Use",
        "sendertype": "2"
      },
      "primarysourcecountry": "US",
      "transmissiondate": "20171128",
      "companynumb": "US-CELGENEUS-USA-20170907738",
      "occurcountry": "US",
      "receiptdate": "20170929",
      "transmissiondateformat": "102",
      "receiptdateformat": "102",
      "primarysource": {
        "reportercountry": "US",
        "qualification": "2"
      },
      "duplicate": "1",
      "reporttype": "1",
      "reportduplicate": {
        "duplicatenumb": "US-CELGENEUS-USA-20170907738",
        "duplicatesource": "CELGENE"
      },
      "receiver": {
        "receiverorganization": "FDA",
        "receivertype": "6"
      },
      "serious": "1",
      "safetyreportversion": "1",
      "fulfillexpeditecriteria": "2",
      "safetyreportid": "14026459",
      "receivedateformat": "102"
    },
    {
      "receivedate": "20170929",
      "receivedateformat": "102",
      "patient": {
        "reaction": [
          {
            "reactionmeddrapt": "Death",
            "reactionmeddraversionpt": "20.1",
            "reactionoutcome": "5"
          }
        ],
        "drug": [
          {
            "drugstartdateformat": "602",
            "medicinalproduct": "ACTILYSE",
            "drugindication": "ISCHAEMIC STROKE",
            "drugadministrationroute": "065",
            "actiondrug": "6",
            "drugcharacterization": "1",
            "drugstartdate": "2009",
            "activesubstance": {
              "activesubstancename": "ALTEPLASE"
            },
            "drugauthorizationnumb": "103172"
          }
        ]
      },
      "sender": {
        "senderorganization": "FDA-Public Use",
        "sendertype": "2"
      },
      "transmissiondate": "20171128",
      "companynumb": "HU-ROCHE-1976831",
      "occurcountry": "HU",
      "receiptdate": "20170929",
      "transmissiondateformat": "102",
      "receiptdateformat": "102",
      "primarysource": {
        "reportercountry": "HU",
        "qualification": "1"
      },
      "duplicate": "1",
      "seriousnessdeath": "1",
      "reporttype": "1",
      "reportduplicate": {
        "duplicatenumb": "HU-ROCHE-1976831",
        "duplicatesource": "ROCHE"
      },
      "receiver": {
        "receiverorganization": "FDA",
        "receivertype": "6"
      },
      "serious": "1",
      "safetyreportversion": "1",
      "fulfillexpeditecriteria": "1",
      "safetyreportid": "14026543",
      "primarysourcecountry": "HU"
    },
    {
      "seriousnessother": "1",
      "reportduplicate": {
        "duplicatenumb": "IE-ASTELLAS-2017US038994",
        "duplicatesource": "ASTELLAS"
      },
      "safetyreportversion": "1",
      "receiptdate": "20170929",
      "duplicate": "1",
      "seriousnessdeath": "1",
      "receivedate": "20170929",
      "patient": {
        "reaction": [
          {
            "reactionmeddrapt": "Sarcoma",
            "reactionmeddraversionpt": "20.1",
            "reactionoutcome": "5"
          },
          {
            "reactionmeddrapt": "Renal transplant",
            "reactionmeddraversionpt": "20.1",
            "reactionoutcome": "6"
          }
        ],
        "patientonsetage": "42",
        "drug": [
          {
            "medicinalproduct": "PROGRAF",
            "drugindication": "LUNG TRANSPLANT",
            "drugadministrationroute": "065",
            "drugdosagetext": "UNK UNK, UNKNOWN FREQ.",
            "openfda": {
              "manufacturer_name": [
                "Astellas Pharma US, Inc."
              ],
              "unii": [
                "WM0HAQ4WNM"
              ],
              "product_type": [
                "HUMAN PRESCRIPTION DRUG"
              ],
              "rxcui": [
                "198378",
                "198379",
                "198377",
                "261134",
                "313190",
                "108515",
                "108514",
                "108513"
              ],
              "spl_set_id": [
                "7f667de1-9dfa-4bd6-8ba0-15ee2d78873b"
              ],
              "route": [
                "ORAL",
                "INTRAVENOUS"
              ],
              "generic_name": [
                "TACROLIMUS"
              ],
              "brand_name": [
                "PROGRAF"
              ],
              "product_ndc": [
                "0469-0607",
                "0469-3016",
                "0469-0657",
                "0469-0617"
              ],
              "substance_name": [
                "TACROLIMUS"
              ],
              "spl_id": [
                "426e5fca-bc45-400f-9d6e-355559f7554e"
              ],
              "application_number": [
                "NDA050708",
                "NDA050709"
              ],
              "package_ndc": [
                "0469-0617-73",
                "0469-0657-11",
                "0469-3016-01",
                "0469-0607-73",
                "0469-0657-73",
                "0469-0617-11"
              ]
            },
            "drugcharacterization": "1",
            "actiondrug": "6",
            "activesubstance": {
              "activesubstancename": "TACROLIMUS\\TACROLIMUS ANHYDROUS"
            },
            "drugauthorizationnumb": "050708",
            "drugdosageform": "FORMULATION UNKNOWN"
          }
        ],
        "summary": {
          "narrativeincludeclinical": "CASE EVENT DATE: 2015"
        },
        "patientsex": "2",
        "patientonsetageunit": "801"
      },
      "occurcountry": "IE",
      "reporttype": "1",
      "companynumb": "IE-ASTELLAS-2017US038994",
      "safetyreportid": "14026598",
      "sender": {
        "senderorganization": "FDA-Public Use",
        "sendertype": "2"
      },
      "transmissiondate": "20171128",
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
        "reportercountry": "IE",
        "qualification": "5"
      },
      "primarysourcecountry": "IE"
    },
    {
      "receivedate": "20170929",
      "seriousnessother": "1",
      "patient": {
        "reaction": [
          {
            "reactionmeddrapt": "Blood pressure increased",
            "reactionmeddraversionpt": "20.1",
            "reactionoutcome": "2"
          },
          {
            "reactionmeddrapt": "Breast mass",
            "reactionmeddraversionpt": "20.1",
            "reactionoutcome": "3"
          },
          {
            "reactionmeddrapt": "Feeling hot",
            "reactionmeddraversionpt": "20.1",
            "reactionoutcome": "2"
          }
        ],
        "drug": [
          {
            "medicinalproduct": "ESTRADIOL/NORETHISTERONE ACETATE",
            "drugindication": "HORMONE REPLACEMENT THERAPY",
            "drugadministrationroute": "062",
            "drugdosagetext": "1 DF (1 PATCH A DAY 2 TIMES A WEEK), UNK",
            "drugstructuredosageunit": "032",
            "drugbatchnumb": "UNKNOWN",
            "drugcharacterization": "1",
            "actiondrug": "4",
            "activesubstance": {
              "activesubstancename": "ESTRADIOL\\NORETHINDRONE ACETATE"
            },
            "drugauthorizationnumb": "020870",
            "drugdosageform": "TRANSDERMAL PATCH",
            "drugstructuredosagenumb": "1"
          }
        ]
      },
      "sender": {
        "senderorganization": "FDA-Public Use",
        "sendertype": "2"
      },
      "primarysourcecountry": "BR",
      "transmissiondate": "20171128",
      "companynumb": "BR-NOVEN PHARMACEUTICALS, INC.-BR2017000747",
      "occurcountry": "BR",
      "receiptdate": "20170929",
      "transmissiondateformat": "102",
      "receiptdateformat": "102",
      "primarysource": {
        "reportercountry": "BR",
        "qualification": "5"
      },
      "duplicate": "1",
      "reporttype": "1",
      "reportduplicate": {
        "duplicatenumb": "BR-NOVEN PHARMACEUTICALS, INC.-BR2017000747",
        "duplicatesource": "NOVEN"
      },
      "receiver": {
        "receiverorganization": "FDA",
        "receivertype": "6"
      },
      "serious": "1",
      "safetyreportversion": "1",
      "fulfillexpeditecriteria": "2",
      "safetyreportid": "14026752",
      "receivedateformat": "102"
    },
    {
      "receivedate": "20170929",
      "seriousnessother": "1",
      "patient": {
        "reaction": [
          {
            "reactionmeddrapt": "Sinus operation",
            "reactionmeddraversionpt": "20.1",
            "reactionoutcome": "6"
          }
        ],
        "patientonsetage": "22",
        "drug": [
          {
            "drugstartdateformat": "102",
            "medicinalproduct": "AZTREONAM LYSINE",
            "drugindication": "CYSTIC FIBROSIS",
            "drugadministrationroute": "055",
            "drugdosagetext": "75 MG, UNK",
            "drugstructuredosageunit": "003",
            "drugdosageform": "INHALATION VAPOUR, LIQUID",
            "drugcharacterization": "1",
            "drugstartdate": "20161013",
            "actiondrug": "4",
            "activesubstance": {
              "activesubstancename": "AZTREONAM LYSINE"
            },
            "drugauthorizationnumb": "050814",
            "drugbatchnumb": "009367",
            "drugstructuredosagenumb": "75"
          }
        ],
        "summary": {
          "narrativeincludeclinical": "CASE EVENT DATE: 2017"
        },
        "patientsex": "1",
        "patientagegroup": "5",
        "patientonsetageunit": "801"
      },
      "sender": {
        "senderorganization": "FDA-Public Use",
        "sendertype": "2"
      },
      "primarysourcecountry": "US",
      "transmissiondate": "20171128",
      "companynumb": "US-GILEAD-2017-0295794",
      "occurcountry": "US",
      "receiptdate": "20170929",
      "transmissiondateformat": "102",
      "receiptdateformat": "102",
      "primarysource": {
        "reportercountry": "US",
        "qualification": "5"
      },
      "duplicate": "1",
      "reporttype": "2",
      "reportduplicate": {
        "duplicatenumb": "US-GILEAD-2017-0295794",
        "duplicatesource": "GILEAD"
      },
      "receiver": {
        "receiverorganization": "FDA",
        "receivertype": "6"
      },
      "serious": "1",
      "safetyreportversion": "1",
      "fulfillexpeditecriteria": "1",
      "safetyreportid": "14026891",
      "receivedateformat": "102"
    },
    {
      "receivedate": "20170929",
      "receivedateformat": "102",
      "patient": {
        "reaction": [
          {
            "reactionmeddrapt": "Blood glucose abnormal",
            "reactionmeddraversionpt": "20.1",
            "reactionoutcome": "3"
          },
          {
            "reactionmeddrapt": "Incorrect dosage administered",
            "reactionmeddraversionpt": "20.1",
            "reactionoutcome": "6"
          }
        ],
        "patientsex": "1",
        "drug": [
          {
            "medicinalproduct": "HUMULIN R",
            "drugindication": "TYPE 2 DIABETES MELLITUS",
            "drugadministrationroute": "065",
            "drugdosagetext": "UNK, UNKNOWN",
            "openfda": {
              "manufacturer_name": [
                "Eli Lilly and Company"
              ],
              "unii": [
                "1Y17CTI5SR"
              ],
              "product_type": [
                "HUMAN OTC DRUG"
              ],
              "rxcui": [
                "311034",
                "311036"
              ],
              "spl_set_id": [
                "b519bd83-038c-4ec5-a231-a51ec5cc291f"
              ],
              "route": [
                "PARENTERAL"
              ],
              "generic_name": [
                "INSULIN HUMAN"
              ],
              "brand_name": [
                "HUMULIN R"
              ],
              "pharm_class_cs": [
                "Insulin [Chemical/Ingredient]"
              ],
              "product_ndc": [
                "0002-8215"
              ],
              "pharm_class_epc": [
                "Insulin [EPC]"
              ],
              "substance_name": [
                "INSULIN HUMAN"
              ],
              "spl_id": [
                "464f0735-d48a-4d10-827c-02142a2f4928"
              ],
              "application_number": [
                "NDA018780"
              ],
              "nui": [
                "N0000004931",
                "N0000175944"
              ],
              "package_ndc": [
                "0002-8215-01",
                "0002-8215-99",
                "0002-8215-17"
              ]
            },
            "drugcharacterization": "1",
            "actiondrug": "4",
            "activesubstance": {
              "activesubstancename": "INSULIN HUMAN"
            },
            "drugauthorizationnumb": "018780",
            "drugdosageform": "INJECTION"
          },
          {
            "medicinalproduct": "METFORMIN",
            "drugindication": "TYPE 2 DIABETES MELLITUS",
            "drugadministrationroute": "065",
            "drugdosagetext": "UNK, UNKNOWN",
            "openfda": {
              "manufacturer_name": [
                "Amneal Pharmaceuticals",
                "Mylan Pharmaceuticals Inc."
              ],
              "unii": [
                "786Z46389E"
              ],
              "product_type": [
                "HUMAN PRESCRIPTION DRUG"
              ],
              "rxcui": [
                "861007",
                "861010",
                "1807917",
                "861004",
                "1807894"
              ],
              "spl_set_id": [
                "45cc673a-e125-4d76-8013-89a089eb0ae9",
                "7e41818c-60e9-4bcf-9586-7bb8d33d5e89"
              ],
              "route": [
                "ORAL"
              ],
              "generic_name": [
                "METFORMIN"
              ],
              "brand_name": [
                "METFORMIN HYDROCHLORIDE"
              ],
              "product_ndc": [
                "0378-6001",
                "0378-6002",
                "65162-177",
                "65162-174",
                "65162-175"
              ],
              "substance_name": [
                "METFORMIN HYDROCHLORIDE"
              ],
              "spl_id": [
                "20e236bb-e297-4f7a-9179-380f1044060e",
                "6fc4d215-86ac-416e-8e9a-5f84b78ef2db"
              ],
              "application_number": [
                "ANDA200690",
                "ANDA077880"
              ],
              "package_ndc": [
                "65162-177-10",
                "65162-174-50",
                "0378-6002-91",
                "0378-6001-91",
                "65162-174-11",
                "65162-174-10",
                "65162-175-10",
                "65162-175-11",
                "65162-177-11",
                "65162-177-50",
                "65162-175-50"
              ]
            },
            "drugcharacterization": "2",
            "activesubstance": {
              "activesubstancename": "METFORMIN HYDROCHLORIDE"
            },
            "actiondrug": "5",
            "drugadditional": "3"
          }
        ]
      },
      "sender": {
        "senderorganization": "FDA-Public Use",
        "sendertype": "2"
      },
      "transmissiondate": "20171128",
      "fulfillexpeditecriteria": "2",
      "occurcountry": "US",
      "receiptdate": "20170929",
      "transmissiondateformat": "102",
      "receiptdateformat": "102",
      "primarysource": {
        "reportercountry": "US",
        "qualification": "3"
      },
      "duplicate": "1",
      "reporttype": "1",
      "reportduplicate": {
        "duplicatenumb": "US-ELI_LILLY_AND_COMPANY-US201709011146",
        "duplicatesource": "ELI LILLY AND CO"
      },
      "receiver": {
        "receiverorganization": "FDA",
        "receivertype": "6"
      },
      "serious": "2",
      "safetyreportversion": "1",
      "companynumb": "US-ELI_LILLY_AND_COMPANY-US201709011146",
      "safetyreportid": "14026947",
      "primarysourcecountry": "US"
    },
    {
      "receivedate": "20170929",
      "receivedateformat": "102",
      "patient": {
        "reaction": [
          {
            "reactionmeddrapt": "Multiple sclerosis relapse",
            "reactionmeddraversionpt": "20.1",
            "reactionoutcome": "6"
          },
          {
            "reactionmeddrapt": "Somnolence",
            "reactionmeddraversionpt": "20.1",
            "reactionoutcome": "6"
          }
        ],
        "patientsex": "2",
        "drug": [
          {
            "medicinalproduct": "TYSABRI",
            "drugindication": "MULTIPLE SCLEROSIS",
            "drugadministrationroute": "065",
            "actiondrug": "5",
            "openfda": {
              "manufacturer_name": [
                "Biogen Inc."
              ],
              "unii": [
                "3JB47N2Q2P"
              ],
              "product_type": [
                "HUMAN PRESCRIPTION DRUG"
              ],
              "rxcui": [
                "603541",
                "477484"
              ],
              "spl_set_id": [
                "c5fdde91-1989-4dd2-9129-4f3323ea2962"
              ],
              "route": [
                "INTRAVENOUS"
              ],
              "generic_name": [
                "NATALIZUMAB"
              ],
              "brand_name": [
                "TYSABRI"
              ],
              "product_ndc": [
                "64406-008"
              ],
              "pharm_class_epc": [
                "Integrin Receptor Antagonist [EPC]"
              ],
              "substance_name": [
                "NATALIZUMAB"
              ],
              "spl_id": [
                "b6ed509c-81c4-4ba6-9560-60bdae9f37e7"
              ],
              "pharm_class_moa": [
                "Integrin Receptor Antagonists [MoA]"
              ],
              "application_number": [
                "BLA125104"
              ],
              "nui": [
                "N0000175774",
                "N0000175775"
              ],
              "package_ndc": [
                "64406-008-01"
              ]
            },
            "drugcharacterization": "1",
            "activesubstance": {
              "activesubstancename": "NATALIZUMAB"
            },
            "drugauthorizationnumb": "125104",
            "drugdosageform": "UNKNOWN",
            "drugadditional": "3"
          }
        ]
      },
      "sender": {
        "senderorganization": "FDA-Public Use",
        "sendertype": "2"
      },
      "transmissiondate": "20171128",
      "fulfillexpeditecriteria": "2",
      "occurcountry": "US",
      "receiptdate": "20170929",
      "transmissiondateformat": "102",
      "receiptdateformat": "102",
      "primarysource": {
        "reportercountry": "US",
        "qualification": "5"
      },
      "duplicate": "1",
      "reporttype": "1",
      "reportduplicate": {
        "duplicatenumb": "US-BIOGEN-2017BI00462854",
        "duplicatesource": "BIOGEN"
      },
      "receiver": {
        "receiverorganization": "FDA",
        "receivertype": "6"
      },
      "serious": "2",
      "safetyreportversion": "1",
      "companynumb": "US-BIOGEN-2017BI00462854",
      "safetyreportid": "14027089",
      "primarysourcecountry": "US"
    },
    {
      "receivedate": "20170929",
      "receivedateformat": "102",
      "patient": {
        "reaction": [
          {
            "reactionmeddrapt": "Metrorrhagia",
            "reactionmeddraversionpt": "20.1",
            "reactionoutcome": "6"
          },
          {
            "reactionmeddrapt": "Rash",
            "reactionmeddraversionpt": "20.1",
            "reactionoutcome": "6"
          }
        ],
        "drug": [
          {
            "medicinalproduct": "COMBIPATCH",
            "drugindication": "PRODUCT USED FOR UNKNOWN INDICATION",
            "drugadministrationroute": "062",
            "drugdosagetext": "HIGHER DOSAGE, UNK",
            "drugbatchnumb": "UNK",
            "drugcharacterization": "1",
            "actiondrug": "6",
            "activesubstance": {
              "activesubstancename": "ESTRADIOL\\NORETHINDRONE ACETATE"
            },
            "drugauthorizationnumb": "020870",
            "drugdosageform": "TRANSDERMAL SYSTEM"
          }
        ]
      },
      "sender": {
        "senderorganization": "FDA-Public Use",
        "sendertype": "2"
      },
      "transmissiondate": "20171128",
      "fulfillexpeditecriteria": "2",
      "occurcountry": "US",
      "receiptdate": "20170929",
      "transmissiondateformat": "102",
      "receiptdateformat": "102",
      "primarysource": {
        "reportercountry": "US",
        "qualification": "5"
      },
      "duplicate": "1",
      "reporttype": "1",
      "reportduplicate": {
        "duplicatenumb": "US-NOVEN PHARMACEUTICALS, INC.-US2017000294",
        "duplicatesource": "NOVEN"
      },
      "receiver": {
        "receiverorganization": "FDA",
        "receivertype": "6"
      },
      "serious": "2",
      "safetyreportversion": "1",
      "companynumb": "US-NOVEN PHARMACEUTICALS, INC.-US2017000294",
      "safetyreportid": "14027166",
      "primarysourcecountry": "US"
    },
    {
      "receivedate": "20170929",
      "receivedateformat": "102",
      "patient": {
        "reaction": [
          {
            "reactionmeddrapt": "Product adhesion issue",
            "reactionmeddraversionpt": "20.1",
            "reactionoutcome": "6"
          },
          {
            "reactionmeddrapt": "Product quality issue",
            "reactionmeddraversionpt": "20.1",
            "reactionoutcome": "6"
          },
          {
            "reactionmeddrapt": "No adverse event",
            "reactionmeddraversionpt": "20.1",
            "reactionoutcome": "6"
          },
          {
            "reactionmeddrapt": "Product storage error",
            "reactionmeddraversionpt": "20.1",
            "reactionoutcome": "3"
          }
        ],
        "drug": [
          {
            "medicinalproduct": "COMBIPATCH",
            "drugindication": "NIGHT SWEATS",
            "actiondrug": "4",
            "drugcharacterization": "1",
            "activesubstance": {
              "activesubstancename": "ESTRADIOL\\NORETHINDRONE ACETATE"
            },
            "drugauthorizationnumb": "020870",
            "drugdosageform": "TRANSDERMAL SYSTEM"
          },
          {
            "drugstartdateformat": "602",
            "medicinalproduct": "COMBIPATCH",
            "drugindication": "HOT FLUSH",
            "drugadministrationroute": "062",
            "drugdosagetext": "50/140 MG, UNKNOWN",
            "drugdosageform": "TRANSDERMAL SYSTEM",
            "drugcharacterization": "1",
            "drugstartdate": "2015",
            "actiondrug": "4",
            "activesubstance": {
              "activesubstancename": "ESTRADIOL\\NORETHINDRONE ACETATE"
            },
            "drugauthorizationnumb": "020870",
            "drugbatchnumb": "80678"
          }
        ],
        "summary": {
          "narrativeincludeclinical": "CASE EVENT DATE: 2015"
        }
      },
      "sender": {
        "senderorganization": "FDA-Public Use",
        "sendertype": "2"
      },
      "transmissiondate": "20171128",
      "fulfillexpeditecriteria": "2",
      "occurcountry": "US",
      "receiptdate": "20170929",
      "transmissiondateformat": "102",
      "receiptdateformat": "102",
      "primarysource": {
        "reportercountry": "US",
        "qualification": "2"
      },
      "duplicate": "1",
      "reporttype": "1",
      "reportduplicate": {
        "duplicatenumb": "US-NOVEN PHARMACEUTICALS, INC.-US2017000412",
        "duplicatesource": "NOVEN"
      },
      "receiver": {
        "receiverorganization": "FDA",
        "receivertype": "6"
      },
      "serious": "2",
      "safetyreportversion": "1",
      "companynumb": "US-NOVEN PHARMACEUTICALS, INC.-US2017000412",
      "safetyreportid": "14027236",
      "primarysourcecountry": "US"
    },
    {
      "receivedate": "20170929",
      "receivedateformat": "102",
      "patient": {
        "reaction": [
          {
            "reactionmeddrapt": "Death",
            "reactionmeddraversionpt": "20.1",
            "reactionoutcome": "5"
          }
        ],
        "patientsex": "1",
        "drug": [
          {
            "medicinalproduct": "EXTRANEAL",
            "drugindication": "PERITONEAL DIALYSIS",
            "drugadministrationroute": "033",
            "actiondrug": "6",
            "openfda": {
              "manufacturer_name": [
                "Baxter Healthcare Corporation"
              ],
              "unii": [
                "2NX48Z0A9G",
                "M4I0D6VV5M",
                "02F3473H9O",
                "451W47IQ8X",
                "TU7HW0W0QT"
              ],
              "product_type": [
                "HUMAN PRESCRIPTION DRUG"
              ],
              "rxcui": [
                "1100742",
                "1100746"
              ],
              "spl_set_id": [
                "d5b85158-b0d6-4855-9d07-8d1b3ad9ab71"
              ],
              "route": [
                "INTRAPERITONEAL"
              ],
              "generic_name": [
                "ICODEXTRIN, SODIUM CHLORIDE, SODIUM LACTATE, CALCIUM CHLORIDE, MAGNESIUM CHLORIDE"
              ],
              "brand_name": [
                "EXTRANEAL"
              ],
              "product_ndc": [
                "0941-0679"
              ],
              "substance_name": [
                "SODIUM CHLORIDE",
                "CALCIUM CHLORIDE",
                "ICODEXTRIN",
                "SODIUM LACTATE",
                "MAGNESIUM CHLORIDE"
              ],
              "spl_id": [
                "3ab6f82d-5a8d-4eaf-82b8-15be5ec3a1df"
              ],
              "application_number": [
                "NDA021321"
              ],
              "package_ndc": [
                "0941-0679-53",
                "0941-0679-52",
                "0941-0679-06",
                "0941-0679-05"
              ]
            },
            "drugcharacterization": "1",
            "activesubstance": {
              "activesubstancename": "CALCIUM CHLORIDE\\ICODEXTRIN\\MAGNESIUM CHLORIDE\\SODIUM CHLORIDE\\SODIUM LACTATE"
            },
            "drugauthorizationnumb": "021321",
            "drugdosageform": "SOLUTION FOR PERITONEAL DIALYSIS"
          },
          {
            "medicinalproduct": "DIANEAL PD-2 PERITONEAL DIALYSIS SOLUTION WITH DEXTROSE",
            "drugindication": "PERITONEAL DIALYSIS",
            "drugadministrationroute": "033",
            "actiondrug": "6",
            "drugcharacterization": "1",
            "activesubstance": {
              "activesubstancename": "CALCIUM CHLORIDE\\DEXTROSE\\MAGNESIUM CHLORIDE\\SODIUM CHLORIDE\\SODIUM LACTATE"
            },
            "drugauthorizationnumb": "017512",
            "drugdosageform": "SOLUTION FOR PERITONEAL DIALYSIS"
          },
          {
            "medicinalproduct": "DIANEAL PD-2 PERITONEAL DIALYSIS SOLUTION WITH DEXTROSE",
            "drugindication": "PERITONEAL DIALYSIS",
            "drugadministrationroute": "033",
            "actiondrug": "6",
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
      "transmissiondate": "20171128",
      "companynumb": "HK-BAXTER-2017BAX034144",
      "occurcountry": "HK",
      "receiptdate": "20170929",
      "transmissiondateformat": "102",
      "receiptdateformat": "102",
      "primarysource": {
        "reportercountry": "HK",
        "qualification": "5"
      },
      "duplicate": "1",
      "seriousnessdeath": "1",
      "reporttype": "1",
      "reportduplicate": {
        "duplicatenumb": "HK-BAXTER-2017BAX034144",
        "duplicatesource": "BAXTER"
      },
      "receiver": {
        "receiverorganization": "FDA",
        "receivertype": "6"
      },
      "serious": "1",
      "safetyreportversion": "1",
      "fulfillexpeditecriteria": "1",
      "safetyreportid": "14027382",
      "primarysourcecountry": "HK"
    },
    {
      "receivedate": "20170825",
      "receivedateformat": "102",
      "patient": {
        "reaction": [
          {
            "reactionmeddrapt": "Vision blurred",
            "reactionmeddraversionpt": "20.1",
            "reactionoutcome": "2"
          }
        ],
        "patientsex": "2",
        "drug": [
          {
            "medicinalproduct": "NEO AND POLYMYXIN B SULFATES + DEX",
            "drugindication": "PRODUCT USED FOR UNKNOWN INDICATION",
            "drugcharacterization": "1",
            "drugadministrationroute": "047",
            "drugseparatedosagenumb": "1",
            "drugstructuredosageunit": "031",
            "drugdosagetext": "3 GTT, QD",
            "drugstructuredosagenumb": "3",
            "drugintervaldosageunitnumb": "1",
            "actiondrug": "1",
            "activesubstance": {
              "activesubstancename": "DEXAMETHASONE\\NEOMYCIN SULFATE\\POLYMYXIN B SULFATE"
            },
            "drugintervaldosagedefinition": "804",
            "drugauthorizationnumb": "050023",
            "drugrecurreadministration": "3",
            "drugbatchnumb": "282340F",
            "drugadditional": "1"
          }
        ]
      },
      "sender": {
        "senderorganization": "FDA-Public Use",
        "sendertype": "2"
      },
      "transmissiondate": "20171127",
      "fulfillexpeditecriteria": "2",
      "occurcountry": "US",
      "receiptdate": "20170928",
      "transmissiondateformat": "102",
      "receiptdateformat": "102",
      "primarysource": {
        "reportercountry": "US",
        "qualification": "5"
      },
      "duplicate": "1",
      "reporttype": "1",
      "reportduplicate": {
        "duplicatenumb": "PHEH2017US026843",
        "duplicatesource": "SANDOZ"
      },
      "receiver": {
        "receiverorganization": "FDA",
        "receivertype": "6"
      },
      "serious": "2",
      "safetyreportversion": "4",
      "companynumb": "PHEH2017US026843",
      "safetyreportid": "13905257",
      "primarysourcecountry": "US"
    },
    {
      "receivedate": "20170825",
      "receivedateformat": "102",
      "patient": {
        "reaction": [
          {
            "reactionmeddrapt": "Sudden death",
            "reactionmeddraversionpt": "20.1",
            "reactionoutcome": "5"
          }
        ],
        "patientonsetage": "81",
        "drug": [
          {
            "drugstartdateformat": "102",
            "medicinalproduct": "ENTRESTO",
            "drugindication": "CARDIAC FAILURE CHRONIC",
            "drugadministrationroute": "065",
            "drugdosagetext": "50 MG, UNK (FREQUENCY 2)",
            "drugstructuredosageunit": "003",
            "openfda": {
              "manufacturer_name": [
                "Novartis Pharmaceuticals Corporation"
              ],
              "unii": [
                "17ERJ0MKGI",
                "80M03YXJ7I"
              ],
              "product_type": [
                "HUMAN PRESCRIPTION DRUG"
              ],
              "rxcui": [
                "1656349",
                "1656340",
                "1656351",
                "1656356",
                "1656346",
                "1656354"
              ],
              "spl_set_id": [
                "000dc81d-ab91-450c-8eae-8eb74e72296f"
              ],
              "route": [
                "ORAL"
              ],
              "generic_name": [
                "SACUBITRIL AND VALSARTAN"
              ],
              "brand_name": [
                "ENTRESTO"
              ],
              "product_ndc": [
                "0078-0696",
                "0078-0659",
                "0078-0777"
              ],
              "pharm_class_epc": [
                "Angiotensin 2 Receptor Blocker [EPC]"
              ],
              "substance_name": [
                "SACUBITRIL",
                "VALSARTAN"
              ],
              "spl_id": [
                "fe25b914-9d1e-4b86-bab5-947c78a30fcb"
              ],
              "pharm_class_moa": [
                "Angiotensin 2 Receptor Antagonists [MoA]"
              ],
              "application_number": [
                "NDA207620"
              ],
              "nui": [
                "N0000175561",
                "N0000000070"
              ],
              "package_ndc": [
                "0078-0777-67",
                "0078-0659-35",
                "0078-0777-35",
                "0078-0696-61",
                "0078-0777-61",
                "0078-0696-35",
                "0078-0696-67",
                "0078-0659-61",
                "0078-0777-20",
                "0078-0696-20",
                "0078-0659-67",
                "0078-0659-20"
              ]
            },
            "drugcharacterization": "1",
            "drugstartdate": "20170622",
            "actiondrug": "6",
            "activesubstance": {
              "activesubstancename": "SACUBITRIL\\VALSARTAN"
            },
            "drugauthorizationnumb": "207620",
            "drugdosageform": "TABLET",
            "drugstructuredosagenumb": "50"
          }
        ],
        "summary": {
          "narrativeincludeclinical": "CASE EVENT DATE: 20170811"
        },
        "patientsex": "1",
        "patientonsetageunit": "801"
      },
      "sender": {
        "senderorganization": "FDA-Public Use",
        "sendertype": "2"
      },
      "transmissiondate": "20171127",
      "companynumb": "PHHY2017AR125242",
      "occurcountry": "AR",
      "receiptdate": "20170927",
      "transmissiondateformat": "102",
      "receiptdateformat": "102",
      "primarysource": {
        "reportercountry": "AR",
        "qualification": "5"
      },
      "duplicate": "1",
      "seriousnessdeath": "1",
      "reporttype": "2",
      "reportduplicate": {
        "duplicatenumb": "PHHY2017AR125242",
        "duplicatesource": "NOVARTIS"
      },
      "receiver": {
        "receiverorganization": "FDA",
        "receivertype": "6"
      },
      "serious": "1",
      "safetyreportversion": "2",
      "fulfillexpeditecriteria": "1",
      "safetyreportid": "13905341",
      "primarysourcecountry": "AR"
    },
    {
      "seriousnessother": "1",
      "reportduplicate": {
        "duplicatenumb": "US-WEST-WARD PHARMACEUTICALS CORP.-US-H14001-17-03090",
        "duplicatesource": "WESTWARD"
      },
      "safetyreportversion": "1",
      "receiptdate": "20170825",
      "seriousnesslifethreatening": "1",
      "duplicate": "1",
      "seriousnesshospitalization": "1",
      "receivedate": "20170825",
      "patient": {
        "reaction": [
          {
            "reactionmeddrapt": "Toxicity to various agents",
            "reactionmeddraversionpt": "20.1",
            "reactionoutcome": "1"
          },
          {
            "reactionmeddrapt": "Respiratory failure",
            "reactionmeddraversionpt": "20.1",
            "reactionoutcome": "1"
          },
          {
            "reactionmeddrapt": "Brain oedema",
            "reactionmeddraversionpt": "20.1",
            "reactionoutcome": "1"
          },
          {
            "reactionmeddrapt": "Accidental overdose",
            "reactionmeddraversionpt": "20.1",
            "reactionoutcome": "1"
          },
          {
            "reactionmeddrapt": "Accidental exposure to product",
            "reactionmeddraversionpt": "20.1",
            "reactionoutcome": "1"
          }
        ],
        "patientonsetage": "10",
        "patientsex": "2",
        "patientonsetageunit": "802",
        "drug": [
          {
            "medicinalproduct": "OXYCODONE",
            "drugindication": "PRODUCT USED FOR UNKNOWN INDICATION",
            "drugadministrationroute": "048",
            "openfda": {
              "manufacturer_name": [
                "Amneal Pharmaceuticals LLC",
                "Collegium Pharmaceutical, Inc."
              ],
              "unii": [
                "C1ENJ2TE6C",
                "CD35PMG570"
              ],
              "product_type": [
                "HUMAN PRESCRIPTION DRUG"
              ],
              "rxcui": [
                "1049683",
                "1049621",
                "1790533",
                "1791558",
                "1791574",
                "1791576",
                "1049618",
                "1791567",
                "1049686",
                "1791560",
                "1791580",
                "1049611",
                "1791582",
                "1790527",
                "1791569"
              ],
              "spl_set_id": [
                "094b64b3-cd32-4de5-afb6-ea00d9caad74",
                "b0a5ded2-8ee2-49ca-a86c-2b28ae40f60c"
              ],
              "route": [
                "ORAL"
              ],
              "generic_name": [
                "OXYCODONE"
              ],
              "brand_name": [
                "OXYCODONE HYDROCHLORIDE",
                "XTAMPZA ER"
              ],
              "product_ndc": [
                "65162-047",
                "24510-140",
                "65162-051",
                "65162-050",
                "24510-110",
                "24510-115",
                "65162-048",
                "65162-049",
                "24510-120",
                "24510-130"
              ],
              "pharm_class_epc": [
                "Opioid Agonist [EPC]"
              ],
              "substance_name": [
                "OXYCODONE HYDROCHLORIDE",
                "OXYCODONE"
              ],
              "spl_id": [
                "b8168438-f108-4a97-bdf4-f2b96e1368ba",
                "9670cffb-1222-4756-8a24-c69d5b8dc435"
              ],
              "pharm_class_moa": [
                "Full Opioid Agonists [MoA]"
              ],
              "application_number": [
                "ANDA203638",
                "NDA208090"
              ],
              "nui": [
                "N0000175690",
                "N0000175684"
              ],
              "package_ndc": [
                "65162-049-50",
                "24510-120-10",
                "65162-048-03",
                "65162-051-10",
                "24510-130-10",
                "65162-047-10",
                "65162-048-10",
                "65162-051-50",
                "65162-050-03",
                "65162-047-50",
                "24510-115-10",
                "24510-110-10",
                "24510-140-20",
                "65162-049-25",
                "65162-049-03",
                "24510-120-20",
                "65162-050-25",
                "24510-130-20",
                "65162-047-03",
                "65162-051-03",
                "65162-048-25",
                "65162-051-25",
                "65162-047-25",
                "65162-050-10",
                "24510-140-10",
                "24510-115-20",
                "24510-110-20",
                "65162-049-10"
              ]
            },
            "drugcharacterization": "1",
            "activesubstance": {
              "activesubstancename": "OXYCODONE"
            },
            "drugauthorizationnumb": "203208"
          }
        ]
      },
      "occurcountry": "US",
      "reporttype": "1",
      "companynumb": "US-WEST-WARD PHARMACEUTICALS CORP.-US-H14001-17-03090",
      "safetyreportid": "13905529",
      "sender": {
        "senderorganization": "FDA-Public Use",
        "sendertype": "2"
      },
      "transmissiondate": "20171127",
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
        "literaturereference": "DURAN D,MESSINA R,BESLOW L,MONTEJO J,KARIMY J,CHARUTA GAVANKAR F,SHERIDAN A,SZE G,YARMAN Y,DILUNA M,KAHLE K. MALIGNANT CEREBELLAR EDEMA SUBSEQUENT TO ACCIDENTAL PRESCRIPTION OPIOID INTOXICATION IN CHILDREN. FRONTIERS IN NEUROLOGY 2017;8:1-6.",
        "reportercountry": "US",
        "qualification": "1"
      },
      "primarysourcecountry": "US"
    },
    {
      "receivedate": "20170825",
      "receivedateformat": "102",
      "patient": {
        "reaction": [
          {
            "reactionmeddrapt": "Fatigue",
            "reactionmeddraversionpt": "20.1",
            "reactionoutcome": "2"
          },
          {
            "reactionmeddrapt": "Hepatic function abnormal",
            "reactionmeddraversionpt": "20.1",
            "reactionoutcome": "2"
          },
          {
            "reactionmeddrapt": "Liver function test increased",
            "reactionmeddraversionpt": "20.1",
            "reactionoutcome": "2"
          }
        ],
        "patientsex": "2",
        "patientagegroup": "5",
        "drug": [
          {
            "medicinalproduct": "BETAFERON",
            "drugindication": "MULTIPLE SCLEROSIS",
            "drugcharacterization": "1",
            "drugadministrationroute": "058",
            "drugseparatedosagenumb": "1",
            "drugdosagetext": "UNK, QOD",
            "drugbatchnumb": "UNKNOWN",
            "drugintervaldosageunitnumb": "2",
            "actiondrug": "1",
            "activesubstance": {
              "activesubstancename": "INTERFERON BETA-1B"
            },
            "drugintervaldosagedefinition": "804",
            "drugauthorizationnumb": "103471",
            "drugdosageform": "POWDER AND SOLVENT FOR SOLUTION FOR INJECTION",
            "drugadditional": "1"
          }
        ]
      },
      "sender": {
        "senderorganization": "FDA-Public Use",
        "sendertype": "2"
      },
      "transmissiondate": "20171127",
      "companynumb": "JP-BAYER-2017-151644",
      "occurcountry": "JP",
      "receiptdate": "20170825",
      "transmissiondateformat": "102",
      "receiptdateformat": "102",
      "primarysource": {
        "reportercountry": "JP",
        "qualification": "3"
      },
      "duplicate": "1",
      "reporttype": "1",
      "seriousnesshospitalization": "1",
      "reportduplicate": {
        "duplicatenumb": "JP-BAYER-2017-151644",
        "duplicatesource": "BAYER"
      },
      "receiver": {
        "receiverorganization": "FDA",
        "receivertype": "6"
      },
      "serious": "1",
      "safetyreportversion": "1",
      "fulfillexpeditecriteria": "2",
      "safetyreportid": "13905550",
      "primarysourcecountry": "JP"
    },
    {
      "receivedate": "20170828",
      "receivedateformat": "102",
      "patient": {
        "reaction": [
          {
            "reactionmeddrapt": "Death",
            "reactionmeddraversionpt": "20.1",
            "reactionoutcome": "5"
          }
        ],
        "patientonsetage": "69",
        "drug": [
          {
            "medicinalproduct": "DIANEAL",
            "drugindication": "PERITONEAL DIALYSIS",
            "drugadministrationroute": "033",
            "actiondrug": "6",
            "drugbatchnumb": "W27P7",
            "drugcharacterization": "1",
            "activesubstance": {
              "activesubstancename": "CALCIUM CHLORIDE\\DEXTROSE\\MAGNESIUM CHLORIDE\\SODIUM CHLORIDE\\SODIUM LACTATE"
            },
            "drugauthorizationnumb": "017512",
            "drugdosageform": "SOLUTION FOR PERITONEAL DIALYSIS"
          },
          {
            "medicinalproduct": "EXTRANEAL",
            "drugindication": "PERITONEAL DIALYSIS",
            "drugadministrationroute": "033",
            "actiondrug": "6",
            "openfda": {
              "manufacturer_name": [
                "Baxter Healthcare Corporation"
              ],
              "unii": [
                "2NX48Z0A9G",
                "M4I0D6VV5M",
                "02F3473H9O",
                "451W47IQ8X",
                "TU7HW0W0QT"
              ],
              "product_type": [
                "HUMAN PRESCRIPTION DRUG"
              ],
              "rxcui": [
                "1100742",
                "1100746"
              ],
              "spl_set_id": [
                "d5b85158-b0d6-4855-9d07-8d1b3ad9ab71"
              ],
              "route": [
                "INTRAPERITONEAL"
              ],
              "generic_name": [
                "ICODEXTRIN, SODIUM CHLORIDE, SODIUM LACTATE, CALCIUM CHLORIDE, MAGNESIUM CHLORIDE"
              ],
              "brand_name": [
                "EXTRANEAL"
              ],
              "product_ndc": [
                "0941-0679"
              ],
              "substance_name": [
                "SODIUM CHLORIDE",
                "CALCIUM CHLORIDE",
                "ICODEXTRIN",
                "SODIUM LACTATE",
                "MAGNESIUM CHLORIDE"
              ],
              "spl_id": [
                "3ab6f82d-5a8d-4eaf-82b8-15be5ec3a1df"
              ],
              "application_number": [
                "NDA021321"
              ],
              "package_ndc": [
                "0941-0679-53",
                "0941-0679-52",
                "0941-0679-06",
                "0941-0679-05"
              ]
            },
            "drugbatchnumb": "W26L4",
            "drugcharacterization": "1",
            "activesubstance": {
              "activesubstancename": "CALCIUM CHLORIDE\\ICODEXTRIN\\MAGNESIUM CHLORIDE\\SODIUM CHLORIDE\\SODIUM LACTATE"
            },
            "drugauthorizationnumb": "021321",
            "drugdosageform": "SOLUTION FOR PERITONEAL DIALYSIS"
          }
        ],
        "summary": {
          "narrativeincludeclinical": "CASE EVENT DATE: 20170822"
        },
        "patientsex": "1",
        "patientonsetageunit": "801"
      },
      "sender": {
        "senderorganization": "FDA-Public Use",
        "sendertype": "2"
      },
      "transmissiondate": "20171127",
      "companynumb": "AU-BAXTER-2017BAX030881",
      "occurcountry": "AU",
      "receiptdate": "20170828",
      "transmissiondateformat": "102",
      "receiptdateformat": "102",
      "primarysource": {
        "reportercountry": "AU",
        "qualification": "3"
      },
      "duplicate": "1",
      "seriousnessdeath": "1",
      "reporttype": "1",
      "reportduplicate": {
        "duplicatenumb": "AU-BAXTER-2017BAX030881",
        "duplicatesource": "BAXTER"
      },
      "receiver": {
        "receiverorganization": "FDA",
        "receivertype": "6"
      },
      "serious": "1",
      "safetyreportversion": "1",
      "fulfillexpeditecriteria": "1",
      "safetyreportid": "13911526",
      "primarysourcecountry": "AU"
    },
    {
      "receivedate": "20170828",
      "receivedateformat": "102",
      "patient": {
        "reaction": [
          {
            "reactionmeddrapt": "Pain",
            "reactionmeddraversionpt": "20.1",
            "reactionoutcome": "6"
          },
          {
            "reactionmeddrapt": "Headache",
            "reactionmeddraversionpt": "20.1",
            "reactionoutcome": "6"
          },
          {
            "reactionmeddrapt": "Drug ineffective",
            "reactionmeddraversionpt": "20.1",
            "reactionoutcome": "6"
          }
        ],
        "patientweight": "23.1",
        "patientsex": "1",
        "drug": [
          {
            "drugstartdateformat": "102",
            "medicinalproduct": "NUTROPIN AQ",
            "drugadministrationroute": "058",
            "drugdosagetext": "DILUTED WITH 1 ML",
            "drugstructuredosageunit": "003",
            "drugcharacterization": "1",
            "drugstartdate": "20001208",
            "activesubstance": {
              "activesubstancename": "SOMATROPIN"
            },
            "drugauthorizationnumb": "020522",
            "drugstructuredosagenumb": "18",
            "actiondrug": "5",
            "drugadditional": "3"
          },
          {
            "drugstartdateformat": "102",
            "medicinalproduct": "NUTROPIN AQ",
            "drugindication": "GROWTH HORMONE DEFICIENCY",
            "drugadministrationroute": "065",
            "actiondrug": "5",
            "drugcharacterization": "1",
            "drugstartdate": "19970926",
            "activesubstance": {
              "activesubstancename": "SOMATROPIN"
            },
            "drugauthorizationnumb": "020522",
            "drugdosageform": "SOLUTION FOR INJECTION",
            "drugadditional": "3"
          },
          {
            "drugstartdateformat": "610",
            "medicinalproduct": "NUTROPIN AQ",
            "drugindication": "DWARFISM",
            "drugadministrationroute": "065",
            "drugenddateformat": "610",
            "drugdosagetext": "CHANGED TO POWDER FORM",
            "drugcharacterization": "1",
            "drugstartdate": "200007",
            "activesubstance": {
              "activesubstancename": "SOMATROPIN"
            },
            "drugauthorizationnumb": "020522",
            "actiondrug": "5",
            "drugenddate": "200008",
            "drugadditional": "3"
          }
        ]
      },
      "sender": {
        "senderorganization": "FDA-Public Use",
        "sendertype": "2"
      },
      "transmissiondate": "20171127",
      "fulfillexpeditecriteria": "2",
      "occurcountry": "US",
      "receiptdate": "20170828",
      "transmissiondateformat": "102",
      "receiptdateformat": "102",
      "primarysource": {
        "reportercountry": "US",
        "qualification": "1"
      },
      "duplicate": "1",
      "reporttype": "1",
      "reportduplicate": {
        "duplicatenumb": "US-ROCHE-1141620",
        "duplicatesource": "ROCHE"
      },
      "receiver": {
        "receiverorganization": "FDA",
        "receivertype": "6"
      },
      "serious": "2",
      "safetyreportversion": "1",
      "companynumb": "US-ROCHE-1141620",
      "safetyreportid": "13911557",
      "primarysourcecountry": "US"
  },{
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
