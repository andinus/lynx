// +build openbsd

package lynx

import "golang.org/x/sys/unix"

// Pledge is just a wrapper to unix.Pledge.
func Pledge(promises, execpromises string) error {
	return unix.Pledge(promises, execpromises)
}

// PledgePromises is just a wrapper to unix.PledgePromises.
func PledgePromises(promises string) error {
	return unix.PledgePromises(promises)
}

// PledgeExecpromises is just a wrapper to unix.PledgeExecpromises.
func PledgeExecpromises(execpromises string) error {
	return unix.PledgeExecpromises(execpromises)
}
