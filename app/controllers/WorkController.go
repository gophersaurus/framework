package controllers

import (
	gf "git.target.com/gophersaurus/gophersaurus/vendor/git.target.com/gophersaurus/framework"
)

var Work = &workController{}

type workController struct {
}

func (w *workController) Index(resp *gf.Response, req *gf.Request) {
	sum := 0

	jobs := []gf.Job{}
	for i := 0; i < 10; i++ {
		jobs = append(jobs, &SampleJob{i, &sum})
	}

	err := gf.WorkerPool(5, jobs...)
	if err != nil {
		resp.RespondWithErr(err)
		return
	}

	resp.Body(sum)
	resp.Respond()
}

type SampleJob struct {
	value int
	sum   *int
}

func (s *SampleJob) Do() (gf.Result, error) {
	return &SampleResult{s.value * 2, s.sum}, nil
}

type SampleResult struct {
	value int
	sum   *int
}

func (s *SampleResult) Process() error {
	*s.sum = (*s.sum) + s.value
	return nil
}
