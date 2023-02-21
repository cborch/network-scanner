package main

import (
	"fmt"
	"log"
	"net"

	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"github.com/google/gopacket/pcap"
)

func main() {
	graph := new(NetworkNodeGraph)

	listNetworkDevices()
	packetChannel := createPacketSource("en0")
	readPackets(packetChannel, graph)
	fmt.Println(graph.vertices)
}

func listNetworkDevices() {
	devices, err := pcap.FindAllDevs()
	if err != nil {
		log.Fatal("Failed to find network devices", err)
	}

	for _, device := range devices {
		fmt.Println("Network Device:", device.Name)
		fmt.Println("Using IPs:")
		for _, address := range device.Addresses {
			fmt.Println(address.IP)
		}
		fmt.Println()
	}
}

func createPacketSource(networkDevice string) chan gopacket.Packet {
	handle, err := pcap.OpenLive(networkDevice, 1600, true, pcap.BlockForever)
	if err != nil {
		log.Fatal("Failed to establish live handle", err)
	}

	// Create a BPF Filter to narrow the scope of packets we'll have in our packet source
	err = handle.SetBPFFilter("")
	if err != nil {
		log.Fatal("Failed to attach BPF filter to handle", err)
	}

	// Create the packet source
	packetSource := gopacket.NewPacketSource(handle, handle.LinkType())
	packetChannel := packetSource.Packets()
	return packetChannel
}

func readPackets(packetChannel chan gopacket.Packet, graph *NetworkNodeGraph) {
	for packet := range packetChannel {
		ipv4Layer := packet.Layer(layers.LayerTypeIPv4)
		ipv6Layer := packet.Layer(layers.LayerTypeIPv6)

		if ipv4Layer != nil || ipv6Layer != nil {
			src, dst := getPacketSrceDst(packet)
			addToGraph(graph, src, dst)
		}

		fmt.Println("___________________________")
		graph.printEdges()
	}
}

func getPacketSrceDst(networkPacket gopacket.Packet) (string, string) {
	netFlow := networkPacket.NetworkLayer().NetworkFlow()

	src, dst := netFlow.Endpoints()
	srcName := lookupName(src.String())
	dstName := lookupName(dst.String())
	return srcName, dstName
}

func addToGraph(graph *NetworkNodeGraph, srcName string, dstName string) {
	if graph.contains(srcName) {
		graph.addEdge(srcName, dstName)
	} else {
		graph.addVertex(srcName)
	}
}

func lookupName(address string) string {
	names, err := net.LookupAddr(address)
	if err != nil {
		return address
	}
	return names[0]
}
