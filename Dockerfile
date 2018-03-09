FROM odise/busybox-curl
ADD micro /micro
WORKDIR /
ENTRYPOINT [ "/micro" ]