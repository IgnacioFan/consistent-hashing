/*
Playground: https://go.dev/play/p/hVoE2jQxQOB
*/
package main

import (
	"fmt"
	"hash/crc32"
	"sort"
	"strconv"
)

func main() {
	ch := NewConsistentHash(3)
	// Add 3 physical nodes respectively.
	// Each physical node has 3 replicas,
	// so a total of nine virtual nodes are created.
	ch.AddNode(NewNode(1, "192.168.1.1", 3))
	ch.AddNode(NewNode(2, "192.168.1.2", 3))
	ch.AddNode(NewNode(3, "192.168.1.3", 3))

	fmt.Println(ch.GetNode("key1")) // 192.168.1.1
	fmt.Println(ch.GetNode("key2")) // 192.168.1.3
	fmt.Println(ch.GetNode("key3")) // 192.168.1.1
	fmt.Println(ch.GetNode("key4")) // 192.168.1.2
	fmt.Println(ch.GetNode("key4")) // 192.168.1.2
}

// store the hashes of the virtual nodes
type HashRing []uint32

// represents a physical node in the cluster
type Node struct {
	ID       int
	IP       string
	Hash     uint32
	Replicas int
}

func NewNode(id int, ip string, replicas int) *Node {
	node := &Node{
		ID:       id,
		IP:       ip,
		Replicas: replicas,
	}
	node.Hash = crc32.ChecksumIEEE([]byte(node.IP))
	return node
}

// Nodes: store the virtual nodes
// Replicas: specify the number of replicas for each physical node
// HashRing: store the hashes of the virtual nodes
// VirtualNode:  map virtual node hashes to IP addresses
type ConsistentHash struct {
	Nodes       map[uint32]*Node
	Replicas    int
	HashRing    HashRing
	VirtualNode map[uint32]string
}

func NewConsistentHash(replicas int) *ConsistentHash {
	return &ConsistentHash{
		Nodes:       make(map[uint32]*Node),
		Replicas:    replicas,
		HashRing:    HashRing{},
		VirtualNode: make(map[uint32]string),
	}
}

// adds a new physical node to the cluster by creating multiple
// virtual nodes and adding them to the hash ring and the Nodes
// and VirtualNode maps
func (ch *ConsistentHash) AddNode(node *Node) {
	for i := 0; i < ch.Replicas; i++ {
		virtualKey := strconv.Itoa(node.ID) + "_" + strconv.Itoa(i)
		virtualHash := crc32.ChecksumIEEE([]byte(virtualKey))
		ch.Nodes[virtualHash] = node
		ch.VirtualNode[virtualHash] = node.IP
		ch.HashRing = append(ch.HashRing, virtualHash)
	}
}

// takes a key and uses its hash to find the nearest virtual
// node in the hash ring. It then returns the IP address of
// the physical node that owns that virtual node.
func (ch *ConsistentHash) GetNode(key string) string {
	hash := crc32.ChecksumIEEE([]byte(key))
	idx := sort.Search(len(ch.HashRing), func(i int) bool {
		return ch.HashRing[i] >= hash
	})
	if idx == len(ch.HashRing) {
		idx = 0
	}
	virtualHash := ch.HashRing[idx]
	return ch.VirtualNode[virtualHash]
}
