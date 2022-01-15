gofmt -s -w .

d="/usr/local/go/src/user-api-gateway"
mkdir -p "$d"
cp -r server "$d"
cp -r controllers "$d"
cp -r crypto "$d"

go run main.go
