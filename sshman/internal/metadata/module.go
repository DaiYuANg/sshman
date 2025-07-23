package metadata

import (
	"github.com/adrg/xdg"
	"go.uber.org/fx"
	"go.uber.org/zap"
	"path"
)

var Module = fx.Module("metadata", fx.Provide(newMetadataStore), fx.Invoke(lifecycle))

func newMetadataStore(logger *zap.SugaredLogger) (*Store, error) {
	storePath := path.Join(xdg.DataHome, "metadata")
	logger.Debugw("creating metadata store", "path", storePath)
	return NewStore(storePath)
}

func lifecycle(store *Store, lc fx.Lifecycle) {
	lc.Append(fx.StopHook(func() error {
		return store.Close()
	}))
}
