package tiny

import (
	"testing"
)

var (
	url	  = 1337
	enc	  = "pjz2p"
	alphabet  = "mn6j2c4rv8bpygw95z7hsdaetxuk3fq"
	blockSize = 24
	minLength = 5
)

func TestEncodeUrl(t *testing.T) {
	tiny := New()
	expected := enc
	if result :=  tiny.EncodeUrl(url, 5); result != expected {
		t.Fatalf("EncodeUrl: unexpected output:\nExpected = %v\nResult = %v\n", expected, result)
	}
}

func TestDecodeUrl(t *testing.T) {
	tiny := New()
	expected := url
	if result :=  tiny.DecodeUrl(enc); result != expected {
		t.Fatalf("EncodeUrl: unexpected output:\nExpected = %v\nResult = %v\n", expected, result)
	}
}

func TestEncode(t *testing.T) {
	tiny := New()
	expected := 15482880
	if result :=  tiny.encode(567); result != expected {
		t.Fatalf("EncodeUrl: unexpected output:\nExpected = %v\nResult = %v\n", expected, result)
	}
}

func TestDecode(t *testing.T) {
	tiny := New()
	expected := 6635520
	if result :=  tiny.decode(678); result != expected {
		t.Fatalf("EncodeUrl: unexpected output:\nExpected = %v\nResult = %v\n", expected, result)
	}

}

func TestEnbase(t *testing.T) {
	tiny := New()
	expected := "mmmcw"
	if result :=  tiny.enbase(169, minLength); result != expected {
		t.Fatalf("EncodeUrl: unexpected output:\nExpected = %v\nResult = %v\n", expected, result)
	}

}

func TestDebase(t *testing.T) {
	tiny := New()
	expected := 22154543
	if result :=  tiny.debase("toshd"); result != expected {
		t.Fatalf("EncodeUrl: unexpected output:\nExpected = %v\nResult = %v\n", expected, result)
	}

}

func TestGetMask(t *testing.T) {
	var in, out = blockSize, 16777215
	if x := getMask(in); x != out {
		t.Errorf("getMask(%v) = %v, want %v", in, x, out)
	}
}

func TestReverseString(t *testing.T) {
	const in, out = "1337", "7331"
	if x := reverseString(in); x != out {
		t.Errorf("reverseString(%v) = %v, want %v", in, x, out)
	}
}

