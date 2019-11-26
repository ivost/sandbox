FROM golang:1.13 as foundation

WORKDIR /build
COPY go.mod go.sum ./

RUN go mod download

FROM foundation as builder

COPY . .

RUN make build

RUN ls -alh /build/hello

FROM gcr.io/distroless/base as runtime

COPY --from=builder /build/hello /bin/app

ENTRYPOINT ["/bin/app"]