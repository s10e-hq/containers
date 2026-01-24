# syntax=docker/dockerfile:1

FROM docker.io/library/alpine:3.23
ARG TARGETARCH
ARG VENDOR
ARG VERSION

USER root
WORKDIR /app

RUN \
    apk add --no-cache \
        bash \
        ca-certificates \
        catatonit \
        chromaprint \
        coreutils \
        curl \
        ffmpeg \
        fuse3 \
        icu-libs \
        jq \
        libintl \
        vim \
        shadow \
        sqlite-libs \
        tzdata \
        rclone \
    && mkdir -p /app/bin \
    && curl -fsSL -o /tmp/app.tgz "https://github.com/sirrobot01/decypharr/releases/download/${VERSION}/decypharr_Linux_${TARGETARCH/amd64/x86_64}.tar.gz" \
    && tar xzf /tmp/app.tgz -C /app/bin \
    && chown -R root:root /app \
    && chmod -R 755 /app \
    && chmod +x /app/bin/decypharr

COPY . /

USER nobody:nobody
WORKDIR /config
EXPOSE 8282/tcp
VOLUME ["/config"]

ENTRYPOINT [ "/usr/bin/catatonit", "--", "/entrypoint.sh" ]
