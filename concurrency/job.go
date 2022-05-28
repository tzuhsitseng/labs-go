package concurrency

import (
	"errors"
	"fmt"
	"math/rand"
)

type JobType int

type Job interface {
	GetID() int
	GetType() JobType
}

type AddJob struct {
	_     struct{}
	Left  int
	Right int
}

func (j AddJob) GetID() int {
	return rand.Int()
}

func (j AddJob) GetType() JobType {
	return Add
}

type SubtractJob struct {
	_     struct{}
	Left  int
	Right int
}

func (j SubtractJob) GetType() JobType {
	return Subtract
}

func (j SubtractJob) GetID() int {
	return rand.Int()
}

type Executor interface {
	DoJob(job Job) error
}

type AddExecutor struct {
}

func (e AddExecutor) DoJob(job Job) error {
	addJob, ok := job.(AddJob)
	if !ok {
		return errors.New("not add job")
	}

	fmt.Println("do add job:", addJob.Left+addJob.Right)
	return nil
}

type SubtractExecutor struct {
}

func (e SubtractExecutor) DoJob(job Job) error {
	subJob, ok := job.(SubtractJob)
	if !ok {
		return errors.New("not subtract job")
	}

	fmt.Println("do subtract job:", subJob.Left-subJob.Right)
	return nil
}

func NewExecutor(job Job) (Executor, error) {
	switch job.GetType() {
	case Add:
		return &AddExecutor{}, nil
	case Subtract:
		return &SubtractExecutor{}, nil
	}
	return nil, errors.New("no matched executor")
}

const (
	Add JobType = iota + 1
	Subtract
)
