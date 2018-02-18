FROM golang:1.9.2-alpine

WORKDIR /go/src/github.com/antonybudianto/resto-api

RUN apk --update add git
RUN go-wrapper download -u github.com/golang/dep/cmd/dep \
    && go-wrapper install github.com/golang/dep/cmd/dep

COPY . .

ADD files/bin/wait-for.sh /wait-for.sh
RUN chmod +x /wait-for.sh

RUN dep ensure

ENTRYPOINT [ "/wait-for.sh", "db:3306", "--" ]

CMD [ "sh", "-c", "go run cmd/apiapp/*.go" ]
