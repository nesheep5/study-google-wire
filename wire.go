// +build wireinject
// The build tag makes sure the stub is not built in the final build.

package main

import (
	"context"

	"wire-sample/foobarbaz"

	"github.com/google/wire"
)

func initializeBaz(ctx context.Context) (foobarbaz.Baz, error) {
	wire.Build(foobarbaz.SuperSet)
	return foobarbaz.Baz{}, nil
}
