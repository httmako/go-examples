package hash

import (
    "fmt"
    "flag"
	"os"
    "crypto/sha256"
    "tharja/handler"
)

func init() {
    handler.Register("sha256","Calculates sha256 of input",hash)
}

func hash(input []string) {
    fs := flag.NewFlagSet("foo", flag.ExitOnError)
    fileRead := fs.Bool("f", false, "calculate hash of file, if this is not set input is 1:1 hashed")
    fs.Parse(input)
	tail := fs.Args()
    
    if len(tail) == 0 {
		fs.PrintDefaults()
        handler.Exit("no file provided")
    }
	for _, name := range tail {
		h := sha256.New()
		if *fileRead {
			fBytes, err := os.ReadFile(name)
			if err != nil {
				handler.Exit("could not read file: "+name)
			}
			h.Write(fBytes)
		}else{
			h.Write([]byte(name))
		}
		fmt.Printf("%x  %s\n", h.Sum(nil), name)
	}
}
