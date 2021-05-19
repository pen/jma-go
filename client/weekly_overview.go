package client

import (
	"context"
	"fmt"

	"github.com/pen/jma-go/overview/week"
)

func (c *Client) GetWeeklyOverview(ctx context.Context, pathCode string) (*week.WeeklyOverview, error) {
	resp, err := callAPI(ctx, c.httpClient, fmt.Sprintf("bosai/forecast/data/overview_week/%s.json", pathCode))
	if err != nil {
		return nil, fmt.Errorf("API failed: pathCode: [%s] error: %w", pathCode, err)
	}
	defer resp.Body.Close()

	weeklyOverview, err := week.Parse(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to parse weekoverview: pathCode: [%s] error: %w", pathCode, err)
	}

	return weeklyOverview, nil
}
