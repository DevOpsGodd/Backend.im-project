package main

import (
	"log"
	"websocket-git/internal/handlers"
	"websocket-git/internal/services"

	"github.com/gin-gonic/gin"
)

func main() {
	// Set Gin to debug mode
	gin.SetMode(gin.DebugMode)

	// Initialize services
	namespaceService := services.NewNamespaceService()

	// Initialize handlers
	wsHandler := handlers.NewWebSocketHandler(namespaceService)

	// Setup router with default middleware (Logger and Recovery)
	r := gin.Default()
	r.GET("/ws", wsHandler.HandleConnection)

	log.Printf("WebSocket server starting on :8080")
	if err := r.Run(":8080"); err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
