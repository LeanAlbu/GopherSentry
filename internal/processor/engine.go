package processor

import(
	"strings"
	"sync"
)

type Job struct{
	ID int
	Line string
}

type Result struct{
	JobID int
	IsThreat bool
	Message string
}


func Worker(id int, jobs <-chan Job, results chan<- Result, wg *sync.WaitGroup){
	defer wg.Done()

	for job := range jobs{
		isThreat := false
		message := "Log normal"

		lineUpper := strings.ToUpper(job.Line)
		if strings.Contains(lineUpper, "DROP TABLE") || strings.Contains(lineUpper, "OR 1=1"){
			isThreat = true
			message = "SQL INJECTION PATTERN DETECTED"
		} else if strings.Contains(lineUpper, "FAILED LOGIN") && strings.Contains(lineUpper, "ROOT"){
			isThreat = true
			message = "Brute force attempt detected on root user"
		}

		results <- Result{
			JobID: job.ID,
			IsThreat: isThreat,
			Message: message,
		}
	}
}
