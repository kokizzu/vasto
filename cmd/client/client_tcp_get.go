package client

import (
	"errors"
	"fmt"

	"github.com/chrislusf/vasto/pb"
	"github.com/chrislusf/vasto/topology"
)

var (
	NotFoundError = errors.New("NotFound")
)

func (c *VastoClient) Get(keyspace string, key []byte, options ...topology.AccessOption) ([]byte, error) {

	conn, replica, err := c.ClusterListener.GetConnectionByPartitionKey(keyspace, key, options...)
	if err != nil {
		return nil, err
	}

	request := &pb.Request{
		Get: &pb.GetRequest{
			Replica: uint32(replica),
			Key:     key,
		},
	}

	requests := &pb.Requests{}
	requests.Requests = append(requests.Requests, request)

	responses, err := pb.SendRequests(conn, requests)
	conn.Close()
	if err != nil {
		return nil, fmt.Errorf("get error: %v", err)
	}

	if len(responses.Responses) == 0 {
		return nil, NotFoundError
	}

	response := responses.Responses[0]

	if response.Get.Status != "" {
		return nil, fmt.Errorf(response.Get.Status)
	}

	kv := response.Get.KeyValue
	if kv == nil {
		return nil, NotFoundError
	}

	return kv.Value, nil
}