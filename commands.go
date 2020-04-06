package lynx

import (
	"fmt"
	"os"
	"strings"

	"golang.org/x/sys/unix"
)

// UnveilCommands takes a slice of commands & unveils them one by one,
// it will return an error if unveil fails at any step. "no such file
// or directory" error is ignored.
func UnveilCommands(commands []string) error {
	// Get $PATH & split it in a list.
	pathList := strings.Split(os.Getenv("PATH"), ":")

	// We work on unveiling each command one by one.
	for _, cmd := range commands {
		// Unveil this command on every PATH.
		for _, path := range pathList {
			err := unix.Unveil(fmt.Sprintf("%s/%s",
				path, cmd), "rx")

			// "no such file or directory" error is
			// ignored because binaries are not placed in
			// every PATH.
			if err != nil && err.Error() != "no such file or directory" {
				// Better error message could be
				// returned like one that includes the
				// path on which unveil failed.
				return err
			}
		}
	}
	// Returning nil because err can be "no such file or
	// directory" which needs to be ignored.
	return nil
}
