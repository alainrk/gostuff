package main

import (
    "fmt"
    "bytes"
    "strconv"
)

type LinkedLister interface {
    PushHead()
    PushBack()
    Head()
    Tail()
    Remove()
    Length()
    Visit()
}

type Node struct {
    key string
    value int
    next *Node
}

type LinkedList struct {
    head *Node
    tail *Node
    length int
}

// Gonna use pointer receiver here to mutate all the stuff
func (ll *LinkedList) PushHead(key string, value int) int {
    n := Node{key, value, ll.head}

    if ll.head == nil {
        ll.tail = &n
    }
    ll.head = &n
    ll.length += 1

    return ll.length
}

// Gonna use pointer receiver here to mutate all the stuff
func (ll *LinkedList) PushBack(key string, value int) int {
    n := Node{key, value, nil}

    if ll.head == nil {
        ll.head = &n
    } else {
        ll.tail.next = &n
    }
    ll.tail = &n
    ll.length += 1

    return ll.length
}

// Remove all the instances O(n)
func (ll *LinkedList) Remove(key string) {
    if ll.head == nil {
        return
    }
    if ll.head.key == key {
        ll.head = ll.head.next
        ll.length -= 1
    }
    var prev *Node = nil
    curr := ll.head

    for curr != nil {
        if curr.key == key {
            prev.next = curr.next
            ll.length -= 1
        }
        prev = curr
        curr = curr.next
    }
}

func (ll LinkedList) Length() int {
    return ll.length
}

func (ll LinkedList) Head() *Node {
    return ll.head
}

func (ll LinkedList) Tail() *Node {
    return ll.tail
}

// Gonna use value receiver here because i'm just reading
func (ll LinkedList) Visit() string {
    var visit bytes.Buffer
    visit.WriteString("Visit: [ ")
    curr := ll.head
    for curr != nil {
        visit.WriteString(curr.key)
        visit.WriteString(": ")
        visit.WriteString(strconv.Itoa(curr.value))
        if curr.next != nil {
            visit.WriteString(",")
        }
        visit.WriteString(" ")
        curr = curr.next
    }
    visit.WriteString("]")
    var res string = visit.String()
    return res
}

func main() {
    ll := LinkedList{}
    ll.PushHead("Head1", 3)
    ll.PushHead("Head2", 2)
    ll.PushHead("Head3", 1)

    fmt.Println("Tail:", strconv.Itoa(ll.Tail().value))

    ll.PushBack("Back1", 4)
    ll.PushBack("Back2", 5)
    ll.PushBack("Back3", 6)

    visit := ll.Visit()
    fmt.Println(visit)

    ll.Remove("Head2")
    ll.Remove("Back2")
    ll.Remove("Head3")

    visit = ll.Visit()
    fmt.Println(visit)

    ll.Remove("Back3")
    ll.Remove("Back1")
    ll.Remove("Head1")

    visit = ll.Visit()
    fmt.Println(visit)

    head := ll.Head()
    tail := ll.Tail()
    if head != nil {
        fmt.Println("Head:", strconv.Itoa(head.value))
        fmt.Println("Tail:", strconv.Itoa(tail.value))
    }
    fmt.Println("Length:", ll.Length())
}
