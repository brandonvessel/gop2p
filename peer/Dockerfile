From golang:alpine

WORKDIR /usr/src/app

COPY . .

RUN go build -o p2ppeer ./*.go

CMD [ "./p2ppeer", "-b", "10.10.10.2", "-p", "7777"]