package discovery

type RequestType byte

const (
	joinRequest RequestType = iota
	leaveRequest
	watchRequest
	snapshotRequest
	heartbeatRequest
	lastRequestType
)

type Request interface {
	Type() RequestType
}

type JoinRequest struct {
	Request
	Host       string
	Port       uint16
	Group      string
	CustomData []byte
}

func (r *JoinRequest) Type() RequestType      { return joinRequest }
func (r *LeaveRequest) Type() RequestType     { return leaveRequest }
func (r *WatchRequest) Type() RequestType     { return watchRequest }
func (r *SnapshotRequest) Type() RequestType  { return snapshotRequest }
func (r *HeartbeatRequest) Type() RequestType { return heartbeatRequest }

type LeaveRequest struct {
	Request
	Host  string
	Port  uint16
	Group string
}

// A snapshot only applies to a single group.
type SnapshotRequest struct {
	Request
	Group string
}

// A watch request contains a list of groups to watch on the connection.
type WatchRequest struct {
	Request
	Groups []string
}

type HeartbeatRequest struct {
	Request
}
