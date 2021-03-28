package main

import (
    "testing"
)

func TestLinkedList(t *testing.T) {
    ll := LinkedList{}
    ll.PushHead("Head1", 3)
    ll.PushHead("Head2", 2)
    ll.PushHead("Head3", 1)

    tail := ll.Tail().value
    if tail != 3 {
        t.Errorf("Wrong tail = %d; want 3", tail)
    }
}

