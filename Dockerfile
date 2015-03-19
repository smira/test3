FROM busybox:ubuntu-14.04

ADD ./test3 /usr/bin/

EXPOSE 3000

ENTRYPOINT ["test3"]