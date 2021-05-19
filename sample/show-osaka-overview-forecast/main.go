package main

import (
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/pen/jma-go/client"
	"github.com/pen/jma-go/overview"
)

func main() {
	ctx := context.Background()
	c := client.New()

	overview, err := c.GetOverview(ctx, "270000")
	if err != nil {
		fmt.Printf("%s\n", err.Error())
		os.Exit(1)
	}

	dumpOverview(overview)
}

func dumpOverview(f *overview.Overview) {
	fmt.Printf("%s %sの天気です。\n",
		f.ReportedAt.Format("2006年01月02日 15時04分"),
		f.AreaName,
	)

	if f.Headline != nil {
		fmt.Printf("[%s]\n", *f.Headline)
	}

	text := strings.ReplaceAll(f.Text, "\n\n", "\n")

	fmt.Printf("%s\n(%s)\n", text, f.OfficeName)
}
