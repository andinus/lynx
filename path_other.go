// +build !openbsd

package lynx

// UnveilPath takes a path, permission & unveils it, it does nothing
// on non OpenBSD systems.
func UnveilPath(_ string, _ string) error {
	return nil
}

// UnveilPathStrict is just a wrapper around unix.Unveil. It does
// nothing on non OpenBSD systems.
func UnveilPathStrict(_ string, _ string) error {
	return nil
}
