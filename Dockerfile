FROM golang:1.10

EXPOSE 8081
ENV PORT 8081

COPY . /go/src/github.com/bigbluebutton/bbb-api-recordings
WORKDIR /go/src/github.com/bigbluebutton/bbb-api-recordings

RUN go get -u github.com/beego/bee \
    && go get -u github.com/kardianos/govendor \
    && govendor add +external

CMD ["bee", "run"]
