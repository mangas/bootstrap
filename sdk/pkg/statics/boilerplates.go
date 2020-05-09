package statics

/*
BPAsset satisfies AssetEmbedder interface
This one is namespaced for root of the boilerplate directory
see boilerplate directory in bootstrap to see the content.
*/

import (
	"errors"
	"fmt"
	"github.com/getcouragenow/bootstrap/sdk/pkg/common/embed"
	_ "github.com/getcouragenow/bootstrap/statiks/bpcore"
	_ "github.com/getcouragenow/bootstrap/statiks/bplyft"
	_ "github.com/getcouragenow/bootstrap/statiks/bproot"
	_ "github.com/getcouragenow/bootstrap/statiks/bptool"
	"github.com/rakyll/statik/fs"
	"net/http"
)

var (
	namespaces = []string{"core", "lyft", "tool", "root"}
)

type BPAsset struct {
	fsys http.FileSystem // the rakyll fs
}

func filterNS(arg string) bool {
	for _, ns := range namespaces {
		if ns == arg {
			return true
		}
	}
	return false
}

// NewBPAsset function to filter valid namespace
// for now this will be hardcoded, later down the line,
// it will be generated.
func NewBPAsset(namespaceArg string) (embed.AssetEmbedder, error) {
	found := filterNS(namespaceArg)
	if !found {
		return nil, errors.New(
			fmt.Sprintf("namespace not found: %s", namespaceArg),
		)
	}
	namespace := fmt.Sprintf("bp%s", namespaceArg)
	return newBPAsset(namespace)
}

// NewBPAsset will return BPAsset
func newBPAsset(namespace string) (embed.AssetEmbedder, error) {
	bfs, err := fs.NewWithNamespace(namespace)
	if err != nil {
		return nil, err
	}
	return &BPAsset{
		fsys: bfs,
	}, nil
}

func (r *BPAsset) GetFS() http.FileSystem { return r.fsys }
func (r *BPAsset) WriteAllFiles(outputPath string) error {
	return writeAllFiles(r.fsys, outputPath)
}
func (r *BPAsset) ReadSingleFile(name string) ([]byte, error) {
	return readSingleFile(r.fsys, name)
}
