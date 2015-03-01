package controllers

import (
	"fmt"

	"git.target.com/gophersaurus/gf.v1"
)

var Work = &workController{}

type workController struct {
}

func (w *workController) Index(resp gf.Responder, req gf.Requester) {
	sum := 0

	jobs := []gf.Job{}
	for i := 0; i < 10; i++ {
		jobs = append(jobs, &SampleJob{i, &sum})
	}

	err := gf.NewWorkerPool(5, jobs...).RunJobs()
	if err != nil {
		resp.WriteErrs(err.Error())
		return
	}

	resp.Read(sum)
	resp.JSON()
}

func double(i int) int {
	return i * 2
}

type SampleJob struct {
	value int
	sum   *int
}

func (s *SampleJob) DoWork() {
	if s.value == 7 {
		panic("no sevens")
	}
	s.value = double(s.value)
	fmt.Printf("value: %v\n", s.value)
}

func (s *SampleJob) ProcessResult() {
	*s.sum = (*s.sum) + s.value
	fmt.Printf("sum: %v\n", *s.sum)
}
