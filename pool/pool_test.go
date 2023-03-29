package pool

import (
	"fmt"
	"github.com/AgentNemo00/myhttp/checksum"
	"reflect"
	"testing"
)

func TestPool_Do(t *testing.T) {
	type fields struct {
		amount  int
		workers []Worker
		errors  map[string]error
	}
	tests := []struct {
		name   string
		fields fields
		want   map[string]checksum.Checksum
	}{
		{
			name: "empty",
			fields: fields{
				amount: 2,
				workers: []Worker{func() result {
					return result{
						Name:     "one",
						Checksum: "",
						Error:    fmt.Errorf("some error"),
					}
				}},
				errors: map[string]error{},
			},
			want: map[string]checksum.Checksum{},
		},
		{
			name: "one correct",
			fields: fields{
				amount: 2,
				workers: []Worker{
					func() result {
						return result{
							Name:     "one",
							Checksum: "",
							Error:    fmt.Errorf("some error"),
						}
					}, func() result {
						return result{
							Name:     "two",
							Checksum: "two-checksum",
							Error:    nil,
						}
					}},
				errors: map[string]error{},
			},
			want: map[string]checksum.Checksum{"two": checksum.Checksum("two-checksum")},
		},
		{
			name: "all correct",
			fields: fields{
				amount: 2,
				workers: []Worker{
					func() result {
						return result{
							Name:     "one",
							Checksum: "one-checksum",
							Error:    nil,
						}
					}, func() result {
						return result{
							Name:     "two",
							Checksum: "two-checksum",
							Error:    nil,
						}
					}},
				errors: map[string]error{},
			},
			want: map[string]checksum.Checksum{"one": checksum.Checksum("one-checksum"), "two": checksum.Checksum("two-checksum")},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Pool{
				amount:  tt.fields.amount,
				workers: tt.fields.workers,
				errors:  tt.fields.errors,
			}
			if got := p.Do(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Do() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPool_worker(t *testing.T) {
	type fields struct {
		amount  int
		workers []Worker
		errors  map[string]error
	}
	type args struct {
		jobs    chan Worker
		results chan result
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		trigger func(chan Worker)
		want    result
	}{
		{
			name:   "jobs",
			fields: fields{},
			args: args{
				jobs:    make(chan Worker),
				results: make(chan result),
			},
			trigger: func(workers chan Worker) {
				workers <- func() result {
					return result{
						Name:     "result",
						Checksum: "check",
						Error:    nil,
					}
				}
			},
			want: result{
				Name:     "result",
				Checksum: "check",
				Error:    nil,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Pool{
				amount:  tt.fields.amount,
				workers: tt.fields.workers,
				errors:  tt.fields.errors,
			}
			go p.worker(tt.args.jobs, tt.args.results)
			tt.trigger(tt.args.jobs)
			workerResult := <-tt.args.results
			if !reflect.DeepEqual(workerResult, tt.want) {
				t.Errorf("Error: got: %v, want: %v", workerResult, tt.want)
			}
		})
	}
}

func TestNewPool(t *testing.T) {
	type args struct {
		parallel int
	}
	tests := []struct {
		name string
		args args
		want *Pool
	}{
		{
			name: "new pool",
			args: args{parallel: 10},
			want: &Pool{
				amount:  10,
				workers: make([]Worker, 0),
				errors:  make(map[string]error, 0),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewPool(tt.args.parallel); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewPool() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPool_AddWorker(t *testing.T) {
	type fields struct {
		amount  int
		workers []Worker
		errors  map[string]error
	}
	type args struct {
		worker Worker
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "add worker",
			fields: fields{
				amount:  0,
				workers: make([]Worker, 0),
				errors:  nil,
			},
			args: args{func() result {
				return result{
					Name:     "one",
					Checksum: "one-checksum",
					Error:    nil,
				}
			}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Pool{
				amount:  tt.fields.amount,
				workers: tt.fields.workers,
				errors:  tt.fields.errors,
			}
			p.AddWorker(tt.args.worker)
			if !reflect.DeepEqual(p.workers[0](), tt.args.worker()) {
				t.Errorf("Error: got: %v, want: %v", p.workers[0](), tt.args.worker())
			}
		})
	}
}

func TestPool_Errors(t *testing.T) {
	type fields struct {
		amount  int
		workers []Worker
		errors  map[string]error
	}
	tests := []struct {
		name   string
		fields fields
		want   map[string]error
	}{
		{
			name: "error",
			fields: fields{
				errors: make(map[string]error),
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Pool{
				amount:  tt.fields.amount,
				workers: tt.fields.workers,
				errors:  tt.fields.errors,
			}
			if got := p.Errors(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Errors() = %v, want %v", got, tt.want)
			}
		})
	}
}
