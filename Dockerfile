FROM alpine:3.17 AS Build

COPY --chmod=755 . /build
WORKDIR /build
RUN apk add --no-cache bash curl gcc git go musl-dev
RUN bash build.sh release docker

FROM alpine:3.17

ENV S6_SERVICES_GRACETIME=30000 \
    S6_KILL_GRACETIME=60000 \
    S6_CMD_WAIT_FOR_SERVICES_MAXTIME=0 \
    S6_SYNC_DISKS=1 \
    LANG=C.UTF-8 \
    PS1="\[\e[32m\][\[\e[m\]\[\e[36m\]\u \[\e[m\]\[\e[37m\]@ \[\e[m\]\[\e[34m\]\h\[\e[m\]\[\e[32m\]]\[\e[m\] \[\e[37;35m\]in\[\e[m\] \[\e[33m\]\w\[\e[m\] \[\e[32m\][\[\e[m\]\[\e[37m\]\d\[\e[m\] \[\e[m\]\[\e[37m\]\t\[\e[m\]\[\e[32m\]]\[\e[m\] \n\[\e[1;31m\]$ \[\e[0m\]" \
    TZ=Asia/Shanghai \
    PUID=911 \
    PGID=911 \
    GIN_MODE=release

RUN apk add --no-cache \
        tzdata \
        bash \
        s6-overlay && \
    rm -rf \
	    /tmp/* \
        /var/cache/apk/* \
	    /var/tmp/*

COPY --chmod=755 ./docker/rootfs /
COPY --chmod=755 --from=Build /build/bin/onelist /app/onelist

ENTRYPOINT [ "/init" ]

EXPOSE 5245
VOLUME [ "/config" ]