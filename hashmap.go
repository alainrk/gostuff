package main

import (
    "fmt"
    "bytes"
    "strconv"
)

// type LinkedLister interface {
//     PushHead()
//     Remove()
//     Length()
//     Visit()
// }

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

    ll.head = &n
    ll.length += 1
    fmt.Println(ll.head)

    return ll.length
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
    ll.PushBack("Back1", 4)
    ll.PushBack("Back2", 5)
    ll.PushBack("Back3", 6)

    visit := ll.Visit()
    fmt.Printf(visit)
}
