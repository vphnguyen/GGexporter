FROM golang:1.20rc1-alpine3.17 as build
MAINTAINER nguyenvp2 

WORKDIR /GGexporter
COPY . .

RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o runfile

FROM gcr.io/distroless/base-debian10
WORKDIR /COPY --from=build /GGexporter/runfile /GGexporter/runfile


EXPOSE 9101
USER nonroot:nonroot


CMD ["/GGexporter/runfile -mH 10.0.0.202  -rU http://localhost:1616  -debug"]