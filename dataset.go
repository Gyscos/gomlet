package gomlet

type DataSet struct {
	Inputs []Input
	Labels []int
}

func (d *DataSet) Add(i Input, l int) {
	d.Inputs = append(d.Inputs, i)
	d.Labels = append(d.Labels, l)
}
