
FROM golang:1.19-bullseye as builder

ADD . /go/dachsurl
WORKDIR /go/dachsurl
RUN make clean && make && adduser --disabled-login --disabled-password nonroot

FROM scratch

COPY --from=builder /go/dachsurl/dachsurl /usr/bin/dachsurl
COPY --from=builder /etc/passwd /etc/passwd
USER nonroot

ENTRYPOINT [ "/usr/bin/dachsurl" ]
