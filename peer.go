package geecache

import pb "GeeCache/geecachepb"

type PeerPicker interface {
	PickPeer(key string) (peer PeerGetter, ok bool)
}

type PeerGetter interface {
	// Get Get(group string, key string) ([]byte, error)
	Get(in *pb.Request, out *pb.Response) error
}

