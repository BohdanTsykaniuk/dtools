package main

const (
	Started = iota
	Failed
	Done
)

type IsVerifiable interface {
	IsConflict(IsVerifiable) bool
}
type Job struct {
	JobId string
	Data  interface{}
}

type FailedJob struct {
	JobId       string
	ErrorData   interface{}
	DataToError interface{}
}

type CompletedJob struct {
	JobId         string
	ResultData    interface{}
	DataToSuccess interface{}
}

type TerminateDispatchJob struct {
}
