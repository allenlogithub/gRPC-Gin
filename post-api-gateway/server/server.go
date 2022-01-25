package server

// import "github.com/vsouza/go-gin-boilerplate/config"

func Init() {
	r := NewRouter()
	r.Run(":80")
}
