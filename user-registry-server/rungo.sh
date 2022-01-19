gofmt -s -w .

d="/usr/local/go/src/user-registry-server"
mkdir -p "$d"
cp -r proto "$d"
cp -r config "$d"
cp -r server "$d"
cp -r controllers "$d"
cp -r databases "$d"
cp -r crypto "$d"

go run main.go
