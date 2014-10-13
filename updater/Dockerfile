FROM scratch
WORKDIR /home/weave
ADD ./weave-updater /home/weave/
ENTRYPOINT ["/home/weave/weave-updater", "-api", "unix:///var/run/docker.sock", "-endpoint", "http://weavedns.weave:6785"]
