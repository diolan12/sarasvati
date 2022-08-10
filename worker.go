package main

import (
	"fmt"
	"os/exec"

	"github.com/fatih/color"
)

type Worker struct {
	index  int
	id     int
	status string
}

// Here's the worker, of which we'll run several
// concurrent instances. These workers will receive
// work on the `jobs` channel and send the corresponding
// results on `results`. We'll sleep a second per job to
// simulate an expensive task.
func worker(id int, jobs <-chan string, results chan<- string, statusBar *StatusBar, worker Worker) {

	for j := range jobs {
		// color.Green("worker " + strconv.Itoa(id) + " started  job " + strconv.Itoa(j))
		// fmt.Println("worker", id, "started  job", j)
		worker.status = "mapping " + j
		statusBar.Add(0, worker)

		// rand.Seed(time.Now().UnixNano())
		// random := time.Duration(rand.Int31n(100))
		// time.Sleep(time.Millisecond * random)
		cmd := exec.Command("sarasvati", "index", "map", j)
		cmd.Run()
		// if err != nil {
		// log.Fatal(err)
		// }
		// args := []string{"bot", "index", "map", j}
		// index(args)

		worker.status = "idle"
		statusBar.Add(1, worker)
		// time.Sleep(time.Second)
		// color.Green("worker " + strconv.Itoa(id) + " finished  job " + strconv.Itoa(j))
		// fmt.Println("worker", id, "finished job", j)
		results <- j
	}
}
func getAllRegencies() ([]string, []string, []string) {
	jobProvinces := []string{}
	jobRegencies := []string{}
	jobDistricts := []string{}
	rootFolder := outputDir
	helperLoads(rootFolder, &provinces)
	for _, province := range provinces {
		// fmt.Println(province.Name)
		jobProvinces = append(jobProvinces, province.ID)
		provinceFolder := rootFolder + "/" + province.ID + "-" + province.Name
		helperLoads(provinceFolder, &regencies)
		for _, regency := range regencies {
			// fmt.Println("\t|--- ", regency.Name)
			jobRegencies = append(jobRegencies, regency.ID)
			districtFolder := provinceFolder + "/" + regency.ID + "-" + regency.Name
			helperLoads(districtFolder, &districts)
			for _, district := range districts {
				// fmt.Println("\t\t|--- ", district.Name)
				jobDistricts = append(jobDistricts, district.ID)
			}
		}
	}
	return jobProvinces, jobRegencies, jobDistricts
}
func work(args []string) {
	// getAllRegencies()
	provs, regs, dists := getAllRegencies()
	// os.Exit(0)
	// scrappingProvinces()
	// jobList := []string{
	// 	"11", "12", "13", "14", "15", "16", "17", "18", "19", "21", "31", "32", "33", "34", "35", "36", "51", "52", "53", "61", "62", "63", "64", "65", "71", "72", "73", "74", "75", "76", "81", "82", "91", "94",
	// }
	// fmt.Println(provs)

	defer dispatcher(dists)

	defer dispatcher(regs)

	defer dispatcher(provs)

}

func dispatcher(jobList []string) {

	color.Yellow("[ Sarasvati Bot Dispatcher ]")
	// In order to use our pool of workers we need to send
	// them work and collect their results. We make 2
	// channels for this.
	numJobs := len(jobList)
	workersCount := 8
	workersList := []Worker{}
	jobs := make(chan string, numJobs)
	results := make(chan string, numJobs)
	statusBar := NewStatusBar(0, numJobs, workersCount)

	// This starts up 3 workers, initially blocked
	// because there are no jobs yet.
	for w := 0; w < workersCount; w++ {
		work := Worker{index: w, id: w, status: "idle"}
		workersList = append(workersList, work)
		go worker(w, jobs, results, statusBar, work)
	}
	statusBar.AddWorkers(workersList)
	go statusBar.Run()
	for w := 0; w < workersCount; w++ {
		fmt.Print("\n")
	}
	fmt.Print("\n")
	// go bl(statusBar)
	// Here we send 5 `jobs` and then `close` that
	// channel to indicate that's all the work we have.
	for j := 0; j < numJobs; j++ {
		// statusBar.Refresh()
		jobs <- jobList[j]
		// time.Sleep(time.Millisecond * 10)
		// fmt.Println("sent job", j)
	}
	defer close(jobs)

	// Finally we collect all the results of the work.
	// This also ensures that the worker goroutines have
	// finished. An alternative way to wait for multiple
	// goroutines is to use a [WaitGroup](waitgroups).
	for a := 1; a <= numJobs; a++ {
		<-results
		// statusBar.Refresh()
		// println("received result", a)
	}
	fmt.Println("\n\n\n\njobs closed")
}

// ###############################################################################

func bufWrite() {
	color.Green("Scrapping Provinces")
	startBar()
	startStatusBar()
}
