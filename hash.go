package myhash

import "fmt"

const arraySize = 7

// hash_table
type HashTable struct {
	array [arraySize]*bucket
}

// insert data into the hash array
func (h *HashTable) Insert(k string) {
	index := hash(k)
	h.array[index].insert(k)
}

// search a key into the array
func (h *HashTable) Search(k string) {
	index := hash(k)
	if h.array[index].search(k) {
		fmt.Println(k, "is here!")
	} else {
		fmt.Println(k, "is not here..")
	}
}

// delete an element from the array
func (h *HashTable) Delete(k string) {
	index := hash(k)
	h.array[index].delete(k)
}

// bucket
type bucket struct {
	head   *bucket_node
	length int
}

// insert data into the bucket
func (b *bucket) insert(k string) {
	if !b.search(k) {
		newNode := &bucket_node{data: k}
		newNode.next = b.head
		b.head = newNode
		b.length++
	} else {
		fmt.Println(k, "already in..")
	}
}

// search for a key in the bucket
func (b *bucket) search(k string) bool {
	currentNode := b.head
	for currentNode != nil {
		if currentNode.data == k {
			return true
		}
		currentNode = currentNode.next
	}
	return false
}

// delete an element from the bucket (linked list)
func (b *bucket) delete(k string) {
	// if no data yet
	if b.length == 0 {
		fmt.Println("There is no one here..yet..")
		return
	}
	// if it's the head
	if b.head.data == k {
		b.head = b.head.next
		b.length--
		return
	}

	previousNode := b.head
	for previousNode.next != nil {
		if previousNode.next.data == k {
			//delete
			previousNode.next = previousNode.next.next
			b.length--
			return
		}
		previousNode = previousNode.next
	}
}

// bucket_node
type bucket_node struct {
	data string
	next *bucket_node
}

// key to index
func hash(k string) int {
	sum := 0
	for _, v := range k {
		sum += int(v)
	}
	return sum % arraySize
}

// init the hash table creating bucket in each slot of the array
func Init() *HashTable {
	hash_table := &HashTable{}
	for i := range hash_table.array {
		hash_table.array[i] = &bucket{}
	}
	return hash_table
}

func Help() {
	fmt.Println(`
	insert -> insert data
	search -> search for data
	delete -> delete data
	exit -> quit
	`)
}
