rm p2pclient
go build -o p2pclient client/*.go
echo "running \"./p2pclient -b 127.0.0.1 -p 7777\""
./p2pclient -b 127.0.0.1 -p 7777