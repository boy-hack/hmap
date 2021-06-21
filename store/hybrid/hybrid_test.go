package hybrid

import (
	"fmt"

	"testing"
)

func TestHybrid(t *testing.T) {
	hm, err := New(DefaultDiskOptions)
	if err != nil {
		fmt.Errorf(err.Error())
		return
	}
	hm.Set("test11", nil)
	hm.Set("test12", nil)
	hm.Set("test13", nil)
	hm.Set("test14", nil)
	hm.Set("test15", nil)
	t.Log(hm.Empty())

	hm.Scan(func(bytes []byte, bytes2 []byte) error {
		t.Log(string(bytes), bytes2)
		return nil
	})
	hm.Del("test11")
	t.Log("---------------------")
	hm.Scan(func(bytes []byte, bytes2 []byte) error {
		t.Log(string(bytes), bytes2)
		return nil
	})
	hm.Del("test12")
	hm.Del("test13")
	hm.Del("test14")
	hm.Del("test15")
	t.Log(hm.Empty())

}
