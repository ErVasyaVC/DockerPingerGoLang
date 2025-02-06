package main

import (
	"context"
	"fmt"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
)

func main() {
	apiClient, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		panic(err)
	}
	defer apiClient.Close()

	containers, err := apiClient.ContainerList(context.Background(), container.ListOptions{All: true})
	if err != nil {
		panic(err)
	}

	for _, ctr := range containers {
		cntr, err := apiClient.ContainerInspect(context.Background(), ctr.ID)
		if err != nil {
			panic(err)
		} else {
			fmt.Println("Имя контейнера:", cntr.Name)
			fmt.Println("Статус контейнера:", cntr.State.Status)
			fmt.Println("IP-адрес контейнера:", cntr.NetworkSettings.IPAddress)
			fmt.Println()
		}

	}
}
