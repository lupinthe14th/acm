FROM golang AS build-env
ENV CGO_ENABLED 0

ADD . /pub_src

WORKDIR /pub_src/
RUN go build -gcflags "all=-N -l" -o /pub ./publisher

FROM scratch

COPY --from=build-env /pub /
CMD ["/pub"]
