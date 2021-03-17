package configor

import (
	"fmt"
	"github.com/jinzhu/configor"
	"github.com/mix-go/cli/argv"
	"github.com/mix-go/cli-skeleton/globals"
)

func init()  {
	// Conf support YAML, JSON, TOML, Shell Environment
	if err := configor.Load(&globals.Config, fmt.Sprintf("%s/../conf/config.json", argv.Program().Dir)); err != nil {
		panic(err)
	}
}
