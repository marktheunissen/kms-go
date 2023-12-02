// Copyright 2023 - MinIO, Inc. All rights reserved.
// Use of this source code is governed by the AGPLv3
// license that can be found in the LICENSE file.

package kms

import (
	"io"
	"net/http"
	"time"

	"github.com/minio/kms-go/kms/internal/headers"
	pb "github.com/minio/kms-go/kms/protobuf"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// readResponse reads the response body into v using the
// response content encoding.
//
// readBesponse assumes that the response body is limited
// to a reasonable size. It returns an error if it cannot
// determine the response content length before decoding.
func readResponse[M any, P pb.Pointer[M], T pb.Unmarshaler[P]](r *http.Response, v T) error {
	if r.ContentLength < 0 {
		return Error{http.StatusLengthRequired, "request content length is negative"}
	}
	// TODO(aead): consider limiting body to a max. content length

	buf := make([]byte, r.ContentLength)
	if _, err := io.ReadFull(r.Body, buf); err != nil {
		return err
	}

	var m M
	var p P = &m
	if r.Header.Get(headers.ContentType) == headers.ContentTypeBinary {
		if err := proto.Unmarshal(buf, p); err != nil {
			return err
		}
	} else {
		if err := protojson.Unmarshal(buf, p); err != nil {
			return err
		}
	}
	return v.UnmarshalPB(p)
}

// NodeStatusResponse contains status information about a single
// KMS cluster node.
type NodeStatusResponse struct {
	// Version is the version of the KMS server. It's the timestamp of
	// the latest commit formatted as 'yyyy-mm-ddThh-mm-ssZ'. For example,
	// "2023-12-01T16-06-52Z"
	Version string

	// APIVersion is the API version supported by the KMS server.
	// For example, "v1".
	APIVersion string

	// Endpoint is the KMS server endpoint as 'host' or 'host:port'.
	Endpoint string

	// Role is the current role the KMS server node has within the cluster.
	// Either, "Leader", "Follower" or "Candidate".
	Role string

	// Commit is the number of state changes applied to this KMS server.
	Commit uint64

	// Nodes is a list of KMS server nodes within the KMS cluster as a map
	// of node IDs to KMS server addresses of the form 'host' or 'host:port'.
	Nodes map[int]string

	// ID is the node ID of this KMS server. It only changes if the node
	// joins a cluster.
	ID int

	// LeaderID is the ID of the current cluster leader or negative if
	// the cluster has no leader.
	LeaderID int

	// LastHeartbeat is the duration since the KMS server has sent or received
	// a heartbeat. As long as there is a cluster leader, it should be lower
	// than the ElectionTimeout.
	LastHeartbeat time.Duration

	// HeartbeatInterval defines the frequency in which this KMS server, as cluster
	// leader, sends heartbeats to its follower nodes. All nodes within a cluster
	// should use the same heartbeat interval.
	HeartbeatInterval time.Duration

	// ElectionTimeout defines how long a KMS server node waits for heartbeats before
	// it considers the cluster leaders as down and starts a leader election to become
	// the cluster leader itself.
	//
	// Each cluster node should have a slightly different election timeout to avoid
	// spliting votes. Typically, base election timeout + random jitter. The average
	// or base election timeout of all cluster nodes should be balanced with the
	// HeartbeatInterval to prevent nodes from starting elections even though a leader
	// is present. A reasonable default may be:
	//
	//   ElectionTimeout = 3 * HeartbeatInterval.
	ElectionTimeout time.Duration

	// UpTime is the amount of time the KMS server is up and running.
	UpTime time.Duration

	// OS identifies the operating system the KMS server is running on.
	// For example, "linux" or "darwin".
	OS string

	// CPUArch is the CPU architecture of the KMS server. For example, "amd64".
	CPUArch string

	// CPUs is the number of logical CPUs that can execite the KMS server process.
	// However, the KMS server may not use all of these CPUs. It might be limited
	// to fewer CPUs.
	CPUs uint

	// UsableCPUs is the number of CPUs actually used by the KMS server process.
	// Unless the KMS server has been limited to fewer CPUs, equal to CPUs field.
	UsableCPUs uint

	// HeapMemInUse is the amount of heap memory currently occupied by the KMS server.
	// The total amount of memory used by the KMS server process is HeapMemInUse +
	// StackMemInUse.
	HeapMemInUse uint64

	// StackMemInUse is the amount of stack memory currently occupied by the KMS server.
	// The total amount of memory used by the KMS server process is HeapMemInUse +
	// StackMemInUse.
	StackMemInUse uint64
}

// MarshalPB converts the NodeStatusResponse into its protobuf representation.
func (s *NodeStatusResponse) MarshalPB(v *pb.NodeStatusResponse) error {
	v.Version = s.Version
	v.APIVersion = s.APIVersion
	v.Addr = s.Endpoint
	v.Role = s.Role
	v.Commit = s.Commit
	v.Nodes = make(map[uint32]string, len(s.Nodes))
	for id, node := range s.Nodes {
		v.Nodes[uint32(id)] = node
	}
	v.ID = uint32(s.ID)
	v.LeaderID = int64(s.LeaderID)
	v.LastHeartbeat = pb.Duration(s.LastHeartbeat)
	v.HeartbeatInterval = pb.Duration(s.HeartbeatInterval)
	v.ElectionTimeout = pb.Duration(s.ElectionTimeout)
	v.OS = s.OS
	v.Arch = s.CPUArch
	v.CPUs = uint32(s.CPUs)
	v.UsableCPUs = uint32(s.UsableCPUs)
	v.HeapMemInUse = s.HeapMemInUse
	v.StackMemInUse = s.StackMemInUse
	return nil
}

// UnmarshalPB initializes the NodeStatusResponse from its protobuf representation.
func (s *NodeStatusResponse) UnmarshalPB(v *pb.NodeStatusResponse) error {
	s.Version = v.Version
	s.APIVersion = v.APIVersion
	s.Endpoint = v.Addr
	s.Role = v.Role
	s.Commit = v.Commit
	s.Nodes = make(map[int]string, len(v.Nodes))
	for id, node := range v.Nodes {
		s.Nodes[int(id)] = node
	}
	s.ID = int(v.ID)
	s.LeaderID = int(v.LeaderID)
	s.LastHeartbeat = v.LastHeartbeat.AsDuration()
	s.HeartbeatInterval = v.HeartbeatInterval.AsDuration()
	s.ElectionTimeout = v.ElectionTimeout.AsDuration()
	s.OS = v.OS
	s.CPUArch = v.Arch
	s.CPUs = uint(v.CPUs)
	s.UsableCPUs = uint(v.UsableCPUs)
	s.HeapMemInUse = v.HeapMemInUse
	s.StackMemInUse = v.StackMemInUse
	return nil
}

// StatusResponse contains status information about a KMS cluster.
//
// The overall view of the current cluster status, in particular
// which nodes are reachable, may vary from node to node in case
// of network partitions. For example, two nodes within two network
// partitions will consider themselves as up and their peer as down.
type StatusResponse struct {
	// NodesUp is a map of node IDs to the corresponding node status
	// information.
	NodesUp map[int]*NodeStatusResponse

	// NodesDown is a map of node IDs to node addresses containing
	// all nodes that were not reachable or failed to respond in time.
	NodesDown map[int]string
}

// MarshalPB converts the StatusResponse into its protobuf representation.
func (s *StatusResponse) MarshalPB(v *pb.StatusResponse) error {
	v.NodesUp = make(map[uint32]*pb.NodeStatusResponse, len(s.NodesUp))
	for id, resp := range s.NodesUp {
		stat := new(pb.NodeStatusResponse)
		if err := resp.MarshalPB(stat); err != nil {
			return err
		}
		v.NodesUp[uint32(id)] = stat
	}

	v.NodesDown = make(map[uint32]string, len(s.NodesDown))
	for id, addr := range s.NodesDown {
		v.NodesDown[uint32(id)] = addr
	}
	return nil
}

// UnmarshalPB initializes the StatusResponse from its protobuf representation.
func (s *StatusResponse) UnmarshalPB(v *pb.StatusResponse) error {
	s.NodesUp = make(map[int]*NodeStatusResponse, len(v.NodesUp))
	for id, resp := range v.NodesUp {
		stat := new(NodeStatusResponse)
		if err := stat.UnmarshalPB(resp); err != nil {
			return err
		}
		s.NodesUp[int(id)] = stat
	}

	s.NodesDown = make(map[int]string, len(v.NodesDown))
	for id, addr := range v.NodesDown {
		s.NodesDown[int(id)] = addr
	}
	return nil
}

// EncryptResponse contains the ciphertext of an encrypted message
// and the key version used to encrypt the message.
type EncryptResponse struct {
	// Version identifies the particular key within a key ring used to encrypt
	// the message.
	Version int

	// Ciphertext is the encrypted message.
	Ciphertext []byte
}

// MarshalPB converts the EncryptResponse into its protobuf representation.
func (r *EncryptResponse) MarshalPB(v *pb.EncryptResponse) error {
	v.Version = uint32(r.Version)
	v.Ciphertext = r.Ciphertext
	return nil
}

// UnmarshalPB initializes the EncryptResponse from its protobuf representation.
func (r *EncryptResponse) UnmarshalPB(v *pb.EncryptResponse) error {
	r.Version = int(v.Version)
	r.Ciphertext = v.Ciphertext
	return nil
}

// DecryptResponse contains the decrypted plaintext message.
type DecryptResponse struct {
	// Plaintext is the decrypted message.
	Plaintext []byte
}

// MarshalPB converts the DecryptResponse into its protobuf representation.
func (r *DecryptResponse) MarshalPB(v *pb.DecryptResponse) error {
	v.Plaintext = r.Plaintext
	return nil
}

// UnmarshalPB initializes the DecryptResponse from its protobuf representation.
func (r *DecryptResponse) UnmarshalPB(v *pb.DecryptResponse) error {
	r.Plaintext = v.Plaintext
	return nil
}

// GenerateKeyResponse contains data encryption key that consists of a plaintext
// data encryption key and an encrypted ciphertext. Applications should use, but
// never store, the plaintext data encryption key for crypto. operations and store
// the ciphertext and key version.
type GenerateKeyResponse struct {
	// Version identifies the particular key within a key ring used to generate
	// and encrypt this data encryption key.
	Version int

	// Plaintext is the plain data encryption key. It may be used by applications
	// to perform crypto. operations.
	Plaintext []byte

	// Ciphertext is the encrypted data encryption key. Applications should store
	// it to obtain the plain data encryption key in the future again.
	Ciphertext []byte
}

// MarshalPB converts the GenerateKeyResponse into its protobuf representation.
func (r *GenerateKeyResponse) MarshalPB(v *pb.GenerateKeyResponse) error {
	v.Version = uint32(r.Version)
	v.Plaintext = r.Plaintext
	v.Ciphertext = r.Ciphertext
	return nil
}

// UnmarshalPB initializes the GenerateKeyResponse from its protobuf representation.
func (r *GenerateKeyResponse) UnmarshalPB(v *pb.GenerateKeyResponse) error {
	r.Version = int(v.Version)
	r.Plaintext = v.Plaintext
	r.Ciphertext = v.Ciphertext
	return nil
}