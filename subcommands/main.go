package main

import (
    "tharja/handler"
    _ "tharja/mod/hash"
	_ "tharja/mod/webserve"
)

func main() {
    handler.ExecuteCommand()
}

