package main

import (
	"context"
	"fmt"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
	"net"
	"time"
)

type PingResult struct {
	ContainerID string `json:"container_id"`
	IPAddress   string `json:"ip_address"`
	PingTime    int    `json:"ping_time"` // Время ответа (мс)
	LastSuccess string `json:"last_success"`
}

func pingContainers() {
	Client, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		fmt.Println("Error connecting to Docker: ", err)
		return
	}
	defer Client.Close()

	containers, err := Client.ContainerList(context.Background(), container.ListOptions{All: true})
	if err != nil {
		fmt.Println("Error getting the container list: ", err)
		return
	}

	for _, ctr := range containers {
		ip := getContainerIP(Client, ctr.ID)
		if ip == "" {
			continue
		}

		pingTime := ping(ip, 80)
		sendPingResult(ip, pingTime)

	}
}

func ping(ip string, port int) int {
	start := time.Now()
	conn, err := net.DialTimeout("tcp", fmt.Sprintf("%s:%d", ip, port), time.Second)
	if err != nil {
		return -1
	}
	conn.Close()
	return int(time.Since(start).Milliseconds())
}

//func ping(ip string, ports nat.PortMap) bool {
//	for _, bindings := range ports {
//		if bindings != nil && len(bindings) > 0 {
//			conn, err := net.DialTimeout("tcp", ip+":"+bindings[0].HostPort, time.Second)
//			if err != nil {
//				return false
//			}
//			conn.Close()
//			return true
//		}
//	}
//	return false
//}

func getContainerIP(Client *client.Client, containerID string) string {
	containerInfo, err := Client.ContainerInspect(context.Background(), containerID)
	if err != nil {
		fmt.Println("Error receiving the container's IP address: ", err)
		return ""
	}
	return containerInfo.NetworkSettings.IPAddress
}

func sendPingResult(ip string, pingTime int) {
	result := PingResult{
		IPAddress:   ip,
		PingTime:    pingTime,
		LastSuccess: time.Now().Format(time.RFC3339),
	}
	fmt.Println(result.ContainerID, result.IPAddress, result.PingTime, result.LastSuccess)
}
