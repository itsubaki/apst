package ranking

import (
	"fmt"

	"github.com/itsubaki/apst/client"
	"github.com/itsubaki/apst/util"
	cli "gopkg.in/urfave/cli.v1"
)

func Action(c *cli.Context) {
	b := client.Ranking(
		util.Limit(c.String("limit")),
		util.Genre(c.String("genre")),
		c.String("feed"),
		c.String("country"),
	)

	kw := util.Keyword(c.Args())
	list := NewFeed(b).Select(kw)
	for i := 1; i < len(c.Args()); i++ {
		list = list.Select(c.Args().Get(i))
	}

	for _, app := range list {
		fmt.Println(app)
	}
}
