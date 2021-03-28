package main

import (
    "hashmap"
)

func TestHashMap(t *testing.T) {
    hm := HashMap{}
    hm.add("key1", 3)

		val := hm.get("key1")
    if val != 3 {
        t.Errorf("Wrong value [key1] = %d; want 3", val)
    }
}

