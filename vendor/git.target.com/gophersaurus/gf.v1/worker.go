package gf

type Job interface {
	DoWork()
	ProcessResult()
}

type workerPool struct {
	Size int
	Jobs []Job
}

func NewWorkerPool(size int, jobs ...Job) *workerPool {
	return &workerPool{size, jobs}
}

func (w *workerPool) RunJobs() error {
	var jobErr error
	w.runJobs(&jobErr)
	return jobErr
}

func (w *workerPool) runJobs(err *error) {
	defer Relieve(err)
	jobCount := len(w.Jobs)
	jobChan := make(chan Job, jobCount)
	resultChan := make(chan Job, jobCount)

	var workErr error
	for i := 0; i < w.Size; i++ {
		go newWorker().work(jobChan, resultChan, &workErr)
	}

	for _, j := range w.Jobs {
		if workErr != nil {
			*err = workErr
			return
		}
		jobChan <- j
	}
	close(jobChan)
	if workErr != nil {
		*err = workErr
		return
	}

	for r := 0; r < jobCount; r++ {
		result, isOpen := <-resultChan
		if !isOpen {
			if workErr != nil {
				*err = workErr
			}
			return
		}
		result.ProcessResult()
	}
}

type worker struct {
}

func newWorker() *worker {
	return &worker{}
}

func (w *worker) work(jobs <-chan Job, results chan<- Job, workErr *error) {
	for j := range jobs {
		if *workErr != nil {
			return
		}
		w.workSafe(j, workErr)
		if *workErr != nil {
			close(results)
			return
		}
		results <- j
	}
}

func (w *worker) workSafe(job Job, err *error) {
	defer Relieve(err)
	job.DoWork()
}
