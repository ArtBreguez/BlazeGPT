FROM golang:1.19-bullseye
RUN apt-get update && apt-get install -y git
COPY id_rsa /root/.ssh/id_rsa
RUN chmod 600 /root/.ssh/id_rsa
RUN ssh-keyscan github.com >> /root/.ssh/known_hosts
RUN git clone git@github.com:ArtBreguez/BlazeGPT.git 
WORKDIR ./BlazeGPT
RUN go get github.com/spf13/viper
RUN go get github.com/sirupsen/logrus
RUN go build BrazinoGPT
ENTRYPOINT ["./BrazinoGPT"]