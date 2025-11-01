FROM golang:alpine AS build
WORKDIR /src
COPY . /src/
RUN GOOS=linux GOARCH=amd64 go build -o /bin/backblaze-personal-exporter

FROM scratch
COPY --from=build /bin/backblaze-personal-exporter /bin/backblaze-personal-exporter
CMD ["/bin/backblaze-personal-exporter", "--backblazeData", "/data" ]

