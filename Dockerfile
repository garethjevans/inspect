FROM --platform=${BUILDPLATFORM} alpine:3.14.1

ARG TARGETOS
ARG TARGETARCH
ARG TARGETPLATFORM

LABEL maintainer="Gareth Evans <gareth@bryncynfelin.co.uk>"
COPY dist/inspect-${TARGETOS}_${TARGETOS}_${TARGETARCH}/inspect /usr/bin/inspect

ENTRYPOINT [ "/usr/bin/inspect" ]

CMD ["--help"]
