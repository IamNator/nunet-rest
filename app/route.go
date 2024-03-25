package app

import (
	"fmt"
	"strings"

	"github.com/gin-gonic/gin"

	"nunet/pkg"
)

func Run(port int) error {

	router := gin.Default()
	router.Use(pkg.CorsMiddleware()) // attach cors middleware

	ctrl := Controller{
		Peers: make(map[string]string),
		Addrs: []string{
			fmt.Sprintf("http://localhost:%d", port),
		},
	}

	router.GET("/health", ctrl.HandleHealthRequest)
	router.POST("/peer", ctrl.HandleRegisterPeer)
	router.POST("/deploy", ctrl.HandleJob)

	// Start listening for incoming connections with port handling logic
	fmt.Println("Listening for deployment requests...")
retry:
	if err := router.Run(fmt.Sprintf(":%d", port)); err != nil {
		if strings.Contains(err.Error(), "already in use") {
			port++
			fmt.Printf("Port %d already in use, retrying with port %d\n", port-1, port)
			goto retry
		}
		return err
	}

	return nil
}
