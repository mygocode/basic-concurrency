# Concurrency with Go

The purpose of this simple program is to implement he concurrency code with Go. `GoRoutines`, `WaitGroups` and `Channels` are the tools Go provides for concurrency.   
This program is creating some Jobs and those jobs are performed by some Workers. Users will provide the workers and jobs data to the program.   
First we will start the workers concurrently and they will wait to receive data from the channel. Please check the worker().

``` go
	// Start the workers
	for i := 0; i < noOfWorkers; i++ {
		go worker(i, &wg, jobChannel, jobResultChannel)
	}
```

Then we will add the jobs to the channel. 

```go
	// Send jobs to worker
	for _, job := range jobs {
		jobChannel <- job
	}
```
As soon as the `jobChannel` start receiving jobs, the workers will fetch the job and start the work.   
After that we will close the channels. Here we are waiting for a while to make sure that workers have completed their work and added result to the result channel
```go 
	close(jobChannel)
	wg.Wait()
	close(jobResultChannel)
  ```
  
Finally, we will read the results from JobResults channel and this is `synchronous`.
```go
 	var jobResults []JobResult
	// Receive job results from workers
	for result := range jobResultChannel {
		jobResults = append(jobResults, result)
	}
```   

Please check the code for better understanding. Happy GO Coding :tulip: 
