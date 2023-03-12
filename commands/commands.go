package commands

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/atotto/clipboard"
	"github.com/google/uuid"
	"github.com/urfave/cli/v2"
)

var UuidCommand = &cli.Command{
	Name:            "uuid",
	Description:     "A tool to generate uuid and copy it your clipboard automaticlly",
	HideHelpCommand: true,
	UsageText:       "uuid [count] -- count is set to 1 if not passed",
	Usage:           "uuid [count]",
	Action: func(ctx *cli.Context) error {
		var i int
		i, err := strconv.Atoi(ctx.Args().Get(0))
		if err != nil {
			i = 1
		}
		count := i
		uuids := []string{}
		for i := 0; i < count; i++ {
			uuids = append(uuids, uuid.New().String())
		}
		uuidsString := strings.Join(uuids, "\n")
		clipboard.WriteAll(uuidsString)
		fmt.Println(uuidsString)
		return nil
	},
}
