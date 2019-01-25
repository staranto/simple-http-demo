ARG DFROM
FROM ${DFROM}

ARG ARCH
COPY build/pounder.${ARCH} /tmp/pounder

HEALTHCHECK --interval=2s --timeout=1s --start-period=10s --retries=1 \
    CMD wget --quiet --tries=1 --spider http://localhost:8282/live || exit 1

CMD ["/tmp/pounder", "http://counter:8181", "blue"]
