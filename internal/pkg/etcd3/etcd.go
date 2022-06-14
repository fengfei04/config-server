package etcd3

import (
	"context"
	"fmt"
	"time"

	clientv3 "go.etcd.io/etcd/client/v3"
)

var (
	config clientv3.Config
	client *clientv3.Client
	err    error
)

func init() {
	config = clientv3.Config{
		Endpoints:   []string{"0.0.0.0:2379"},
		DialTimeout: 5 * time.Second,
	}
	if client, err = clientv3.New(config); err != nil {
		panic(err.Error())
	}
}

func IsExist(key string) (bool, error) {
	kv := clientv3.NewKV(client)
	resp, err := kv.Get(context.TODO(), key, clientv3.WithCountOnly())
	if err != nil {
		return false, err
	}
	if resp.Count == 0 {
		return false, nil
	}
	return true, nil
}

func SaveKV(key, value string) error {
	kv := clientv3.NewKV(client)
	if _, err := kv.Put(context.TODO(), key, value, clientv3.WithPrevKV()); err != nil {
		return err
	}
	return nil
}

func GetKV(key string) ([]byte, error) {
	kv := clientv3.NewKV(client)
	resp, err := kv.Get(context.TODO(), key)
	if err != nil {
		return nil, err
	}
	if len(resp.Kvs) == 0 {
		return nil, fmt.Errorf("key[%s] is not exist", key)
	}
	return resp.Kvs[0].Value, nil
}

func DelKV(key string) error {
	kv := clientv3.NewKV(client)
	if _, err := kv.Delete(context.TODO(), key, clientv3.WithPrevKV()); err != nil {
		return err
	}
	return nil
}
