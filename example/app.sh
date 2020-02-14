#!/bin/bash
set -e
i=0
while true; do
    ((i=i+1))
    if (( "$i" >= 5 )); then
      sdnotify ready
      exit
    else
      sdnotify watchdog
      sleep 1
    fi
done
