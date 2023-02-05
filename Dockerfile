FROM golang:1.19-bullseye
WORKDIR /BlazeGPT/
COPY . .
RUN mkdir logs
RUN go get github.com/spf13/viper
RUN go get github.com/sirupsen/logrus
RUN go build BrazinoGPT
ENTRYPOINT ["./BrazinoGPT"]
