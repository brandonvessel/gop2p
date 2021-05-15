From golang:alpine

WORKDIR /usr/src/app

COPY . .

RUN go build -o p2pclient ./*.go

CMD [ "./p2pclient", "-b", "10.10.10.2", "-p", "7777"]