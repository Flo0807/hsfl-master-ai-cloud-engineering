package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/Flo0807/hsfl-master-ai-cloud-engineering/src/api-gateway/config"
	"github.com/Flo0807/hsfl-master-ai-cloud-engineering/src/api-gateway/handler"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	configPath := flag.String("config", "./config.json", "Path to the configuration file")
	flag.Parse()

	cfg, err := config.LoadConfig(*configPath)

	if err != nil {
		log.Fatal("Error loading config file: ", err)
	}

	// Configuration for the Reverse Proxy
	var services []handler.Service

	for _, service := range cfg.Services {
		services = append(services, handler.Service{
			Name:        service.Name,
			ContextPath: service.ContextPath,
			TargetURL:   service.TargetURL,
		})
	}

	proxyCfg := handler.ReverseProxyConfig{
		Services: services,
	}

	// Create a new Reverse Proxy
	reverseProxy := handler.NewReverseProxy(proxyCfg)

	// Use the Reverse Proxy as a http.Handler in http.ListenAndServe
	http.Handle("/", reverseProxy)

	// Start the Reverse Proxy
	log.Printf("Starting HTTP server on port %s", cfg.HttpServerPort)
	addr := fmt.Sprintf("0.0.0.0:%s", cfg.HttpServerPort)
	err = http.ListenAndServe(addr, nil)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
	}
}
