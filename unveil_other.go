// +build !openbsd

package lynx

// Unveil takes a path, permission & unveils it, it does nothing
// on non OpenBSD systems.
func Unveil(_ string, _ string) error {
	return nil
}

// UnveilPath takes a path, permission & unveils it, it does nothing
// on non OpenBSD systems. Unveil should be used instead of
// UnveilPath.
func UnveilPath(_ string, _ string) error {
	return nil
}

// UnveilStrict is just a wrapper around unix.Unveil. It does
// nothing on non OpenBSD systems.
func UnveilStrict(_ string, _ string) error {
	return nil
}

// UnveilPathStrict is just a wrapper around unix.Unveil. It does
// nothing on non OpenBSD systems. UnveilStrict should be used instead
// of UnveilPathStrict.
func UnveilPathStrict(_ string, _ string) error {
	return nil
}
