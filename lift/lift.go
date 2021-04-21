package lift

type workout map[string][]float32
type cycle [3]workout

//Lift represents one of the 5/3/1 main lifts.
type Lift struct {
	Name        string
	OneRepMax   float32
	TrainingMax float32
	Cycle       cycle
}

//GetOneRep calculates estimated 1rm based on a given set.
func (l *Lift) GetOneRep(weight, reps float32) {
	l.OneRepMax = weight*reps*float32(.0333) + weight

}

//GetTM calculates the 5/3/1 training max based on the repmax.
func (l *Lift) GetTM() {
	l.TrainingMax = l.OneRepMax * .9

}

//GetCycle calculates the training cycle based on the training max weight.
func (l *Lift) GetCycle() {
	l.Cycle = cycle{
		workout{"week1": []float32{l.TrainingMax * .65, l.TrainingMax * .75, l.TrainingMax * .85}},
		workout{"week2": []float32{l.TrainingMax * .70, l.TrainingMax * .80, l.TrainingMax * .9}},
		workout{"week3": []float32{l.TrainingMax * .75, l.TrainingMax * .85, l.TrainingMax * .95}},
	}
}
