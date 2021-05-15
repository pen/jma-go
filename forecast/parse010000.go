package forecast

import (
	"encoding/json"
	"fmt"
	"io"
	"time"
)

type rawOfficeForecast struct {
	OfficeCode string           `json:"officeCode"`
	Name       string           `json:"name"`
	Srf        rawRangeForecast `json:"srf"` // short range forecast?
	Week       rawRangeForecast `json:"week"`
}

type rawRangeForecast struct {
	PublishingOffice string             `json:"publishingOffice"`
	ReportDateTime   time.Time          `json:"reportDateTime"`
	TimeSeries       []*rawAreaTimeItem `json:"timeSeries"`
	TempAverage      *rawAreaItem       `json:"tempAverage"`
	PrecipAverage    *rawAreaItem       `json:"precipAverage"`
}

type rawAreaTimeItem struct {
	TimeDefines []time.Time `json:"timeDefines"`
	Areas       rawAreaItem `json:"areas"`
}

func Parse010000(reader io.Reader) ([]*Forecast, error) {
	var rawOfficeForecasts []*rawOfficeForecast

	if err := json.NewDecoder(reader).Decode(&rawOfficeForecasts); err != nil {
		return nil, fmt.Errorf("decode error: %w", err)
	}

	forecastMap := map[string]*Forecast{}

	for _, rawOfficeForecast := range rawOfficeForecasts {
		rawRangeForecasts := []*rawRangeForecast{
			&rawOfficeForecast.Srf,
			&rawOfficeForecast.Week,
		}

		for _, rawRangeForecast := range rawRangeForecasts {
			for _, rawAreaTimeItem := range rawRangeForecast.TimeSeries {
				rawAreaItem := &rawAreaTimeItem.Areas
				rawAreaItem.Area.Code = rawOfficeForecast.OfficeCode
				rawAreaItem.Area.Name = rawOfficeForecast.Name

				for i, comesAt := range rawAreaTimeItem.TimeDefines {
					key := fmt.Sprintf("%s-%s", rawAreaItem.Area.Code, comesAt.Format("2006010215"))

					forecast := forecastMap[key]
					if forecast == nil {
						forecast = &Forecast{Area: rawAreaItem.Area, ComesAt: comesAt}
						forecastMap[key] = forecast
					}

					forecast.ReportedAt = rawRangeForecast.ReportDateTime

					updateForecast(forecast, rawAreaItem, i)
				}
			}
		}
	}

	forecasts := make([]*Forecast, len(forecastMap))
	i := 0

	for _, forecast := range forecastMap {
		forecasts[i] = forecast
		i++
	}

	return forecasts, nil
}
