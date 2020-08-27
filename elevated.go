// +build !windows

package elevated

import "os"

func IsElevated() bool {
	return os.Geteuid() == 0
}