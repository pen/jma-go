package main

import (
	"context"
	"fmt"
	"os"

	"github.com/pen/jma-go/client"
	"github.com/pen/jma-go/forecast"
)

func main() {
	ctx := context.Background()
	c := client.New()

	forecasts, err := c.GetForecasts(ctx, "130000")
	if err != nil {
		fmt.Printf("%s\n", err.Error())
		os.Exit(1)
	}

	for _, forecast := range forecasts {
		dumpForecast(forecast)
	}
}

func dumpForecast(f *forecast.Forecast) {
	fmt.Printf("-- %s %s %s[%s]\n", f.ReportedAt.Format("2006.0102.1504"), f.ComesAt.Format("0102.1504"), f.Area.Name, f.Area.Code)

	if f.Weather != nil {
		fmt.Printf("天: %s", f.Weather.Code)
		if f.Weather.Text != nil {
			fmt.Printf(" [%s]", *f.Weather.Text)
		}
		fmt.Println("")
	}

	if f.Wind != nil {
		fmt.Printf("風: [%s]\n", f.Wind.Text)
	}

	if f.Wave != nil {
		fmt.Printf("波: [%s]\n", f.Wave.Text)
	}

	if f.Precipitation != nil {
		fmt.Printf("降: %d%%", f.Precipitation.Probability)

		if f.Precipitation.Reliability != nil {
			fmt.Printf("(%s)", *f.Precipitation.Reliability)
		}

		fmt.Println("")
	}

	if f.Temperature != nil {
		sp := ""
		if f.Temperature.Base != nil {
			fmt.Printf("温: %d", f.Temperature.Base)
			sp = " "
		}
		if f.Temperature.Min != nil {
			fmt.Print(sp)
			fmt.Printf("高:%d(%d〜%d)", f.Temperature.Max.Base, f.Temperature.Max.Lower, f.Temperature.Max.Upper)
			fmt.Printf(" 低:%d(%d〜%d)", f.Temperature.Min.Base, f.Temperature.Min.Lower, f.Temperature.Min.Upper)
		}
		fmt.Println("")
	}
}
