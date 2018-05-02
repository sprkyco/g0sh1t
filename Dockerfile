FROM golang:alpine

ADD . /go/src/sprkyco/g0sh1t
RUN go install sprkyco/g0sh1t
CMD ["/go/bin/g0sh1t"]
EXPOSE 8000