package area

import (
	"encoding/json"
	"fmt"
	"io"
)

type rawArea struct {
	Name       string   `json:"name"`
	EnName     string   `json:"enName"`
	Kana       string   `json:"kana"`
	OfficeName string   `json:"officeName"`
	Parent     string   `json:"parent"`
	Children   []string `json:"children"`
}

func Parse(reader io.Reader) ([]*Area, error) {
	var rawAreaMaps struct {
		Centers  map[string]*rawArea `json:"centers"`
		Offices  map[string]*rawArea `json:"offices"`
		Class10s map[string]*rawArea `json:"class10s"`
		Class15s map[string]*rawArea `json:"class15s"`
		Class20s map[string]*rawArea `json:"class20s"`
	}

	if err := json.NewDecoder(reader).Decode(&rawAreaMaps); err != nil {
		return nil, fmt.Errorf("decode error: %w", err)
	}

	var areas []*Area

	addAreas(&areas, rawAreaMaps.Centers, "center")
	addAreas(&areas, rawAreaMaps.Offices, "office")
	addAreas(&areas, rawAreaMaps.Class10s, "class10")
	addAreas(&areas, rawAreaMaps.Class15s, "class15")
	addAreas(&areas, rawAreaMaps.Class20s, "class20")

	return areas, nil
}

func addAreas(areas *[]*Area, rawAreaMap map[string]*rawArea, class string) {
	for code, raw := range rawAreaMap {
		*areas = append(*areas, &Area{
			Class:      class,
			Code:       code,
			ParentCode: raw.Parent,
			Name:       raw.Name,
			NameEn:     raw.EnName,
			NameKana:   raw.Kana,
			OfficeName: raw.OfficeName,
		})
	}
}
