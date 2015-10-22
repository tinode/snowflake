package snowflake

import (
	"testing"
)

func TestNext(t *testing.T) {
	sf, err := NewSnowFlake(1)
	if err != nil {
		t.Error(err)
	}

	id, err := sf.Next()
	if err != nil {
		t.Error(err)
	}

	id2, err := sf.Next()
	if err != nil {
		t.Error(err)
	}

	if id >= id2 {
		t.Errorf("id %v is smaller or equal to previous one %v", id2, id)
	}
}

func TestDuplicate(t *testing.T) {

	total := 1000 * 1000
	data := make(map[uint64]int)

	sf, err := NewSnowFlake(1)
	if err != nil {
		t.Error(err)
	}

	var id, pre uint64
	for i := 0; i < total; i++ {

		id, err = sf.Next()
		if err != nil {
			t.Error(err)
		}

		if id < pre {
			t.Errorf("id %v is smaller than previous one %v", id, pre)
		}
		pre = id

		count := data[id]
		if count > 0 {
			t.Errorf("duplicate id %v %d", id, count)
		}
		data[id] = count + 1
	}

	length := len(data)
	t.Logf("map length %v", length)
	if length != total {
		t.Errorf("length does not match expected value; expected %v, actual %d", total, length)
	}

}

func BenchmarkNext(b *testing.B) {
	sf, err := NewSnowFlake(1)
	if err != nil {
		b.Error(err)
	}

	for i := 0; i < b.N; i++ {
		sf.Next()
	}
}
