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

// HashTable com contador de colisões
type HashTable struct {
	table     []*Node
	size      int
	collisions int
}

//Folding Method
func foldHash(key int, size int) int {
	sum := 0
	keyStr := fmt.Sprintf("%d", key) // converte para string
	chunkSize := 2 //tamanho das dobras

	for i := 0; i < len(keyStr); i += chunkSize {
		end := i + chunkSize
		if end > len(keyStr) {
			end = len(keyStr)
		}
		chunk, _ := strconv.Atoi(keyStr[i:end])//converte de volta para int
		sum += chunk
	}
	return sum % size
}

// FNV Hash 32 bits
func fnvHash(key int, size int) int {
	const (
		fnvPrime  = 16777619
		fnvOffset = 2166136261
	)
	hash := uint32(fnvOffset) //truncador
	keyStr := fmt.Sprintf("%d", key)
	for _, char := range keyStr {
		hash *= fnvPrime
		hash ^= uint32(char)//XOR
	}
	return int(hash % uint32(size))
}

// Inicializa a HashTable
func NewHashTable(size int) *HashTable {
	return &HashTable{
		table:     make([]*Node, size),
		size:      size,
		collisions: 0,
	}
}

// Inserção com função de hash configurável
func (h *HashTable) Insert(key int, value int, hashFunc func(int, int) int) {
	index := hashFunc(key, h.size)

    //add colision
	if h.table[index] != nil {
		h.collisions++
	}

    //check empty
	if h.table[index] == nil {
		h.table[index] = &Node{key: key, value: value}
		return
	}

    //normal case
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

// Gera uma key random
func generateRandomKey() int {
	return rand.IntN(10000000) // Usando rand/v2
}

// Testa colisões das keys inseridas
func testCollisions(tableSize int, entries int, hashFunc func(int, int) int, hashName string) {
	ht := NewHashTable(tableSize)
	for i := range entries {
		key := generateRandomKey()
		ht.Insert(key, i, hashFunc)
	}
	fmt.Printf("Função: %s, Tamanho da tabela: %d, Entradas: %d, Colisões: %d\n", hashName, tableSize, entries, ht.collisions)
}

func main() {
	sizes := []int{1000, 10000, 1000000}
	tableSize := 10000

	fmt.Println("=== Testes com Método da Dobra ===")
	for _, entries := range sizes {
		testCollisions(tableSize, entries, foldHash, "Método da Dobra")
	}

	fmt.Println("\n=== Testes com FNV Hash ===")
	for _, entries := range sizes {
		testCollisions(tableSize, entries, fnvHash, "FNV Hash")
	}
}
