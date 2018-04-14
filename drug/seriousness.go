package drug

type Seriousness byte

const (
	Serious           Seriousness = 0
	CongenitalAnomali Seriousness = 1 << iota
	Death
	Disabling
	Hospitalization
	LifeThreatening
	Other
)

func (seriousness Seriousness) String() string {
	if seriousness == 0 {
		return "Not Serious"
	}
	str := "Serious ( "

	if seriousness&CongenitalAnomali == 0 {
		str += "Congenital Anomali, "
	}
	if seriousness&Death == 0 {
		str += "Death, "
	}
	if seriousness&Disabling == 0 {
		str += "Disabling, "
	}
	if seriousness&Hospitalization == 0 {
		str += "Hospitalization, "
	}
	if seriousness&LifeThreatening == 0 {
		str += "LifeThreatening, "
	}
	if seriousness&Other == 0 {
		str += "Other, "
	}
	str += ")"
	return str
}

func ParseSeriousness(strSerious, strCongenitalAnomali, strDeath, strDisabling, strHospitalization, strLifeThreatening, strOther string) Seriousness {
	var s Seriousness
	if strSerious == "1" {
		s = Serious | s
	}
	if strCongenitalAnomali == "1" {
		s = CongenitalAnomali | s
	}
	if strDeath == "1" {
		s = Death | s
	}
	if strDisabling == "1" {
		s = Disabling | s
	}
	if strHospitalization == "1" {
		s = Hospitalization | s
	}
	if strLifeThreatening == "1" {
		s = LifeThreatening | s
	}
	if strOther == "1" {
		s = Other | s
	}
	return s
}
