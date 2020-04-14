// +build !openbsd

package lynx

// UnveilBlock is just a wrapper around unix.UnveilBlock, it does
// nothing on non OpenBSD systems.
func UnveilBlock() error {
	return nil
}
