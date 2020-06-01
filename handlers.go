package main

import (
	"fmt"
	portscanner "github.com/anvie/port-scanner"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"time"
)

func HealthCheck(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	response := fmt.Sprintf("{Status: \"UP\"}")
	writeOKResponse(w, response)
}

func ScanLocalHost(w http.ResponseWriter, r *http.Request, _ httprouter.Params){
	fmt.Fprint(w, "Scanning Beginning!\n")
	// scan localhost with a 2 second timeout per port in 5 concurrent threads
	ps := portscanner.NewPortScanner("localhost", 2*time.Second, 5)

	// get opened port
	fmt.Printf("scanning port %d-%d...\n", 20, 30000)

	openedPorts := ps.GetOpenedPort(20, 30000)

	for i := 0; i < len(openedPorts); i++ {
		port := openedPorts[i]
		fmt.Print(" ", port, " [open]")
		fmt.Println("  -->  ", ps.DescribePort(port))
	}
	writeOKResponse(w, openedPorts)
}

func TestJson(w http.ResponseWriter, r *http.Request, _ httprouter.Params){
	 var x string = "Test Json"
	 writeOKResponse(w, x)
}

