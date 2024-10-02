FROM golang:1.22-alpine AS build
WORKDIR /app
COPY . .
RUN go build -o cloud_storage ./cmd/

FROM alpine
WORKDIR /app
COPY --from=build /app/config.json .
COPY --from=build /app/config/config.yml /app/config/.
COPY --from=build /app/cloud_storage .
ENTRYPOINT [ "./cloud_storage" ]
