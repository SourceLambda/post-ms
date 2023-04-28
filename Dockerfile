FROM golang:1.20-alpine

WORKDIR /post_ms

COPY go.mod ./
COPY go.sum ./

RUN go mod download
COPY . ./

EXPOSE 8090

RUN go build -o /sourcelambda_post_ms
CMD [ "/sourcelambda_post_ms" ]