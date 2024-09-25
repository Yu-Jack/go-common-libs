package ns

import "golang.org/x/sys/unix"

// Join joins all the namespaces in the Joiners.
func (joiners *Joiners) Join() (err error) {
	for _, joiner := range *joiners {
		if joiner.isJoined {
			logrus.Tracef("Already joined namespace: %s", joiner.namespace)
			continue
		}

		if joiner.namespace == types.NamespaceMnt {
			err := unix.Unshare(unix.CLONE_NEWNS)
			if err != nil {
				return errors.Wrapf(err, "failed to unshare namespace: %+s", joiner.namespace)
			}
		}

		if err := unix.Setns(joiner.fd, 0); err != nil {
			return errors.Wrapf(err, "failed to set namespace: %+s", joiner.namespace)
		}

		joiner.isJoined = true
		logrus.Tracef("Joined namespace: %v", joiner.namespace)
	}
	return nil
}

func Gettid() int {
	return unix.Gettid()
}
