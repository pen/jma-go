package client

import (
	"context"
	"fmt"

	"github.com/pen/jma-go/forecast"
)

func (c *Client) GetForecasts(ctx context.Context, pathCode string) ([]*forecast.Forecast, error) {
	resp, err := callAPI(ctx, c.httpClient, fmt.Sprintf("bosai/forecast/data/forecast/%s.json", pathCode))
	if err != nil {
		return nil, fmt.Errorf("API failed: pathCode: [%s] error: %w", pathCode, err)
	}
	defer resp.Body.Close()

	var areas []*forecast.Forecast

	if pathCode == "010000" {
		areas, err = forecast.Parse010000(resp.Body)
	} else {
		areas, err = forecast.Parse(resp.Body)
	}

	if err != nil {
		return nil, fmt.Errorf("failed to parse forecasts: pathCode: [%s] error: %w", pathCode, err)
	}

	return areas, nil
}

func (c *Client) GetZenkokuForecasts(ctx context.Context) ([]*forecast.Forecast, error) {
	return c.GetForecasts(ctx, "010000")
}
