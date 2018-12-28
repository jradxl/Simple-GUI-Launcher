Simple GUI Launcher

Executes the configured program after a delay, unless the button to stop is pressed.

I use Kodi on Debian Mate desktop, which logs in automatically and autostarts Kodi.
The lightdm auto login has delay of 20 second to allow a login to some other account.
Once the login to my Kodi account has been achieved the ~/.config/autostart configuration starts
this launcher, with the option to cancel it, in case I wish to maintain the Kodi account.

Built with Golang and https://github.com/andlabs/ui

Usage:
launcher "program full path"

Example:
launcher /usr/bin/Kodi

No releases, you'll need to build source.

December 2018


