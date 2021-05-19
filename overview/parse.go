package overview

import (
	"encoding/json"
	"fmt"
	"io"
	"time"
)

type rawOverview struct {
	PublishingOffice string    `json:"publishingOffice"`
	ReportDateTime   time.Time `json:"reportDateTime"`
	TargetArea       string    `json:"targetArea"`
	HeadlineText     string    `json:"headlineText"`
	Text             string    `json:"text"`
}

func Parse(reader io.Reader) (*Overview, error) {
	ro := rawOverview{}

	if err := json.NewDecoder(reader).Decode(&ro); err != nil {
		return nil, fmt.Errorf("decode error: %w", err)
	}

	overview := Overview{
		OfficeName: ro.PublishingOffice,
		AreaName:   ro.TargetArea,
		ReportedAt: ro.ReportDateTime,
		Text:       ro.Text,
	}

	if ro.HeadlineText != "" {
		overview.Headline = &ro.HeadlineText
	}

	return &overview, nil
}
