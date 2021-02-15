package keg

import (
	"time"
)

// TStamp outputs	an RFC3338 (ISO8601) timestamp suitable for
// importing into any modern database and as used by JSON extensions and
// YAML structured data files.
func TStamp() string {
	return time.Now().Format(time.RFC3339)
}
