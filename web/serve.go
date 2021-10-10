package web

import (
	"embed"
	"log"
	"minetest_client/frontend"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func Serve() {
	r := mux.NewRouter()
	http.Handle("/", r)

	// webdev flag
	useLocalfs := os.Getenv("WEBDEV") == "true"
	// static files
	r.PathPrefix("/").Handler(http.FileServer(getFileSystem(useLocalfs, frontend.Webapp)))

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}

func getFileSystem(useLocalfs bool, content embed.FS) http.FileSystem {
	if useLocalfs {
		log.Print("using live mode")
		return http.FS(os.DirFS("frontend"))
	}

	log.Print("using embed mode")
	return http.FS(content)
}
