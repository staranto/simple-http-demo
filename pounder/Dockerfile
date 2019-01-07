ARG DFROM
FROM ${DFROM}

ARG ARCH
COPY build/pounder.${ARCH} /tmp/pounder

CMD ["/tmp/pounder", "http://counter:8181", "blue"]
