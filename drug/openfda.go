package drug

const (
	ProductTypeHumanPerscriptionDrug = "HUMAN PRESCRIPTION DRUG"
)

type OpenFDA struct {
	ID    string `json:"_id" bson:"_id"`
	SetID string `json:"set_id,omitempty" bson:"set_id,omitempty"`

	SPLID    []string `json:"spl_id,omitempty" bson:"-"`
	SPLSetID []string `json:"spl_set_id,omitempty" bson:"-"`

	ApplicationNumber []string `json:"application_number,omitempty" bson:"appnumb,omitempty"`

	BrandName        []string `json:"brand_name,omitempty" bson:"bn,omitempty"`
	GenericName      []string `json:"generic_name,omitempty" bson:"gn,omitempty"`
	ManufacturerName []string `json:"manufacturer_name,omitempty" bson:"mn,omitempty"`
	SubstanceName    []string `json:"substance_name,omitempty" bson:"sn,omitempty"`

	NUI []string `json:"nui,omitempty" bson:"nui,omitempty"`

	PackageNDC []string `json:"package_ndc,omitempty" bson:"pandc,omitempty"`
	ProductNDC []string `json:"product_ndc,omitempty" bson:"prndc,omitempty"`

	PharmClassCS  []string `json:"pharm_class_cs,omitempty" bson:"cs,omitempty"`
	PharmClassEPC []string `json:"pharm_class_epc,omitempty" bson:"epc,omitempty"`
	PharmClassMOA []string `json:"pharm_class_moa,omitempty" bson:"moa,omitempty"`
	PharmClassPE  []string `json:"pharm_class_pe,omitempty" bson:"pe,omitempty"`

	ProductType []string `json:"product_type,omitempty" bson:"pt,omitempty"`
	Route       []string `json:"route,omitempty" bson:"route,omitempty"`
	RxCUI       []string `json:"rxcui,omitempty" bson:"rxcui,omitempty"`
	UNII        []string `json:"unii,omitempty" bson:"unii,omitempty"`
	UPC         []string `json:"upc,omitempty" bson:"upc,omitempty"`

	// Undocumented Fields Found In Responses
	IsOriginalPackager []bool `json:"is_original_packager,omitempty" bson:"is_original_packager,omitempty"`
}

func (openFDA OpenFDA) IsHumanPresecriptionDrug() bool {
	for _, tp := range openFDA.ProductType {
		if tp == ProductTypeHumanPerscriptionDrug {
			return true
		}
	}
	return false
}
