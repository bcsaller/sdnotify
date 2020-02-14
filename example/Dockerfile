FROM bcsaller/sdnotify:latest as sdnotify
FROM ubuntu:latest
COPY --from=sdnotify /sdnotify /bin/sdnotify
COPY app.sh /app.sh
WORKDIR /
CMD [ "/bin/bash", "app.sh" ]
