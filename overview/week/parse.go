package week

import (
	"encoding/json"
	"fmt"
	"io"
	"time"
)

type rawWeek struct {
	PublishingOffice string    `json:"publishingOffice"`
	ReportDateTime   time.Time `json:"reportDateTime"`
	TargetArea       string    `json:"targetArea"`
	HeadTitle        string    `json:"headTitle"`
	Text             string    `json:"text"`
}

func Parse(reader io.Reader) (*WeeklyOverview, error) {
	rw := rawWeek{}

	if err := json.NewDecoder(reader).Decode(&rw); err != nil {
		return nil, fmt.Errorf("decode error: %w", err)
	}

	return &WeeklyOverview{
		OfficeName: rw.PublishingOffice,
		ReportedAt: rw.ReportDateTime,
		HeadTitle:  rw.HeadTitle,
		Text:       rw.Text,
	}, nil
}
