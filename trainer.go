package gomlet

type Trainer interface {
	Train(DataSet) Classifier
}
