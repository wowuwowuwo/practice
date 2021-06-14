package code_engine

import (
	"fmt"
	"sort"
)

// author by chenxi
// merge multi sorted arrays, it seems like one simple map-reduce procedure,

// Node assume element type is int
type Node struct {
	val  int
	task *MergeTask
}

type MergeTask struct {
	taskId     int
	rawData    []int
	curPos     int
	sendSignal chan struct{}
	isDone     bool
}

func (task *MergeTask) SendOne() *Node {
	if task.curPos < len(task.rawData) {
		val := task.rawData[task.curPos]
		task.curPos++
		if task.curPos == len(task.rawData) {
			task.isDone = true
		}
		// change raw format to middle format
		return &Node{val: val, task: task}
	}
	return nil
}

type MergeJob struct {
	tasks       []*MergeTask
	doneTaskNum int

	isDone bool

	nodeNum int
	output  []int
}

func NewMergeJob(input [][]int) *MergeJob {
	job := &MergeJob{
		tasks:       make([]*MergeTask, 0, len(input)),
		doneTaskNum: 0,
		isDone:      false,
		nodeNum:     0,
		output:      nil,
	}
	for index, data := range input {
		// skip empty array
		if len(data) == 0 {
			continue
		}
		job.nodeNum += len(data)

		// init one task
		sig := make(chan struct{}, 1)
		task := &MergeTask{
			taskId:     index,
			rawData:    data,
			curPos:     0,
			sendSignal: sig,
			isDone:     false,
		}

		job.tasks = append(job.tasks, task)
	}
	job.output = make([]int, 0, job.nodeNum)
	return job
}

// DoSimpleMerge one simple way to merge input arrays
func (job *MergeJob) DoSimpleMerge() {
	job.runMerge()
}

func (job *MergeJob) runMerge() {
	if len(job.tasks) == 0 || job.nodeNum == 0 {
		job.isDone = true
		return
	}

	// init out heap, test shows that simple heap is faster than std heap about 50%
	//outHeap := &OutHeapStd{data: make([]*Node, 0, len(job.tasks))}
	outHeap := &OutHeapSimple{data: make([]*Node, 0, len(job.tasks))}
	for _, task := range job.tasks {
		// get each task first node
		// todo, it's possible to use batch sending to speed up
		node := task.SendOne()
		outHeap.Add(node)
		//fmt.Printf("heap init: add node: %v\n", node)
	}

	for {
		minNode := outHeap.GetMin()
		if minNode == nil {
			job.isDone = true
			return
		}
		//fmt.Printf("got node: %v\n", minNode)

		// output one min node, change middle format to output format
		job.output = append(job.output, minNode.val)
		if len(job.output) == job.nodeNum {
			job.isDone = true
			//fmt.Printf("output full, job is done: %v\n", minNode)
			return
		}

		// get next node from task of minNode
		if minNode.task.isDone {
			continue
		}
		// simple way
		nextNode := minNode.task.SendOne()
		if nextNode != nil {
			outHeap.Add(nextNode)
			//fmt.Printf("heap running: add node: %v\n", node)
		}
	}
}

func (job *MergeJob) GetOutput() []int {
	// change output format if needed
	job.changeOutputFormat()
	return job.output
}

func (job *MergeJob) changeOutputFormat() {
}

func checkResEqual(ans, expected []int) bool {
	if len(ans) != len(expected) {
		return false
	}
	for i := 0; i < len(ans); i++ {
		if ans[i] != expected[i] {
			return false
		}
	}
	return true
}

// MergeMultiSortedArrays assume input element type is int
func MergeMultiSortedArrays(input [][]int) []int {
	for _, data := range input {
		if !sort.IntsAreSorted(data) {
			// invalid inputs
			fmt.Printf("invalid inputs\n")
			return []int{}
		}
	}
	job := NewMergeJob(input)
	// simple way
	//fmt.Printf("mode simple\n")
	job.DoSimpleMerge()
	output := job.GetOutput()
	return output
}
