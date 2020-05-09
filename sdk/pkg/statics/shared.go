package statics

import (
	"github.com/getcouragenow/bootstrap/sdk/pkg/common/osutil"
	"os"
)

const (
	ErrFileMapNotPopulated = "error filemap is not populated (empty)"
)

func checkAndMakeDir(path string) error {
	exists, err := osutil.Exists(path)
	if err != nil {
		return err
	}
	if !exists {
		if err := os.MkdirAll(path, 0755); err != nil {
			return err
		}
	}
	return nil
}

