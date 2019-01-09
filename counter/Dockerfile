ARG DFROM
FROM ${DFROM}

ARG ARCH
COPY build/counter.${ARCH} /tmp/counter

HEALTHCHECK --interval=3s --timeout=1s --start-period=15s --retries=1 \
    CMD wget --quiet --tries=1 --spider http://localhost:8282/live || exit 1

CMD ["/tmp/counter"]
