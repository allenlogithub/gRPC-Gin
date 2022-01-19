gofmt -s -w .

d="/usr/local/go/src/user-auth-server"
mkdir -p "$d"
cp -r proto "$d"
cp -r config "$d"
cp -r databases "$d"
cp -r server "$d"

go run main.go
