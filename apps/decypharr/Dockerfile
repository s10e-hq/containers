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
        icu-libs \
        jq \
        libintl \
        nano \
        sqlite-libs \
        tzdata \
    && mkdir -p /app/bin \
    && curl -fsSL -o /tmp/app.tgz "https://github.com/sirrobot01/decypharr/releases/download/${VERSION}/decypharr_Linux_${TARGETARCH/amd64/x86_64}.tar.gz" \
    && tar xzf /tmp/app.tgz -C /app/bin --strip-components=1 \
    && chown -R root:root /app && chmod -R 755 /app

COPY . /

USER nobody:nobody
WORKDIR /config
VOLUME ["/config"]

ENTRYPOINT [ "/usr/bin/catatonit", "--", "/entrypoint.sh" ]
