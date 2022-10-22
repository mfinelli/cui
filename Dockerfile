FROM golang:1.19-alpine as build
RUN apk add make
WORKDIR /cui
COPY . /cui
RUN make

FROM alpine
COPY --from=build /cui/cui /usr/bin/cui
CMD ["cui"]
