package main

import(
	"flag"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"os"
)

func main() {
	// get the current working directory
	wd, _ := os.Getwd()

	// set and parse directory flags
	env := flag.String("env", "", "environment variable that correlates with the directory in which the index.html file is located")
	dir := flag.String("dir", "", "directory where the index.html file is located")
	port := flag.String("port", "8080", "port where the application is to be served")
	flag.Parse()

	// create router
	r := mux.NewRouter()

	// look for directory
	if(*env != "" && *dir != "") {
		if _, err := os.Stat(fmt.Sprintf("%s/index.html", *dir)); err == nil {
			r.Handle("/", http.FileServer(http.Dir(*dir)))
			runInfo(wd, *port)

		}
		if _, err := os.Stat(fmt.Sprintf("%s/index.html", *env)); err == nil {
			r.Handle("/", http.FileServer(http.Dir(*env)))
			runInfo(wd, *port)

		} else {
			fmt.Println("-------------------------------------------------------------------------------")
			fmt.Println("")
			fmt.Println("index.html is not located within the input directory")
			fmt.Println("")
			fmt.Println("-------------------------------------------------------------------------------")

			os.Exit(1)
		}
	} else {
		if _, err := os.Stat(fmt.Sprintf("%s/index.html", wd)); err == nil {
			r.Handle("/", http.FileServer(http.Dir(wd)))
			runInfo(wd, *port)

		} else {
			warn()

			os.Exit(1)
		}
	}

	// serve router
	http.ListenAndServe(fmt.Sprintf(":%s", *port), r)
}


func warn() {
	fmt.Println("")
	fmt.Println("")
	fmt.Println("-------------------------------------------------------------------------------")
	fmt.Println("")
	fmt.Println("If index.html is not in the current working directory, please enter the")
	fmt.Println("environment variable that points to the directory where the file is located")
	fmt.Println("")
	fmt.Println("Example: 'AppServe -env=APP_DIR'")
	fmt.Println("")
	fmt.Println("")
	fmt.Println("OR")
	fmt.Println("")
	fmt.Println("")
	fmt.Println("Enter the actual directory where the index.html file is located")
	fmt.Println("")
	fmt.Println("Example: 'AppServe -dir=/User/KodeeMcIntosh/static/'")
	fmt.Println("-------------------------------------------------------------------------------")
}

func runInfo(directory string, port string) {
	fmt.Printf("AppServe is serving %s/index.html on port %s", directory, port)
	// TODO: add some actual logging
}
