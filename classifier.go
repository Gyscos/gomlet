package gomlet

type Classifier interface {
	Classify(Input) int
}

func Evaluate(c Classifier, d DataSet) float32 {
	fails := 0
	for i, _ := range d.Inputs {
		if c.Classify(d.Inputs[i]) != d.Labels[i] {
			fails++
		}
	}
	return float32(len(d.Inputs)-fails) / float32(len(d.Inputs))
}

type DumbClassifier struct{}

func (d *DumbClassifier) Classify(i Input) int {
	return 0
}
