FROM golang:1.20.1-alpine as builder

WORKDIR /home/app

COPY . /home/app

RUN go mod download
RUN go build -o /worker ./cmd/worker


FROM scratch

COPY --from=builder /worker /worker

CMD [ "/worker" ]