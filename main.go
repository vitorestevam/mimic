package main

import (
	"fmt"
	"log"
)

var Log string = `2023-04-24T10:15:27.123Z INFO kubelet: Starting container "web-server" with UUID "f9cc9ef3-3c62-4f51-bd29-1b1cb87f0ec8" in pod "my-app-1234-5678" with UUID "b791c38d-48eb-42c1-a212-2f8cf141d25e" on node "worker-1" with UUID "7fb18790-6d05-49c3-bdd3-73a3125e6b5e"
2023-04-24T10:15:27.456Z INFO kubelet: Container "web-server" with UUID "f9cc9ef3-3c62-4f51-bd29-1b1cb87f0ec8" in pod "my-app-1234-5678" with UUID "b791c38d-48eb-42c1-a212-2f8cf141d25e" is running on node "worker-1" with UUID "7fb18790-6d05-49c3-bdd3-73a3125e6b5e"
2023-04-24T10:15:32.789Z ERROR kubelet: Failed to pull image "my-registry/my-app:latest" for container "web-server" with UUID "f9cc9ef3-3c62-4f51-bd29-1b1cb87f0ec8" in pod "my-app-1234-5678" with UUID "b791c38d-48eb-42c1-a212-2f8cf141d25e" on node "worker-1" with UUID "7fb18790-6d05-49c3-bdd3-73a3125e6b5e": ImagePullBackOff
2023-04-24T10:15:32.789Z WARN kubelet: Restarting container "web-server" with UUID "f9cc9ef3-3c62-4f51-bd29-1b1cb87f0ec8" in pod "my-app-1234-5678" with UUID "b791c38d-48eb-42c1-a212-2f8cf141d25e" on node "worker-1" with UUID "7fb18790-6d05-49c3-bdd3-73a3125e6b5e"
2023-04-24T10:15:33.234Z INFO kubelet: Starting container "web-server" with UUID "3d85e9af-70a5-4aa2-aa2d-10a30b57c66b" in pod "my-app-1234-5678" with UUID "b791c38d-48eb-42c1-a212-2f8cf141d25e" on node "worker-2" with UUID "e0eb24df-8e77-4712-8d6c-87d2c8bf0e32"
`

func main() {
	types := []string{"date", "uuid"}

	replacers := []Replacer{}

	for _, t := range types {
		replacer, err := GetReplacers(t)
		if err != nil {
			log.Println(err)
		}

		replacers = append(replacers, replacer)
	}

	text := Log

	fmt.Println(text)

	for _, r := range replacers {
		text = ApplyReplacer(r, text)
	}

	fmt.Println(text)
}
