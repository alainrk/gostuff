package main

import "fmt"

type LinkedLister interface {
    Add()
    Remove()
    Length()
}

type node struct {
    key string
    value int
    next *node
}

type LinkedList struct {
    head *node
    length int
}

func newNode(key string, value int) *node {
    n := node{key: key, value: value}
    n.next = nil
    return &n
}

// func newLinkedList() *LinkedList {
//     ll := LinkedList{}
//     return &ll
// }

func main() {
    ll := LinkedList{}
    n := newNode("bobbolo", 666)
    ll.head = n
    fmt.Println(n)
    fmt.Println(ll)
}

// func main() {
//     fmt.Println(node{"Bob", 20})
//     fmt.Println(&node{key: "Ann", value: 40})
//     n := node{key: "Sean", value: 666}
//     fmt.Println(n.name)

//     np := &n
//     fmt.Println(np.age)

//     np.value = 51
//     fmt.Println(np.value)
// }