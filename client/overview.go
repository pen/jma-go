package client

import (
	"context"
	"fmt"

	"github.com/pen/jma-go/overview"
)

func (c *Client) GetOverview(ctx context.Context, pathCode string) (*overview.Overview, error) {
	resp, err := callAPI(ctx, c.httpClient, fmt.Sprintf("bosai/forecast/data/overview_forecast/%s.json", pathCode))
	if err != nil {
		return nil, fmt.Errorf("API failed: pathCode: [%s] error: %w", pathCode, err)
	}
	defer resp.Body.Close()

	overview, err := overview.Parse(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to parse overviewforecast: pathCode: [%s] error: %w", pathCode, err)
	}

	return overview, nil
}
