package webserve

import (
    "fmt"
	"os"
	"net/http"
    "time"
    "tharja/handler"
)

func init() {
    handler.Register("webserve","host folder as static webserver",hash)
}

func hash(input []string) {
    dir := ""
    if len(input) == 1 {
        dir = input[0]
    }else{
        cwd, err := os.Getwd()
        if err != nil { panic(err) }
        dir = cwd
    }
    dataRoot, err := os.OpenRoot(dir)
	if err != nil {
		panic(err)
	}
	fmt.Println("Hosting on :1933 with folder:",dir)
	mux := http.NewServeMux()
	mux.Handle("/", http.FileServer(http.FS(dataRoot.FS())))
	fmt.Println(http.ListenAndServe(":1933", AddLoggingToMux(mux)))
    
}

func AddLoggingToMux(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		defer func() {
			re := recover()
            if re == nil {
                re = ""
            }
			fmt.Printf("[%s] %s %s %s (%s) %s\n",time.Now().Format(time.DateTime),r.RemoteAddr,r.Method,r.URL,time.Since(start),re)
		}()
		next.ServeHTTP(w, r)
	})
}

