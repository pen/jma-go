package main

import (
	"context"
	"fmt"
	"os"

	"github.com/pen/jma-go/area"
	"github.com/pen/jma-go/client"
)

func main() {
	ctx := context.Background()
	c := client.New()

	areas, err := c.GetAreas(ctx)
	if err != nil {
		fmt.Printf("%s\n", err.Error())
		os.Exit(1)
	}

	for _, area := range areas {
		dumpArea(area)
	}
}

func dumpArea(a *area.Area) {
	fmt.Printf("%s/%s[%s] %s", a.Class, a.Code, a.ParentCode, a.Name)

	if a.NameKana != "" {
		fmt.Printf("(%s)", a.NameKana)
	}

	fmt.Printf("/%s\n", a.NameEn)
}
