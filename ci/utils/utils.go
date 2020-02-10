package utils

import (
	"time"

	"github.com/magefile/mage/sh"
)

// FlagEnv flags used in build
func FlagEnv(packageName string) map[string]string {
	hash, _ := sh.Output("git", "rev-parse", "--short", "HEAD")
	return map[string]string{
		"PACKAGE":     packageName,
		"COMMIT_HASH": hash,
		"BUILD_DATE":  time.Now().Format("2006-01-02T15:04:05Z0700"),
	}
}
