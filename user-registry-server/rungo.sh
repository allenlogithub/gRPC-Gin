gofmt -s -w .

d="/usr/local/go/src/user-registry-server"
mkdir -p "$d"
cp -r proto "$d"

go run main.go
