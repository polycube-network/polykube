package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/containernetworking/cni/libcni"
	"log"

	//	"github.com/containernetworking/cni/pkg/invoke"
	"os"
)

// session is maintained in /var/lib/networks/network_name

func main() {
	cniContainerId := os.Getenv("CNI_CONTAINERID")
	cniNetns := os.Getenv("CNI_NETNS")
	cniIfname := os.Getenv("CNI_IFNAME")
	cniPath := os.Getenv("CNI_PATH")
	netConfPath := os.Getenv("NETCONF")

	paths := []string{cniPath}
	cniConf := libcni.NewCNIConfigWithCacheDir(paths, "/tmp/cacheDir", nil)

	netConfList, err := libcni.ConfListFromFile(netConfPath)
	if err != nil {
		log.Fatalf("main: parsing netConfList: %v\n", err)
		return
	}

	if	_, err := cniConf.ValidateNetworkList(context.TODO(), netConfList); err != nil {
		log.Fatalf("main: validating network list: %v\n", err)
		return
	}

	runtimeConf := &libcni.RuntimeConf{
		ContainerID: cniContainerId,
		NetNS:       cniNetns,
		IfName:      cniIfname,
	}

	log.Println("Press the Enter Key to perform ADD")
        fmt.Scanln()


	result, err := cniConf.AddNetworkList(context.TODO(), netConfList, runtimeConf)
	if err != nil {
		log.Fatalf("main: adding network: %v\n", err)
		return
	}
	formattedResult, err := json.MarshalIndent(result, "", "\t")
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Println(string(formattedResult))


	log.Print("ADD done")
	log.Println("Press the Enter Key to perform CHECK")
	fmt.Scanln()

	if err = cniConf.CheckNetworkList(context.TODO(), netConfList, runtimeConf); err != nil {
		log.Fatalf("main: checking network: %v\n", err)
		return
	}
	log.Print("CHECK done")
	log.Print("Press the Enter Key to perform DEL")
	fmt.Scanln()

	if err = cniConf.DelNetworkList(context.TODO(), netConfList, runtimeConf); err != nil {
		log.Fatalf("main: deleting network: %v\n", err)
		return
	}
	log.Print("DEL done")
}
