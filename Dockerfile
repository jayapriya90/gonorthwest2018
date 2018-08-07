FROM golang

ADD . $GOPATH/src/
WORKDIR $GOPATH/src/

ENV PROJECT bloomfilter
ENV FILE bloomfilter.go

RUN go get -u "github.com/tylertreat/BoomFilters"
RUN go get -u "github.com/dustin/go-probably"
RUN go get -u "github.com/axiomhq/hyperloglog"

CMD ["sh", "-c", "go run ./${PROJECT}/${FILE}"]