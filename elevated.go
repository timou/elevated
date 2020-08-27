// +build !windows

package elevated

func IsElevated() bool {
	return false
}