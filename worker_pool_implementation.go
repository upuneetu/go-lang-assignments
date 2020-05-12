package main

import(
	"fmt"
	"sync"
	"math/rand"
	"time"
)

type Job struct{
	value int
	id int
}

type Result struct{
	result int
	job Job
}

func sum_digits(number int) int{
	res := 0
	for number>0{
		res+=number%10
		number/=10
	}
	time.Sleep(2*time.Second)
	return res
}

func create_jobs(jobs chan Job, nJobs int){
	for i:=0;i<nJobs;i++{
		var jb = Job{i,rand.Intn(999999)}
		jobs <- jb
	}
	close(jobs)
}

func worker(jobs chan Job, results chan Result, wg *sync.WaitGroup){
	for jb := range jobs{
		output := Result{sum_digits(jb.value), jb}
		results <- output
	}
	wg.Done()
}

func create_worker_pool(jobs chan Job, results chan Result, nWorkers int){
	var wg sync.WaitGroup
	for i:=0;i<nWorkers;i++{
		wg.Add(1)
		go worker(jobs,results,&wg);
	}
	wg.Wait();
	close(results)
}

func get_results(results chan Result, done chan bool){
	for result := range results{
		fmt.Println("Job id:", result.job.id, " number:", result.job.value, " sum:", result.result)
	}
	done <- true
}

func main() {
	startTime := time.Now()
	var jobs = make(chan Job, 10)
	var results = make(chan Result, 10)

	var nWorkers = 20
	go create_worker_pool(jobs, results, nWorkers)

	var nJobs = 100
	go create_jobs(jobs, nJobs)

	done := make(chan bool)

	go get_results(results, done)


	<-done
	endTime := time.Now()
	difference := endTime.Sub(startTime)
	fmt.Println("\n\nTotal Time taken:",difference.Seconds())
}
