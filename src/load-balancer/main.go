package main

import (
	"context"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"os/signal"
	"syscall"

	"github.com/Flo0807/hsfl-master-ai-cloud-engineering/load-balancer/algorithm/round_robin.go"
	"github.com/Flo0807/hsfl-master-ai-cloud-engineering/load-balancer/balancer"
	"github.com/Flo0807/hsfl-master-ai-cloud-engineering/load-balancer/config"
	"github.com/Flo0807/hsfl-master-ai-cloud-engineering/load-balancer/docker_helper"
	"github.com/Flo0807/hsfl-master-ai-cloud-engineering/load-balancer/model"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
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

	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		panic(err)
	}
	defer cli.Close()

	err = docker_helper.CreateNetworkIfNotExists(cli, cfg.NetworkName)
	if err != nil {
		panic(err)
	}

	reader, err := cli.ImagePull(context.Background(), cfg.Image, types.ImagePullOptions{})
	if err != nil {
		panic(err)
	}
	io.Copy(os.Stdout, reader)

	containerIds := docker_helper.StartContainers(cli, cfg.Image, cfg.NetworkName, cfg.Replicas)
	defer docker_helper.RemoveContainers(cli, containerIds)

	targets, err := GetTargets(cli, cfg.NetworkName, containerIds, cfg.Port)
	if err != nil {
		panic(err)
	}

	loadBalancer := balancer.NewBalancer(round_robin.New(), targets, cfg.HealthCheckIntervalSeconds, cfg.HealthCheckPath)

	log.Println("Starting HTTP server on port", cfg.HttpServerPort)
	server := &http.Server{
		Addr:    fmt.Sprintf("0.0.0.0:%s", cfg.HttpServerPort),
		Handler: loadBalancer,
	}

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		<-sigChan
		server.Shutdown(context.Background())
	}()

	server.ListenAndServe()
}

func GetTargets(cli *client.Client, networkName string, containerIDs []string, port int) ([]model.Target, error) {
	targets := make([]model.Target, len(containerIDs))

	for i, containerID := range containerIDs {
		containerJson, err := cli.ContainerInspect(context.Background(), containerID)
		if err != nil {
			return nil, err
		}

		ip := containerJson.NetworkSettings.Networks[networkName].IPAddress

		urlStr := fmt.Sprintf("http://%s:%d", ip, port)

		targetUrl, err := url.Parse(urlStr)
		if err != nil {
			return nil, err
		}

		target := model.Target{
			ContainerId: containerID,
			Url:         targetUrl,
		}

		targets[i] = target
	}

	return targets, nil
}
