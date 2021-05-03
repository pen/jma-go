package area

var ParentClassOf = map[string]string{ //nolint: gochecknoglobals
	"class20": "class15",
	"class15": "class10",
	"class10": "office",
	"office":  "center",
}

type Area struct {
	Class      string
	Code       string
	ParentCode string
	Name       string
	NameEn     string
	NameKana   string
	OfficeName string
}
