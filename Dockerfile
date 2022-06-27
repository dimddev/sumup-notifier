##
## Build
##
FROM golang:1.17-buster AS build

WORKDIR /app

COPY .env.local go.mod go.sum ./
COPY app ./app
RUN go get -d ./...
RUN go build -o sumup-notifier ./app/cmd/server/main.go

##
## Deploy
##
FROM debian:buster-slim
WORKDIR /
COPY --from=build /app/sumup-notifier /sumup-notifier
COPY --from=build /app/.env.local /.env.local
EXPOSE 8080
USER nobody:nogroup
ENTRYPOINT ["./sumup-notifier"]
