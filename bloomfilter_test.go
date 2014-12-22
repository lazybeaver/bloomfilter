package bloomfilter

import "testing"

func TestBloomFilter(t *testing.T) {
	items := []string{"foo", "bar", "baz", "qux"}
	bf := NewBloomFilter(25, 5)

	// Add items to the bloom filter and verify size
	for _, item := range items {
		bf.Add([]byte(item))
	}
	if bf.Size() != len(items) {
		t.Errorf("Unexpected BloomFilter size\nExpected: %d | Actual: %d", len(items), bf.Size())
	}

	// Verify that there are no false negatives
	for _, item := range items {
		if !bf.Contains([]byte(item)) {
			t.Errorf("BloomFilter reports false negative for %s", item)
		}
	}

	// Verify negative result for at least one known case
	if bf.Contains([]byte("quux")) {
		t.Errorf("BloomFilter reports unexpected false positive")
	}
}
