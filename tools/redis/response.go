package redis

import (
	"github.com/gomodule/redigo/redis"
)

type Response struct {
	reply interface{}
	error error
}

func NewResponse(reply interface{}, err error) *Response {
	return &Response{reply: reply, error: err}
}

func (r *Response) Result() (interface{}, error) {
	return r.reply, r.error
}

func (r *Response) Error() error {
	return r.error
}

func (r *Response) Hit() bool {
	if r.error != nil || r.reply == nil {
		return false
	}
	return true
}

func (r *Response) Int64() (int64, error) {
	return redis.Int64(r.reply, r.error)
}

func (r *Response) Bool() (bool, error) {
	return redis.Bool(r.reply, r.error)
}

func (r *Response) Bytes() ([]byte, error) {
	return redis.Bytes(r.reply, r.error)
}

func (r *Response) String() (string, error) {
	return redis.String(r.reply, r.error)
}
