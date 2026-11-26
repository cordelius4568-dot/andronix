package app

import (
	"os"

	"github.com/AndronixApp/andronix-distros/internal/sys"
)

// Host is where andronix runs: distro (inside one), standalone (Andronix's
// own app: ANDRONIX_HOST=standalone), termux, or linux. It is the "host"
// field of the --progress json hello event.
func Host() string {
	switch {
	case os.Getenv("ANDRONIX_DISTRO") != "" || sys.InGuest():
		return "distro"
	case os.Getenv("ANDRONIX_HOST") == "standalone":
		return "standalone"
	case sys.IsTermux():
		return "termux"
	}
	return "linux"
}
