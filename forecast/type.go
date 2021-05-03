package forecast

import (
	"time"
)

type Forecast struct {
	Area          Area
	ReportedAt    time.Time // 予報をいつしたか
	ComesAt       time.Time // いつの天気を予報しているか
	Weather       *Weather
	Wind          *Wind
	Wave          *Wave
	Precipitation *Precipitation
	Temperature   *Temperature
}

type Area struct {
	Code string
	Name string
}

type Weather struct {
	Code string
	Text *string
}

type Wind struct {
	Text string
}

type Wave struct {
	Text string
}

type Precipitation struct {
	Probability int
	Reliability *string
}

type Temperature struct {
	Base *int
	Max  *TemperatureRange
	Min  *TemperatureRange
}

type TemperatureRange struct {
	Base  int
	Upper int
	Lower int
}
