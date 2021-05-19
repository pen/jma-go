package main

import (
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/pen/jma-go/client"
	"github.com/pen/jma-go/overview"
	"github.com/pen/jma-go/overview/week"
)

func main() {
	ctx := context.Background()
	c := client.New()

	pathCode := "270000"

	overview, err := c.GetOverview(ctx, pathCode)
	if err != nil {
		fmt.Printf("%s\n", err.Error())
		os.Exit(1)
	}

	weekly, err := c.GetWeeklyOverview(ctx, pathCode)
	if err != nil {
		fmt.Printf("%s\n", err.Error())
		os.Exit(1)
	}

	dumpOverviewAndWeek(overview, weekly)
}

func dumpOverviewAndWeek(o *overview.Overview, wo *week.WeeklyOverview) {
	fmt.Printf("%s %sの天気です。\n",
		o.ReportedAt.Format("2006年01月02日 15時04分"),
		o.AreaName,
	)

	if o.Headline != nil {
		fmt.Printf("[%s]\n", *o.Headline)
	}

	text := strings.ReplaceAll(o.Text, "\n\n", "\n")
	fmt.Printf("%s\n\n", text)

	text = strings.ReplaceAll(wo.Text, "\n\n", "\n")
	fmt.Printf("%sです。\n%s\n\n", wo.HeadTitle, text)

	fmt.Printf("(%s)\n", o.OfficeName)
}
