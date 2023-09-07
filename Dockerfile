FROM golang:1.21-alpine as source
RUN apk add make
WORKDIR /cui
COPY . /cui
RUN make vendor

FROM source as build
RUN apk add fortify-headers gcc libc-dev
RUN make

FROM build as test
RUN make test

FROM alpine

LABEL org.opencontainers.image.title=cui
LABEL org.opencontainers.image.version=v0.5.0
LABEL org.opencontainers.image.description="http request/response tui"
LABEL org.opencontainers.image.url=https://github.com/mfinelli/cui
LABEL org.opencontainers.image.source=https://github.com/mfinelli/cui
LABEL org.opencontainers.image.licenses=GPL-3.0-or-later

RUN addgroup -S cui && adduser -S cui -G cui
COPY --from=source /cui /usr/src/cui
COPY --from=build /cui/cui /usr/bin/cui
USER cui
CMD ["cui"]
