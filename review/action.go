package review

import (
	"fmt"
	"os"

	"github.com/itsubaki/appstore/client"
	"github.com/itsubaki/appstore/model"
	"github.com/itsubaki/appstore/util"
	"gopkg.in/urfave/cli.v1"
)

func Action(c *cli.Context) {
	if len(c.Args()) < 1 {
		fmt.Println("See: apst review -h")
		os.Exit(1)
	}

	country := c.String("country")
	b := client.Ranking(
		c.String("limit"),
		c.String("genre"),
		c.String("feed"),
		country,
	)

	list := model.NewAppFeed(b).AppList
	for i := 0; i < len(c.Args()); i++ {
		list = list.Select(c.Args().Get(i))
	}

	for i, app := range list {
		b := client.Review(app.ID, country)
		f := model.NewReviewFeed(*app, b)

		switch c.String("output") {
		case "json":
			fmt.Println(f.Json(c.Bool("pretty")))
		default:
			fmt.Println(app)
			for _, r := range f.ReviewList {
				util.ColorPrintln(r.Rating, r.String())
			}

			if c.Bool("stats") {
				fmt.Println(f.Stats())
			}

			if i != (len(list) - 1) {
				fmt.Println("")
			}
		}
	}

}
