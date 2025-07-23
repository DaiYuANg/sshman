package internal

import (
	"fmt"
	"go.uber.org/fx"
)

func component[T any](obj T, name string) interface{} {
	return fx.Annotate(
		obj,
		fx.ResultTags(fmt.Sprintf(`name:"%s"`, name)),
	)
}
