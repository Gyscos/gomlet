package gomlet

import (
	"encoding/json"
	"os"
)

type NeuralNet struct {
	InputSize   int
	LayerSizes  []int
	Transitions []Matrix
}

func NewNeuralNet(inputSize int) *NeuralNet {
	return &NeuralNet{InputSize: inputSize}
}

func (n *NeuralNet) Evaluate(input []float64) []float64 {
	for _, l := range n.Transitions {
		input = l.Apply(input)
	}

	return input
}

func (n *NeuralNet) AddLayer(size int) {

	prevSize := n.InputSize
	if len(n.LayerSizes) > 0 {
		prevSize = n.LayerSizes[len(n.LayerSizes)-1]
	}

	n.LayerSizes = append(n.LayerSizes, size)
	n.Transitions = append(n.Transitions, NewMatrix(size, prevSize))
}

func (n *NeuralNet) Save(filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	err = json.NewEncoder(file).Encode(n)
	if err != nil {
		return err
	}

	return nil
}

func (n *NeuralNet) Load(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return nil
	}
	defer file.Close()

	err = json.NewDecoder(file).Decode(n)
	if err != nil {
		return err
	}

	return err
}
