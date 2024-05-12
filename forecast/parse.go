package forecast

import (
	"encoding/json"
	"fmt"
	"io"
	"strconv"
	"time"
)

type rawForecast struct {
	PublishingOffice string         `json:"publishingOffice"`
	ReportDateTime   time.Time      `json:"reportDateTime"`
	TimeSeries       []*rawTimeItem `json:"timeSeries"`
	TempAverage      *rawAreaItem   `json:"tempAverage"`
	PrecipAverage    *rawAreaItem   `json:"precipAverage"`
}

type rawTimeItem struct {
	TimeDefines []time.Time    `json:"timeDefines"`
	Areas       []*rawAreaItem `json:"areas"`
}

type rawAreaItem struct {
	Area          Area     `json:"area"`
	WeatherCodes  []string `json:"weatherCodes"`
	Weathers      []string `json:"weathers"`
	Winds         []string `json:"winds"`
	Waves         []string `json:"waves"`
	Pops          []string `json:"pops"`
	Reliabilities []string `json:"reliabilities"`
	Temps         []string `json:"temps"`
	TempsMin      []string `json:"tempsMin"`
	TempsMinUpper []string `json:"tempsMinUpper"`
	TempsMinLower []string `json:"tempsMinLower"`
	TempsMax      []string `json:"tempsMax"`
	TempsMaxUpper []string `json:"tempsMaxUpper"`
	TempsMaxLower []string `json:"tempsMaxLower"`
	Min           string   `json:"min"`
	Max           string   `json:"max"`
}

func Parse(reader io.Reader) ([]*Forecast, error) {
	var rawForecasts []*rawForecast

	if err := json.NewDecoder(reader).Decode(&rawForecasts); err != nil {
		return nil, fmt.Errorf("decode error: %w", err)
	}

	forecastMap := map[string]*Forecast{}

	for _, rawForecast := range rawForecasts {
		for _, rawTimeItem := range rawForecast.TimeSeries {
			for _, rawAreaItem := range rawTimeItem.Areas {
				for i, comesAt := range rawTimeItem.TimeDefines {
					key := fmt.Sprintf("%s-%s", rawAreaItem.Area.Code, comesAt.Format("2006010215"))

					forecast := forecastMap[key]
					if forecast == nil {
						forecast = &Forecast{Area: rawAreaItem.Area, ComesAt: comesAt}
						forecastMap[key] = forecast
					}

					forecast.ReportedAt = rawForecast.ReportDateTime

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

//nolint:funlen
func updateForecast(forecast *Forecast, areaItem *rawAreaItem, index int) {
	if len(areaItem.WeatherCodes) > index {
		if forecast.Weather == nil {
			forecast.Weather = &Weather{}
		}

		forecast.Weather.Code = areaItem.WeatherCodes[index]

		if len(areaItem.Weathers) > index {
			forecast.Weather.Text = &areaItem.Weathers[index]
		}
	}

	if len(areaItem.Winds) > index && forecast.Wind == nil {
		forecast.Wind = &Wind{
			Text: areaItem.Winds[index],
		}
	}

	if len(areaItem.Waves) > index && forecast.Wave == nil {
		forecast.Wave = &Wave{
			Text: areaItem.Waves[index],
		}
	}

	if len(areaItem.Pops) > index {
		if forecast.Precipitation == nil {
			forecast.Precipitation = &Precipitation{}
		}

		probability, _ := strconv.Atoi(areaItem.Pops[index])
		forecast.Precipitation.Probability = probability

		if len(areaItem.Reliabilities) > index && areaItem.Reliabilities[index] != "" {
			forecast.Precipitation.Reliability = &areaItem.Reliabilities[index]
		}
	}

	if len(areaItem.Temps) > index {
		if forecast.Temperature == nil {
			forecast.Temperature = &Temperature{}
		}

		base, _ := strconv.Atoi(areaItem.Temps[index])
		forecast.Temperature.Base = &base
	}

	if len(areaItem.TempsMin) > index && areaItem.TempsMin[index] != "" {
		if forecast.Temperature == nil {
			forecast.Temperature = &Temperature{}
		}

		minBase, _ := strconv.Atoi(areaItem.TempsMin[index])
		minUpper, _ := strconv.Atoi(areaItem.TempsMinUpper[index])
		minLower, _ := strconv.Atoi(areaItem.TempsMinLower[index])
		forecast.Temperature.Min = &TemperatureRange{
			Base:  minBase,
			Upper: minUpper,
			Lower: minLower,
		}
	}

	if len(areaItem.TempsMax) > index && areaItem.TempsMax[index] != "" {
		if forecast.Temperature == nil {
			forecast.Temperature = &Temperature{}
		}

		maxBase, _ := strconv.Atoi(areaItem.TempsMax[index])
		maxUpper, _ := strconv.Atoi(areaItem.TempsMaxUpper[index])
		maxLower, _ := strconv.Atoi(areaItem.TempsMaxLower[index])
		forecast.Temperature.Max = &TemperatureRange{
			Base:  maxBase,
			Upper: maxUpper,
			Lower: maxLower,
		}
	}
}
