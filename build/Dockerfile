# builder image
FROM golang:1.16-alpine as builder
RUN mkdir /lms
ADD . /lms
WORKDIR /lms
RUN go build -o lms ./cmd/lms/main.go


# generate clean, final image for end users
FROM alpine:3.14
COPY --from=builder /lms/lms .
COPY --from=builder /lms/configs/config.yaml ./configs/config.yaml

## executable
ENTRYPOINT [ "./lms" ]
CMD [ "--config", "./configs/config.yaml" ]