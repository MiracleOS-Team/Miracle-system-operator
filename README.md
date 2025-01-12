<p align="center">MSO is a utility which allows you to perform maintenance tasks on your Miracle OS installation.</p>
</div>

<br/>

## Help

```md
The Miracle System Operator is a package manager, a system updater and a task automator.

Usage:
  mso [command]

Available Commands:
  android     Manage the mso waydroid subsystem
  completion  Generate the autocompletion script for the specified shell
  config      Manage the system configuration.
  export      Export an application or binary from the subsystem
  help        Help about any command
  install     Install an application inside the subsystem
  pico-init   Initialize the MSO subsystem, used for package management
  remove      Remove an application from the subsystem
  run         Run an application from the subsystem
  search      Search for an application to install inside the subsystem
  shell       Enter the subsystem environment
  sideload    Sideload DEB/APK packages inside the subsystem
  sys         Execute system commands, such as upgrading the system
  tasks       Create and manage tasks
  unexport    Unexport an application or binary from the subsystem
  update      Update the subsystem's package repository
  upgrade     Upgrade the packages inside the subsystem

Flags:
  -h, --help      Show help for mso.
  -v, --version   Show version for mso.

Use "mso [command] --help" for more information about a command.
```
##  MSO as system Shell

To use MSO as your system shell, you can copy the `usr/bin/mso-os-shell` script
to your system's `/usr/bin` directory and set it as your default shell. Your
image needs to implement the `usr/bin/os-shell` script, which will expand the
`$SHELL` environment variable, this is much needed for login shells and other
flags, this also ensures that the user's default shell is respected.

Our `mso-image` already implements this script.
