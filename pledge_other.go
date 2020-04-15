// +build !openbsd

package lynx

// Pledge is just a wrapper to unix.Pledge. It returns nil on
// unsupported systems.
func Pledge(_, _ string) error {
	return nil
}

// PledgePromises is just a wrapper to unix.PledgePromises. It returns
// nil on unsupported systems.
func PledgePromises(_ string) error {
	return nil
}

// PledgeExecpromises is just a wrapper to unix.PledgeExecpromises. It
// returns nil on unsupported systems.
func PledgeExecpromises(_ string) error {
	return nil
}
