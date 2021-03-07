rm p2ppeer
go build -o p2ppeer peer/*.go
echo "running \"./p2ppeer -p 7777\""
./p2ppeer -p 7777