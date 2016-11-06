FROM golang:1.7-alpine

COPY . /go/src/app

WORKDIR /go/src/app

ENV PORT 80

RUN apk --no-cache add curl git && \
  go get ./ && \
  go build && go install -v

expose 80

CMD ["app"]
