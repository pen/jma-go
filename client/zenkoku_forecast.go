package client

import (
	"context"
	"fmt"

	"github.com/pen/jma-go/forecast"
)

func (c *Client) GetZenkokuForecasts(ctx context.Context) ([]*forecast.Forecast, error) {
	pathCode := "010000"

	resp, err := callAPI(ctx, c.httpClient, fmt.Sprintf("bosai/forecast/data/forecast/%s.json", pathCode))
	if err != nil {
		return nil, fmt.Errorf("API failed: pathCode: [%s] error: %w", pathCode, err)
	}
	defer resp.Body.Close()

	areas, err := forecast.Parse010000(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to parse forecasts: pathCode: [%s] error: %w", pathCode, err)
	}

	return areas, nil
}
