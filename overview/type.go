package overview

import (
	"time"
)

type Overview struct {
	OfficeName string
	AreaName   string
	ReportedAt time.Time
	Headline   *string
	Text       string
}
