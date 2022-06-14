package resource

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"

	"config-server/internal/pkg/etcd3"
)

type Resource struct {
	Name      string      `json:"name"`
	Kind      string      `json:"kind"`
	Namespace string      `json:"namespace"`
	Spec      interface{} `json:"spec"`
}

type ResourceList struct {
	items []Resource
}

func (r *Resource) Create(namespace, kind string, body io.Reader) error {
	if err := json.NewDecoder(body).Decode(r); err != nil {
		return err
	}
	if err := r.validateParams(namespace, kind); err != nil {
		return err
	}
	isExist, err := r.isExist()
	if err != nil {
		return err
	}
	if isExist {
		return errors.New("resource is already exist")
	}
	if err := r.save(); err != nil {
		return err
	}
	return nil
}

func (r *Resource) Get(namespace, kind, name string) error {
	r.Namespace, r.Kind, r.Name = namespace, kind, name
	return r.sync()
}

func (r *Resource) Update(namespace, kind, name string, body io.Reader) error {
	if err := json.NewDecoder(body).Decode(r); err != nil {
		return err
	}
	if err := r.validateParams(namespace, kind, name); err != nil {
		return err
	}
	isExist, err := r.isExist()
	if err != nil {
		return err
	}
	if !isExist {
		return errors.New("resource is not exist")
	}
	if err := r.save(); err != nil {
		return err
	}
	return nil
}

func (r *Resource) Delete(namespace, kind, name string) error {
	r.Namespace, r.Kind, r.Name = namespace, kind, name
	return r.delete()
}

func (rl *ResourceList) Get() {

}

func (r *Resource) validateParams(params ...string) error {
	if len(params) >= 1 && params[0] != r.Namespace {
		return errors.New("namespace is not match")
	}
	if len(params) >= 2 && params[1] != r.Kind {
		return errors.New("kind is not match")
	}
	if len(params) >= 3 && params[2] != r.Name {
		return errors.New("name is not match")
	}
	return nil
}

func (r *Resource) isExist() (bool, error) {
	key := fmt.Sprintf("/%s/%s/%s", r.Namespace, r.Kind, r.Name)
	return etcd3.IsExist(key)
}

func (r *Resource) save() error {
	key := fmt.Sprintf("/%s/%s/%s", r.Namespace, r.Kind, r.Name)
	value, err := json.Marshal(r)
	if err != nil {
		return err
	}
	return etcd3.SaveKV(key, string(value))
}

func (r *Resource) sync() error {
	key := fmt.Sprintf("/%s/%s/%s", r.Namespace, r.Kind, r.Name)
	value, err := etcd3.GetKV(key)
	if err != nil {
		return err
	}
	return json.Unmarshal(value, r)
}

func (r *Resource) delete() error {
	key := fmt.Sprintf("/%s/%s/%s", r.Namespace, r.Kind, r.Name)
	return etcd3.DelKV(key)
}
