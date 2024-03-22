package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// Peer represents a peer machine
type Peer struct {
	ID       string `json:"id"`
	Hostname string `json:"hostname"`
	// Add more fields as needed, such as IP address, port, etc.
}

var peers []Peer

// Job represents the job parameters sent by the user
type Job struct {
	Program   string   `json:"program"`
	Arguments []string `json:"arguments"`
}

// HandleDescribeJob handles requests to describe a job
func HandleDescribeJob(w http.ResponseWriter, r *http.Request) {
	var job Job
	err := json.NewDecoder(r.Body).Decode(&job)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Validate job parameters (e.g., check if program exists, etc.)

	// Process the job description (e.g., store in database, etc.)

	// Respond with success message
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"message": "Job description received successfully"})
}

// HandleDeployJob handles requests to deploy a job/container
func HandleDeployJob(w http.ResponseWriter, r *http.Request) {
	// Perform deployment logic (e.g., deploy container on the other machine)

	// Respond with success or failure message
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"message": "Job deployed successfully"})
}

// HandleRegisterPeer handles requests to register a new peer
func HandleRegisterPeer(w http.ResponseWriter, r *http.Request) {
	var peer Peer
	err := json.NewDecoder(r.Body).Decode(&peer)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Validate peer data (e.g., check if ID or hostname already exists, etc.)

	// Add the peer to the list of peers
	peers = append(peers, peer)

	// Respond with success message
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"message": "Peer registered successfully"})
}

func main() {
	http.HandleFunc("/job/describe", HandleDescribeJob)
	http.HandleFunc("/job/deploy", HandleDeployJob)
	http.HandleFunc("/peer/register", HandleRegisterPeer)

	fmt.Println("Server is running on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
