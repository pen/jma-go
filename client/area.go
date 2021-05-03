package client

import (
	"context"
	"fmt"

	"github.com/pen/jma-go/area"
)

func (c *Client) GetAreas(ctx context.Context) ([]*area.Area, error) {
	resp, err := callAPI(ctx, c.httpClient, "bosai/common/const/area.json")
	if err != nil {
		return nil, fmt.Errorf("API error: %w", err)
	}
	defer resp.Body.Close()

	areas, err := area.Parse(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to parse areas: %w", err)
	}

	return areas, nil
}
