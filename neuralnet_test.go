package gomlet

import "testing"

func TestNetLoad(t *testing.T) {
	net := NewNeuralNet(2)
	net.AddLayer(3)
	net.AddLayer(1)

	err := net.Save("test.json")
	if err != nil {
		t.Error(err)
	}
}
