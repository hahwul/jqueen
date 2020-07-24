package job


//Job is struct of job
type Job struct{
	ID       string
	PID      int
	CMD      string
	Name     string
	Desc     string
	Priority int
	Flags    []string
}

// JobQueueItem is struct of QueueItem
type JobQueueItem struct {
	item interface{} 
	prev *JobQueueItem 
}

// JobQueue is struct of Job queue
type JobQueue struct {
	current *JobQueueItem 
	last    *JobQueueItem 
	depth   uint64     
}

// New is init JobQueue
func New() *JobQueue {
	var queue *JobQueue = new(JobQueue)

	queue.depth = 0

	return queue
}

// Enqueue is push item to queue
func (queue *JobQueue) Enqueue(item interface{}) {
	if queue.depth == 0 {
		queue.current = &JobQueueItem{item: item, prev: nil}
		queue.last = queue.current
		queue.depth++
		return
	}

	q := &JobQueueItem{item: item, prev: nil}
	queue.last.prev = q
	queue.last = q
	queue.depth++
}

// Dequeue is pop item in queue
func (queue *JobQueue) Dequeue() interface{} {
	if queue.depth > 0 {
		item := queue.current.item
		queue.current = queue.current.prev
		queue.depth--

		return item
	}

	return nil
}

// TestFunction is only testing function
/*
func TestFunction(){
	var queue *JobQueue = New()
	job := new(Job)
	job.id="asdf"
	queue.Enqueue(job)
	queue.Enqueue(2)
	queue.Enqueue(3)
	queue.Enqueue(4)
	queue.Enqueue(5)

	for i := 1; i < 6; i++ {
		item := queue.Dequeue()
		fmt.Println("Dequeue value :", item)
	}

}
*/
