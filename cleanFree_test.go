package main

import "testing"

func TestCanParseRAMCommandOutput(t *testing.T) {
	raw := `
			  total        used        free      shared     buffers       cache   available
Mem:           15Gi       5.7Gi       5.9Gi       1.2Gi       434Mi       3.6Gi       8.6Gi
Swap:          15Gi          0B        15Gi`
	mem, swap := cleanFree(raw)

	if val, ok := mem["total"]; ok {
		if val != "15Gi" {
			t.Error("Parsed values for SWAP are incorrect")
		}
	} else {
		t.Error("the key 'total' was not set correctly")
	}

	if val, ok := mem["used"]; ok {
		if val != "5.7Gi" {
			t.Error("Parsed values for SWAP are incorrect")
		}
	} else {
		t.Error("the key 'used' was not set correctly")
	}

	if val, ok := mem["free"]; ok {
		if val != "5.9Gi" {
			t.Error("Parsed values for SWAP are incorrect")
		}
	} else {
		t.Error("the key 'free' was not set correctly")
	}

	if val, ok := mem["shared"]; ok {
		if val != "1.2Gi" {
			t.Error("Parsed values for SWAP are incorrect")
		}
	} else {
		t.Error("the key 'shared' was not set correctly")
	}

	if val, ok := swap["total"]; ok {
		if val != "15Gi" {
			t.Error("Parsed values for SWAP are incorrect")
		}
	} else {
		t.Error("the key 'total' was not set correctly")
	}
}

func TestCanParseRAMCommandOutputFromCentOSOutput(t *testing.T) {
	raw := `
			  total       used       free     shared    buffers     cached
Mem:          3.7G       2.7G       1.0G       1.2M       220M       949M
-/+ buffers/cache:       1.6G       2.2G
Swap:         1.0G        41M       982M
	`
	mem, swap := cleanFree(raw)

	if val, ok := mem["total"]; ok {
		if val != "3.7G" {
			t.Error("Parsed values for SWAP are incorrect")
		}
	} else {
		t.Error("the key 'total' was not set correctly")
	}

	if val, ok := mem["used"]; ok {
		if val != "2.7G" {
			t.Error("Parsed values for SWAP are incorrect")
		}
	} else {
		t.Error("the key 'used' was not set correctly")
	}

	if val, ok := mem["free"]; ok {
		if val != "1.0G" {
			t.Error("Parsed values for SWAP are incorrect")
		}
	} else {
		t.Error("the key 'free' was not set correctly")
	}

	if val, ok := mem["shared"]; ok {
		if val != "1.2M" {
			t.Error("Parsed values for SWAP are incorrect")
		}
	} else {
		t.Error("the key 'shared' was not set correctly")
	}

	if val, ok := swap["total"]; ok {
		if val != "1.0G" {
			t.Error("Parsed values for SWAP are incorrect")
		}
	} else {
		t.Error("the key 'total' was not set correctly")
	}
}
