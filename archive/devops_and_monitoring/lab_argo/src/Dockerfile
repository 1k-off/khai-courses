FROM golang:1-alpine as build
WORKDIR /src
RUN apk update && apk add --no-cache ca-certificates git
COPY go.mod go.sum ./
RUN go mod download
COPY . .
WORKDIR /src
RUN CGO_ENABLED=0 GOOS=linux go build -o /app/lab_argo

FROM scratch
LABEL maintainer="Bogdan Kosarevskyi <bogdan.kosarevskyi@gmail.com>"
WORKDIR /app
COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=build /src/mock /app/mock
COPY --from=build /app /app
EXPOSE 3000
CMD ["/app/lab_argo"]