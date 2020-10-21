// +build !linux,!freebsd,!netbsd,!openbsd,!windows,!darwin,!js

package beeep

// SetAppID sets the name of your application
func SetAppID(appID string) {}

// Notify sends desktop notification.
func Notify(title, message string) error {
	return ErrUnsupported
}
