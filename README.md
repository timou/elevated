# Elevated

A cross-platform check for elevated privileges.

On anything other than Windows, it checks whether os.Geteuid() == 0.

On Windows, it checks whether the current process token is elevated.

It has no dependencies.
