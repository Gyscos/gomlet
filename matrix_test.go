package gomlet

import (
	"encoding/json"
	"log"
	"testing"
)

func TestMatrix(t *testing.T) {
	M := NewMatrix(2, 1)
	N := NewMatrix(1, 2)

	M.Set(0, 0, 1)
	M.Set(1, 0, 2)
	N.Set(0, 0, 1)
	N.Set(0, 1, 2)

	X, err := MatTimes(M, N)
	if err != nil {
		t.Error(err)
	}
	log.Println(X)
}

func TestLoad(t *testing.T) {
	M := NewMatrix(2, 1)
	N := NewMatrix(1, 2)

	M.Set(0, 0, 1)
	M.Set(1, 0, 2)
	N.Set(0, 0, 1)
	N.Set(0, 1, 2)

	b, err := json.Marshal(M)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(string(b))
	b, err = json.Marshal(N)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(string(b))
}
