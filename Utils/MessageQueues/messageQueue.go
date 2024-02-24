package MessageQueues

/*
-> Creating worker pool mainly involves below steps/operations to be performed
1. Creating pool of workers listening to jobs channel waiting for a task to be assigned.
2. jobs are being added to job channel
3. Once worker completes the task write the result into result channel
4. Render / print result of each task.
*/

import (
	"TigerPopulation/Domain/Tigers"
	"fmt"
	"sync"
	"time"

	"github.com/golang/glog"
)

type Job struct {
	id   int
	User Tigers.EmailStruct
}
type Result struct {
	job    Job
	status bool // email sent or not
}

var jobs = make(chan Job, 100)
var results = make(chan Result, 100)
var numOfWorkers int = 10

// func to calculate number sum for each digit
func SendEmail(emailInfo Tigers.EmailStruct) bool {
	// send email here
	err := SendMail(emailInfo)
	if err != nil {
		return false
	}
	return true
}

// creating a worker
// each worker keeps on waiting for the task to be assigned once one task is completed untill no task in jobs chennal
func Worker(wg *sync.WaitGroup) {
	for val := range jobs { // receiving task from buffered channel
		resp := Result{val, SendEmail(val.User)}
		// assign to results channel to be read to render the data
		results <- resp
	}
	wg.Done()
}

// create  worker pool
func CreateWorkerPool(n int) {
	var wg sync.WaitGroup
	for i := 0; i < n; i++ {
		wg.Add(1)
		go Worker(&wg)
	}
	wg.Wait()
	// close the result channel
	close(results)
}

// allocate/add jobs to job channel
func AddJobs(users []Tigers.EmailStruct) {
	// add jobs to  jobs channel
	for ind, val := range users {
		job := Job{ind, val}
		jobs <- job
	}
	close(jobs) // once all jobs are assigned close the channel
}

// render the result from result channel

func AnalyseDeliveryReport(done chan bool) {
	// read from result channel and print the result
	for val := range results {
		glog.V(2).Infoln("Result = ", val)
	}
	done <- true
}

func SendEmailToUsers(tigerId int) error {
	start := time.Now()
	ch := make(chan bool)
	//var numOfJobs int = 100
	// fetch users to send email
	users, err := Tigers.GetUsersByTigerId(tigerId)
	if err != nil {
		glog.Errorln("Error getting users : Err", err)
		return err
	}
	// allocate some jobs first
	go AddJobs(users)
	// run worker pool which will create worker inside the go routine
	go AnalyseDeliveryReport(ch)
	go CreateWorkerPool(numOfWorkers)
	<-ch
	fmt.Println("Time taken by worker pool to complete all the tasks (Seconds) = ", time.Now().Sub(start).Seconds())
	return nil
}
