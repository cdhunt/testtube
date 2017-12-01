FROM golang:1.9.2-windowsservercore-1709
WORKDIR /go/src/hereiam/
COPY webserver.go .
RUN go build -o webserver.exe .

FROM microsoft/nanoserver:1709

COPY --from=0 /go/src/hereiam/webserver.exe .
CMD ["webserver.exe"]