package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"os/exec"
	"fmt"
)

func PowerEndpoint(w http.ResponseWriter, req *http.Request) {
	fmt.Println(">>> Sending LIRC KEY_POWER command...")
	out := map[string]bool{
		"success": true,
	}
	cmd := "irsend"
	args := []string{"SEND_ONCE", "roku", "KEY_POWER"}
	lsCmd := exec.Command(cmd, args...)
	lsOut, err := lsCmd.Output()
	if err != nil {
		fmt.Println("Unable to run command; Error: ", err)
		out["success"] = false
	}
	fmt.Println(string(lsOut))

	json.NewEncoder(w).Encode(out)
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/power", PowerEndpoint).Methods("GET")
	log.Fatal(http.ListenAndServe(":5000", router))
}