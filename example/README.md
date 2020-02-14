This example can be tested by running 

```
$ build.sh
```

Then copy the example unit into ~/.config/systemd/user/notify_tester.service and edit the path to the dockersd wrapper.

and try

```
systemctl --user daemon-reload
systemctl --user start notify_test
systemctl --user status notify_test
```

This should block on start till the docker conatiner times out.
