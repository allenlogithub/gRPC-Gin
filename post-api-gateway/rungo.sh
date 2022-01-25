gofmt -s -w .

d="/usr/local/go/src/post-api-gateway"
mkdir -p "$d"
cp -r server "$d"
cp -r client "$d"
cp -r config "$d"
cp -r client "$d"
cp -r proto "$d"

go run main.go
