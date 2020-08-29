package main

import (
	"errors"
	"fmt"
	"math"
	"strconv"
	"sync"
	"time"
)

type Job struct {
	Id int
}

type JobResult struct {
	Output string
}

func main() {

	// Taking input from user
	fmt.Println("Workers: ")
	var w string
	fmt.Scanln(&w)

	fmt.Println("Total jobs to perform: ")
	var j string
	fmt.Scanln(&j)

	totalWorkers, _ := strconv.Atoi(w)
	totalJobs, _ := strconv.Atoi(j)

	concurrentFunc(totalWorkers, totalJobs)
}

func concurrentFunc(noOfWorkers int, noOfjobs int) (result int) {

	//Negative check
	if (math.Signbit(float64(noOfWorkers)) == true) || (math.Signbit(float64(noOfjobs)) == true) {
		errors.New("Negative numbers are not allowed")
		return -1
	}

	start := time.Now()
	var jobs []Job
	for i := 0; i < noOfjobs; i++ {
		jobs = append(jobs, Job{Id: i})
	}

	var wg sync.WaitGroup

	wg.Add(noOfWorkers)
	jobChannel := make(chan Job)
	jobResultChannel := make(chan JobResult, len(jobs))

	// Start the workers
	for i := 0; i < noOfWorkers; i++ {
		go worker(i, &wg, jobChannel, jobResultChannel)
	}

	// Send jobs to worker
	for _, job := range jobs {
		jobChannel <- job
	}

	close(jobChannel)
	wg.Wait()
	close(jobResultChannel)

	var jobResults []JobResult
	// Receive job results from workers
	for result := range jobResultChannel {
		jobResults = append(jobResults, result)
	}

	fmt.Printf("Total jobs completed %d\n", len(jobResults))
	fmt.Printf("Took %s\n", time.Since(start))

	return len(jobResults)

}

func worker(id int, wg *sync.WaitGroup, jobChannel <-chan Job, resultChannel chan JobResult) {
	defer wg.Done()
	for job := range jobChannel {
		resultChannel <- startWork(id, job)
	}
}

func startWork(workerId int, job Job) JobResult {
	//fmt.Printf("Worker #%d Running job %d\n", workerId, job.Id)
	time.Sleep(50 * time.Millisecond)
	return JobResult{Output: "Success"}
}
