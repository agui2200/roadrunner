package http

import (
	"github.com/agui2200/roadrunner/util"
	"github.com/pkg/errors"
)

type rpcServer struct{ svc *Service }

// WorkerList contains list of workers.
type WorkerList struct {
	// Workers is list of workers.
	Workers []*util.State `json:"workers"`
}

// Reset resets underlying RR worker pool and restarts all of it's workers.
func (rpc *rpcServer) Reset(reset bool, r *string) error {
	if rpc.svc == nil || rpc.svc.handler == nil {
		return errors.New("http server is not running")
	}

	*r = "OK"
	return rpc.svc.rr.Reset()
}

// Workers returns list of active workers and their stats.
func (rpc *rpcServer) Workers(list bool, r *WorkerList) (err error) {
	if rpc.svc == nil || rpc.svc.handler == nil {
		return errors.New("http server is not running")
	}

	r.Workers, err = util.ServerState(rpc.svc.rr)
	return err
}
