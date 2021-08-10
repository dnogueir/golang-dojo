FROM golang:1.16

WORKDIR /go/src
ENV PATH="/go/bin:${PATH}"

CMD ["tail", "-f", "/dev/null"]