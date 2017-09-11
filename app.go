package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"os/exec"
	"fmt"
	"time"
)

func PowerEndpoint(w http.ResponseWriter, req *http.Request) {
	fmt.Println(">>> Sending LIRC KEY_POWER command to roku remote...")
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

func PowerOffBluetoothEndpoint(w http.ResponseWriter, req *http.Request) {
	fmt.Println(">>> Sending LIRC KEY_POWER command to vizio remote...")
	out := map[string]bool{
		"success": true,
	}
	cmd := "irsend"
	args := []string{"SEND_ONCE", "count==2", "vizio", "KEY_POWER"}
	lsCmd := exec.Command(cmd, args...)
	_, err := lsCmd.Output()
	if err != nil {
		fmt.Println("Unable to run vizio KEY_POWER command; Error: ", err)
		out["success"] = false
	}

	json.NewEncoder(w).Encode(out)
}

func ConnectToBluetoothEndpoint(w http.ResponseWriter, req *http.Request) {
	fmt.Println(">>> Sending LIRC KEY_POWER command to vizio remote...")
	out := map[string]bool{
		"success": true,
	}
	cmd := "irsend"
	args := []string{"SEND_ONCE", "count==2", "vizio", "KEY_POWER"}
	lsCmd := exec.Command(cmd, args...)
	_, err := lsCmd.Output()
	if err != nil {
		fmt.Println("Unable to run vizio KEY_POWER command; Error: ", err)
		out["success"] = false
	}

	time.Sleep(3 * time.Second)

	cmd = "irsend"
	args = []string{"SEND_ONCE", "count==2", "vizio", "KEY_BLUE"}
	lsCmd = exec.Command(cmd, args...)
	_, err = lsCmd.Output()
	if err != nil {
		fmt.Println("Unable to run vizio KEY_BLUE command; Error: ", err)
		out["success"] = false
	}

	json.NewEncoder(w).Encode(out)
}

func main() {
	// instantiate server
	router := mux.NewRouter()

	// endpoints
	router.HandleFunc("/power", PowerEndpoint).Methods("GET")
	router.HandleFunc("/powerBluetooth", PowerOffBluetoothEndpoint).Methods("GET")
	router.HandleFunc("/connectToBluetooth", ConnectToBluetoothEndpoint).Methods("GET")

	log.Fatal(http.ListenAndServe(":5000", router))
}
