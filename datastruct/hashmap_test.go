package datastruct

import (
    "testing"
)

func TestHashMap(t *testing.T) {
    hm := HashMap{}
    hm.Add("key1", 3)

	val := hm.Get("key1")
    if val != 3 {
        t.Errorf("Wrong value [key1] = %d; want 3", val)
    }
}

