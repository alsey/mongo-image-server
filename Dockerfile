FROM golang
MAINTAINER Alsey.DAI <zamber@gmail.com>

RUN go get github.com/gorilla/mux
RUN go get github.com/nfnt/resize
RUN go get gopkg.in/mgo.v2

ADD . /go/src/mongo-image-server

RUN go install mongo-image-server

ENTRYPOINT [ "/go/bin/mongo-image-server" ]

EXPOSE 3000
