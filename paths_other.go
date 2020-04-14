// +build !openbsd

package lynx

// UnveilPaths takes a map of path, permission & unveils them one by
// one, it does nothing on non OpenBSD systems.
func UnveilPaths(_ map[string]string) error {
	return nil
}

// UnveilPathsStrict takes a map of path, permission & unveils them
// one by one, it does nothing on non OpenBSD systems.
func UnveilPathsStrict(_ map[string]string) error {
	return nil
}
