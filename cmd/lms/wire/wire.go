//+build wireinject

package wire

import (
	"context"
	"github.com/google/wire"
)

func InitializeApp(ctx context.Context, configPath string) (Application, error) {
	wire.Build(
		NewApplication,
		pkgSet,
		userSet,
		serverSet,
		authSet,
	)
	return Application{}, nil
}
