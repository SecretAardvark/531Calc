package lift

type workout map[string][]float32
type cycle [3]workout

type Lift struct {
	Name        string
	OneRepMax   float32
	TrainingMax float32
	Cycle       cycle
}

func (l *Lift) GetOneRep(weight, reps float32) {
	l.OneRepMax = weight*reps*float32(.0333) + weight

}

func (l *Lift) GetTM() {
	l.TrainingMax = l.OneRepMax * .9
	//return l.TrainingMax
}

func (l *Lift) GetCycle() {
	l.Cycle = cycle{
		workout{"week1": []float32{l.TrainingMax * .65, l.TrainingMax * .75, l.TrainingMax * .85}},
		workout{"week2": []float32{l.TrainingMax * .70, l.TrainingMax * .80, l.TrainingMax * .9}},
		workout{"week3": []float32{l.TrainingMax * .75, l.TrainingMax * .85, l.TrainingMax * .95}},
	}
}
