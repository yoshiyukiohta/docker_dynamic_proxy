package main

import (
    "fmt"
    "github.com/fsouza/go-dockerclient"
    "github.com/garyburd/redigo/redis"
    "os"
    "strings"
    "time"
)

var (
    docker_endpoint string
    redis_endpoint string
)

func httpPort(port int64) bool {
    return port == 80 ||
	port == 3000 ||
	port == 5000 ||
	port == 8000 ||
	port == 8080
}

func getDests() map[string]string {
    dests := make(map[string]string)
    client, _ := docker.NewClient(docker_endpoint)
    opts := docker.ListContainersOptions{}
    containers, _ := client.ListContainers(opts)
    for _, container := range containers {
	for _, port := range container.Ports {
	    if port.IP == "0.0.0.0" &&
		httpPort(port.PrivatePort) &&
		port.Type == "tcp" {

		    inspect, _ := client.InspectContainer(container.ID)
		    containerName := strings.TrimLeft(inspect.Name, "/")
		    dest := fmt.Sprintf("%s:%d",
					inspect.NetworkSettings.IPAddress,
					port.PrivatePort)
		    dests[containerName] = dest
		}
	}
    }
    return dests
}

func setToRedis(dests map[string]string) {
    client, err := redis.Dial("tcp", redis_endpoint)
    if err != nil {
	fmt.Println("fail to connect redis server: ", err)
	return
    }
    defer client.Close()
    for name, dest := range dests {
	client.Do("SET", name, dest)
	    }
}

func main() {
    docker_endpoint = "unix:///var/run/docker.sock"
    redis_endpoint = "127.0.0.1:6379"
    // if os.Getenv("DOCKER_HOST") != "" {
    //	docker_endpoint = os.Getenv("DOCKER_HOST")
    // }
    if os.Getenv("REDIS_ENDPOINT") != "" {
	redis_endpoint = os.Getenv("REDIS_ENDPOINT")
    } else if os.Getenv("UPSTREAMS_PORT_6379_TCP_ADDR") != "" &&
	os.Getenv("UPSTREAMS_PORT_6379_TCP_PORT") != "" {

	    redis_endpoint = fmt.Sprintf("%s:%s",
					 os.Getenv("UPSTREAMS_PORT_6379_TCP_ADDR"),
					 os.Getenv("UPSTREAMS_PORT_6379_TCP_PORT"))
	}

    for {
	dests := getDests()
	setToRedis(dests)
	time.Sleep(10 * time.Second)
    }
}
