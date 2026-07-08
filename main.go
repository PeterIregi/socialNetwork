package main

//last time we made a web server with go that displayed a simple message lets do that again .

import (
	"net/http"
	//this is the package that allows us to create a web server in go
)

func main() {
	http.HandleFunc("/", home)
	//this function tells the server to listen for requests to the root URL and call the home function when a request is received.
	http.ListenAndServe(":8080", nil)
	//this function starts the server and listens for incoming requests on port 8080. the second argument is nil because we are not using any custom handlers.
}
func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello Peter!"))
	//this function writes the response to the client. it takes two arguments, the first is the response writer and the second is the request. we are writing a simple message "hello world" to the response writer.
}

//lets see if it rus as we expect it to.
//now if we go to our browe=ser and input localhost:8080 we should see the message "hello world" displayed in the browser.if wr change it to hello peter it will change but we have to stop the running server and restart it to see the changes. this is because go is a compiled language and we need to recompile the code to see the changes
//next we will add a database to our web server.