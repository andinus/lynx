// +build openbsd

package lynx

import "golang.org/x/sys/unix"

// UnveilBlock is just a wrapper around unix.UnveilBlock, it does
// nothing extra. You should use unix.UnveilBlock.
func UnveilBlock() error {
	return unix.UnveilBlock()
}
