FROM ubuntu:22.04

ENV TZ=Asia/Shanghai \
    PUID=911 \
    PGID=911 \
    GIN_MODE=release

ARG OVERLAY_VERSION="v2.2.0.3"
ARG ONELISTP_VERSION="v1.0.1"
ARG DEBIAN_FRONTEND=noninteractive

RUN apt-get update -y && \
    apt-get install -y \
        wget \
        tar \
        xz-utils \
        tzdata && \
    mkdir -p \
        /config \
        /app && \
    if [ "$(uname -m)" = "x86_64" ]; then ARCH=amd64; elif [ "$(uname -m)" = "aarch64" ]; then ARCH=aarch64; elif [ "$(uname -m)" = "armv7l" ]; then ARCH=arm; fi && \
    wget \
        --no-check-certificate \
        -O /app/onelist \
        https://github.com/msterzhang/onelist/releases/download/${ONELISTP_VERSION}/onelist-linux-${ARCH} && \
    if [ "$(uname -m)" = "x86_64" ]; then ARCH=amd64; elif [ "$(uname -m)" = "aarch64" ]; then ARCH=aarch64; elif [ "$(uname -m)" = "armv7l" ]; then ARCH=arm; fi && \
    wget \
        --no-check-certificate \
        -O /tmp/s6-overlay-installer \
        https://github.com/just-containers/s6-overlay/releases/download/${OVERLAY_VERSION}/s6-overlay-${ARCH}-installer && \
    chmod 755 /tmp/s6-overlay-installer && \
    /tmp/s6-overlay-installer / && \
    chmod 755 /app/onelist && \
    apt-get autoremove && \
    apt-get clean && \
    rm -rf \
	    /tmp/* \
	    /var/lib/apt/lists/* \
	    /var/tmp/*

COPY --chmod=755 ./rootfs /

ENTRYPOINT [ "/init" ]

EXPOSE 5245
VOLUME [ "/config" ]