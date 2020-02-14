Simple wrapper to help include a statically linked binary into a docker container for use with the systemd watchdog/notify system.

The wrapper only supports three commands:

sdnotify watchdog
sdnotify ready
sdnotify <any strings>

The final form sets a status message for the unit.


Note that this still needs support from systemd


```
[Unit]
Description=showing docker with systemd-notify support

[Service]
Type=notify
NotifyAccess=all
WatchdogSec=10
ExecStart=/usr/local/bin/dockersd %N notify_tester:latest

[Install]
WantedBy=default.target
```

Where dockersd in this case might look like 

```
#!/bin/bash
# Internally the process must use some form of sd_notify(3) or
# systemd-notify(1) to signal the unit.
SD_NOTIFY_PROXY=/usr/libexec/sdnotify_proxy
SD_SOCK=/tmp/$1.sock

function sdproxy() {
    $SD_NOTIFY_PROXY $SD_SOCK /usr/bin/docker run --volume=$SD_SOCK:$SD_SOCK \
    --env=NOTIFY_SOCKET=$SD_SOCK --env=WATCHDOG_USEC=$WATCHDOG_USEC \
    $@
}

# Invoke the wrapper as 
# $ sdproxy %N dockerargs beyond "run"
sdproxy ${@:2}
```

All of which depends on sdnotify_proxy which can be found at https://github.com/coreos/sdnotify-proxy
