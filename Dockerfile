FROM docker.artifactory.kasikornbank.com:8443/golang:1.15.3 as builder

#PREPARED project
RUN mkdir /build
WORKDIR /build
COPY main.go .
COPY internal .
COPY go.mod .
COPY config .
RUN go mod download

#RUN go build -o /app main.go
ENV GOCACHE=/tmp/.cache
RUN CGO_ENABLED=0 GOOS=linux GOCACHE=/tmp/.cache go build -o /app

FROM docker.artifactory.kasikornbank.com:8443/alpine:3.10
ENV GOTRACEBACK=single XDG_CACHE_HOME=/tmp/.cache
COPY --from=builder /app .
CMD ["./app"]
