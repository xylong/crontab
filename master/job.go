package master

import (
	"github.com/coreos/etcd/clientv3"
	"time"
)

var (
	jobManager JobManager
)

// JobManager 任务管理器
type JobManager struct {
	client *clientv3.Client
	kv     clientv3.KV
	lease  clientv3.Lease
}

// InitJobManager 初始化管理器
func InitJobManager() (err error) {
	client, err := clientv3.New(clientv3.Config{
		Endpoints:   Conf.Endpoints,
		DialTimeout: time.Duration(Conf.DialTimeout) * time.Millisecond,
	})
	if err != nil {
		return
	}
	jobManager.client = client
	jobManager.kv = clientv3.NewKV(client)
	jobManager.lease = clientv3.NewLease(client)
	return
}
