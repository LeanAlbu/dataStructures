package main

import (
	"fmt"
	"math/rand/v2"
	"strconv"
)

type Node struct {
	key   int
	value int
	next  *Node
}

type HashTable struct {
	table     []*Node
	size      int
	collisions int
}

func foldHash(key int, size int) int {
	sum := 0
	keyStr := fmt.Sprintf("%d", key)
	chunkSize := 2
	for i := 0; i < len(keyStr); i += chunkSize {
		end := i + chunkSize
		if end > len(keyStr) {
			end = len(keyStr)
		}
		chunk, _ := strconv.Atoi(keyStr[i:end])
		sum += chunk
	}
	return sum % size
}

func fnvHash(key int, size int) int {
	const (
		fnvPrime  = 16777619
		fnvOffset = 2166136261
	)
	hash := uint32(fnvOffset)
	keyStr := fmt.Sprintf("%d", key)
	for _, char := range keyStr {
		hash *= fnvPrime // Overflow é intencional (FNV-1)
		hash ^= uint32(char)
	}
	return int(hash % uint32(size))
}

func NewHashTable(size int) *HashTable {
	return &HashTable{
		table:     make([]*Node, size),
		size:      size,
		collisions: 0,
	}
}

func (h *HashTable) Insert(key int, value int, hashFunc func(int, int) int) {
	index := hashFunc(key, h.size)
	if h.table[index] != nil {
		h.collisions++ // Conta colisão inicial no índice
	}
	if h.table[index] == nil {
		h.table[index] = &Node{key: key, value: value}
		return
	}
	aux := h.table[index]
	for aux.next != nil {
		if aux.key == key {
			aux.value = value
			return
		}
		aux = aux.next
	}
	if aux.key == key {
		aux.value = value
	} else {
		aux.next = &Node{key: key, value: value}
	}
}

func generateRandomKey() int {
	return rand.IntN(10000000)
}

func testCollisions(tableSize int, entries int, hashFunc func(int, int) int, hashName string) {
	ht := NewHashTable(tableSize)
	usedKeys := make(map[int]bool) // Para rastrear duplicatas, se desejado
	for i := 0; i < entries; i++ {
		key := generateRandomKey()
		ht.Insert(key, i, hashFunc)
		usedKeys[key] = true
	}
	collisionRate := float64(ht.collisions) / float64(entries)
	fmt.Printf("Função: %s, Tamanho: %d, Entradas: %d, Colisões: %d, Taxa: %.2f%%\n",
		hashName, tableSize, entries, ht.collisions, collisionRate*100)
}

func main() {
	entryCounts := []int{1000, 10000, 1000000}
	tableSize := 10000

	fmt.Println("=== Testes com Método da Dobra ===")
	for _, entries := range entryCounts {
		testCollisions(tableSize, entries, foldHash, "Método da Dobra")
	}

	fmt.Println("\n=== Testes com FNV Hash ===")
	for _, entries := range entryCounts {
		testCollisions(tableSize, entries, fnvHash, "FNV Hash")
	}
}
