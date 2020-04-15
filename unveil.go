// +build openbsd

package lynx

import "golang.org/x/sys/unix"

// Unveil takes a path, permission & unveils it, returning an
// error if unveil fails. "no such file or directory" error is
// ignored.
func Unveil(path string, flags string) (err error) {
	err = unix.Unveil(path, flags)

	// "no such file or directory" error is ignored.
	if err != nil && err.Error() != "no such file or directory" {
		// Better error message could be returned like
		// one that includes the path on which unveil
		// failed.
		return err
	}
	// Returning nil because err can be "no such file or
	// directory" which needs to be ignored.
	return nil
}

// UnveilPath is kept for backwards compatibility, use Unveil instead.
func UnveilPath(path string, flags string) (err error) {
	return Unveil(path, flags)
}

// UnveilStrict is just a wrapper around unix.Unveil.
func UnveilStrict(path string, flags string) error {
	return unix.Unveil(path, flags)
}

// UnveilPathStrict is kept for backwards compatibility, use
// UnveilStrict instead.
func UnveilPathStrict(path string, flags string) error {
	return UnveilStrict(path, flags)
}
