package pool

import (
	"github.com/AgentNemo00/myhttp/checksum"
)

// Pool - worker pool
// amount of parallel processes
// workers to use for the jobs
// errors occurred during the process
type Pool struct {
	amount  int
	workers []Worker
	errors  map[string]error
}

func NewPool(parallel int) *Pool {
	return &Pool{
		amount:  parallel,
		workers: make([]Worker, 0),
		errors:  make(map[string]error),
	}
}

func (p *Pool) AddWorker(worker Worker) {
	p.workers = append(p.workers, worker)
}

// Do - triggers the work process
func (p *Pool) Do() map[string]checksum.Checksum {
	results := make(chan result, len(p.workers))
	jobs := make(chan Worker, len(p.workers))
	// initiate parallel working process
	for i := 0; i < p.amount; i++ {
		go p.worker(jobs, results)
	}
	// send jobs
	for _, worker := range p.workers {
		jobs <- worker
	}
	close(jobs)
	// fetch results
	ret := make(map[string]checksum.Checksum)
	for i := 0; i < len(p.workers); i++ {
		workerResult := <-results
		if workerResult.Error != nil {
			p.errors[workerResult.Name] = workerResult.Error
			continue
		}
		ret[workerResult.Name] = workerResult.Checksum
	}
	return ret
}

func (p *Pool) worker(jobs <-chan Worker, results chan<- result) {
	for job := range jobs {
		results <- job()
	}
}

func (p *Pool) Errors() map[string]error {
	return p.errors
}
