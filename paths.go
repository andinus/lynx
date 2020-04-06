// Package lynx is a simple wrapper to unveil.
package lynx

import "golang.org/x/sys/unix"

// UnveilPaths takes a map of path, permission & unveils them one by
// one, it will return an error if unveil fails at any step. "no such
// file or directory" error is ignored.
func UnveilPaths(paths map[string]string) error {
	for k, v := range paths {
		err := unix.Unveil(k, v)

		// "no such file or directory" error is ignored.
		if err != nil && err.Error() != "no such file or directory" {
			// Better error message could be returned like
			// one that includes the path on which unveil
			// failed.
			return err
		}
	}
	// Returning nil because err can be "no such file or
	// directory" which needs to be ignored.
	return nil
}

// UnveilPathsStrict takes a map of path, permission & unveils them
// one by one, it will return an error if unveil fails at any steop.
// No error is ignored.
func UnveilPathsStrict(paths map[string]string) (err error) {
	for k, v := range paths {
		err = unix.Unveil(k, v)

		if err != nil {
			return
		}
	}
	return
}
