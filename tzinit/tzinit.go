package tzinit

import (
	"os"
)

func init() {
	os.Setenv("TZ", "Africa/Cairo")
}
