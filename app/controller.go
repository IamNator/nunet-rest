package app

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"nunet/pkg"
)

type Controller struct {
	Peers map[string]string // map of peer IDs to URLs
	Addrs []string
}

func (ctrl Controller) HandleHealthRequest(c *gin.Context) {

	availableCompute, err := pkg.GetComputeAvailable()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"error":   "Error getting compute availability",
			"details": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "Healthy",
		"data": gin.H{
			"id":        "1234",
			"addresses": []string{},
			"peers":     ctrl.Peers,
			"num_peers": len(ctrl.Peers),
			"network":   "libp2p",
			"cpu":       availableCompute.FreeCPUCores,
			"ram":       availableCompute.FreeRAM,
			"total_cpu": availableCompute.TotalCPUCores,
			"total_ram": availableCompute.TotalRAM,
			"cpu_model": availableCompute.TotalCPUModel,
			"cpu_ghz":   availableCompute.ToalCPUGhz,
		},
	})
}

// Peer represents a peer machine
type Peer struct {
	ID  string `json:"id" binding:"required"`
	URL string `json:"url" binding:"required"`
}

// HandleRegisterPeer handles requests to register a new peer
func (ctrl Controller) HandleRegisterPeer(c *gin.Context) {
	var peer Peer
	if err := c.ShouldBindJSON(&peer); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Add the peer to the list of peers
	ctrl.Peers[peer.ID] = peer.URL

	// Respond with success message
	c.JSON(http.StatusOK, gin.H{"message": "Peer registered successfully"})
}

// Job represents the job parameters sent by the user
type Job struct {
	Program   string   `json:"program"`
	Arguments []string `json:"arguments"`
}

type JobResponse struct {
	Output []string `json:"output"`
	PID    int      `json:"pid"`
}

// HandleJob handles requests to deploy a job/container
func (ctrl Controller) HandleJob(c *gin.Context) {
	var job Job
	if err := c.ShouldBindJSON(&job); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Process the job deployment (e.g., send to peer, etc.)
	output, pid, err := pkg.RunCmd(job.Program, job.Arguments...)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
			"data": gin.H{
				"output": output,
				"pid":    pid,
			},
		})
		return
	}

	// Respond with success message
	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "Job deployed successfully",
		"data": JobResponse{
			Output: output,
			PID:    pid,
		},
	})
}

func (ctrl Controller) HandleJobRequest(c *gin.Context) {
	var job Job
	if err := c.ShouldBindJSON(&job); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// check if there are any peers available
	if len(ctrl.Peers) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No peers available to deploy job"})
		return
	}

	var selectedPeer Peer
	for id, url := range ctrl.Peers {
		selectedPeer = Peer{
			ID:  id,
			URL: url,
		}
		break
	}

	// send job to selected peer
	var response struct {
		Data JobResponse `json:"data"`
	}
	statusCode, err := pkg.SendJobToPeer(selectedPeer.URL, response, job.Program, job.Arguments...)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": "error",
			"error":  err.Error(),
		})
		return
	}

	// Respond with success message
	c.JSON(http.StatusOK, gin.H{
		"message": "Job deployed successfully",
		"data": gin.H{
			"output": response.Data.Output,
			"pid":    response.Data.PID,
			"peer":   selectedPeer,
			"status": statusCode,
		},
	})
}
