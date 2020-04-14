// +build !openbsd

package lynx

// UnveilCommands takes a slice of commands & unveils them one by one,
// it does nothing on non OpenBSD systems.
func UnveilCommands(_ []string) error {
	return nil
}
