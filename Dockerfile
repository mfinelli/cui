FROM golang:1.19-alpine as source
RUN apk add make
WORKDIR /cui
COPY . /cui
RUN go mod vendor

FROM source as build
RUN make

FROM alpine
LABEL org.opencontainers.image.source https://github.com/mfinelli/cui
COPY --from=source /cui /usr/src/cui
COPY --from=build /cui/cui /usr/bin/cui
CMD ["cui"]
