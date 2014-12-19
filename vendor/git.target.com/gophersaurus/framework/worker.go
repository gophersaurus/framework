package gophersauras

import "sync"

type Job interface {
	Do() (Result, error)
}

type Result interface {
	Process() error
}

func WorkerPool(size int, jobs ...Job) error {
	jobCount := len(jobs)
	jobChan := make(chan *Job, jobCount)
	resultChan := make(chan *Result, jobCount)

	var err error

	wg := &sync.WaitGroup{}

	for w := 0; w < size; w++ {
		go worker(jobChan, resultChan, wg, &err)
	}

	for _, job := range jobs {
		jobChan <- job
	}

	wg.Wait()
	if err != nil {
		return err
	}
	close(jobChan)

	for r := 0; r < jobCount; r++ {
		result := <-resultChan
		err := result.Process()
		if err != nil {
			return err
		}
	}
	return nil
}

func worker(jobs <-chan *Job, results chan<- *Result, wg *sync.WaitGroup, err *error) {
	for j := range jobs {
		if err != nil {
			wg.Done()
			return
		}
		result, er := j.Do()
		if er != nil {
			*err = er
			wg.Done()
			return
		}
		results <- j.Do()
	}
	wg.Done()
}
