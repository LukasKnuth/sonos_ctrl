# sonos_ctrl

A simple Sonos Controller CLI, built as an experiment in golang.

## Features

* Automatic discovery of Sonos speakers/groups on the local network
* Caching of discovered speakers/groups for fast control
* CLI actions for simple commands such as "pause"/"play"/"next" etc

## Ideas

Since the program is being built with rxGo, it could also be running in the background indefinately.

Then, it could re-query the state of all Sonos speakers/groups in the background and expose them via a common platform API, such as DBus for Linux. This could then be used by a compatible consumer to display player controls and allow controlling Sonos speakers from 3rd-party applications.

## Sources

* SoCo Project: https://github.com/SoCo/SoCo/