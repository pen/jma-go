package week

import (
	"time"
)

type WeeklyOverview struct {
	OfficeName string
	AreaName   string
	ReportedAt time.Time
	HeadTitle  string
	Text       string
}
