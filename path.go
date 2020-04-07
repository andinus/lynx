package lynx

import "golang.org/x/sys/unix"

// UnveilPath takes a path, permission & unveils it, returning an
// error if unveil fails. "no such file or directory" error is
// ignored.
func UnveilPath(path string, flags string) (err error) {
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

// UnveilPathStrict is just a wrapper around unix.Unveil.
func UnveilPathStrict(path string, flags string) error {
	return unix.Unveil(path, flags)
}
