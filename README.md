Certainly! Here's the modified description:

---

# Project Name

## Overview

This project is a job/container manager application built with GoLang that facilitates communication between nodes. It **utilizes RESTful communication** for interaction between servers/peers. The application allows nodes to connect, exchange messages, and discover other peers on the network.

## Getting Started

### Prerequisites

Make sure you have Go installed on your machine.

### Installation

1. Clone the repository:

```bash
git clone git@github.com:IamNator/nunet-rest.git
```

2. Navigate to the project directory:

```bash
cd nunet-rest
```

### Usage

To start the project, run the following command:

```bash
go run main.go
```

This command will start the application.

### Adding Peers

To add peers to the network, make a POST request to the appropriate endpoint. Example:

```bash
curl -X POST -d '{"id": "peer_id", "url": "peer_address"}' http://localhost:8080/peer
```

Replace `"peer_id"` with the identifier of the peer and `"peer_address"` with the address (url) of the peer you want to add.

### Running a Job

To run a job on the network, make a POST request to the job endpoint with the payload containing the program and arguments. Example:

```bash
curl -X POST -d '{"program": "echo", "arguments": ["hello"]}' http://localhost:8080/job
```

Replace `"echo"` with the program you want to execute, and `["hello"]` with the arguments to pass to the program.


### Checking Health

To check the health of the node and view details, make a GET request to the health endpoint. Example:

```bash
curl http://localhost:8080/health
```

This will display the health status and other details of the node.


#### Expected Response:

**Success (200 OK):**
```json
{
  "status": "success",
  "message": "Success message describing the action"
}

```
```bash
➜  nunet-rest git:(main) ✗ curl -X POST -d '{"id": "peer_id", "url": "peer_address"}' http://localhost:8080/peer -v
Note: Unnecessary use of -X or --request, POST is already inferred.
*   Trying [::1]:8080...
* Connected to localhost (::1) port 8080
> POST /peer HTTP/1.1
> Host: localhost:8080
> User-Agent: curl/8.4.0
> Accept: */*
> Content-Length: 40
> Content-Type: application/x-www-form-urlencoded
> 
< HTTP/1.1 200 OK
< Access-Control-Allow-Headers: Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization
< Access-Control-Allow-Methods: GET, POST, PUT, DELETE, OPTIONS
< Access-Control-Allow-Origin: *
< Content-Type: application/json; charset=utf-8
< Date: Mon, 25 Mar 2024 16:47:27 GMT
< Content-Length: 61
< 
* Connection #0 to host localhost left intact
{"message":"Peer registered successfully","status":"success"}%       
```

**Error (4xx/5xx):**
```json
{
  "status": "error",
  "error": "Error message describing the issue."
}
```
```bash
➜  nunet-rest git:(main) ✗ curl -X POST -d '{"id": "peer_id", "address": "peer_address"}' http://localhost:8080/peer -v
Note: Unnecessary use of -X or --request, POST is already inferred.
*   Trying [::1]:8080...
* Connected to localhost (::1) port 8080
> POST /peer HTTP/1.1
> Host: localhost:8080
> User-Agent: curl/8.4.0
> Accept: */*
> Content-Length: 44
> Content-Type: application/x-www-form-urlencoded
> 
< HTTP/1.1 400 Bad Request
< Access-Control-Allow-Headers: Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization
< Access-Control-Allow-Methods: GET, POST, PUT, DELETE, OPTIONS
< Access-Control-Allow-Origin: *
< Content-Type: application/json; charset=utf-8
< Date: Mon, 25 Mar 2024 16:48:35 GMT
< Content-Length: 89
< 
* Connection #0 to host localhost left intact
{"error":"Key: 'Peer.URL' Error:Field validation for 'URL' failed on the 'required' tag"}%      
```

In case of success, the response will include a JSON object indicating success along with a message. In case of failure, an error message will be provided explaining the issue.

## Contributing

Contributions are welcome! Feel free to open an issue or submit a pull request.

## License

This project is licensed under the [MIT License](LICENSE).