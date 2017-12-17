package store

import (
	"fmt"
	"github.com/chrislusf/vasto/storage/binlog"
	"github.com/chrislusf/vasto/storage/rocks"
	"github.com/chrislusf/vasto/topology"
	"github.com/chrislusf/vasto/topology/cluster_listener"
	"log"
	"context"
)

type shard struct {
	keyspace          string
	id                int
	serverId          int
	db                *rocks.Rocks
	lm                *binlog.LogManager
	clusterRing       *topology.ClusterRing
	clusterListener   *cluster_listener.ClusterListener
	replicationFactor int
	nodeFinishChan    chan bool
	cancelFunc        context.CancelFunc
	isShutdown        bool
	// just to avoid repeatedly create these variables
	nextSegmentKey, nextOffsetKey []byte
	prevSegment                   uint32
	prevOffset                    uint64
	nextSegment                   uint32
	nextOffset                    uint64
}

func (s *shard) String() string {
	return fmt.Sprintf("%s.%d.%d", s.keyspace, s.serverId, s.id)
}

func newShard(keyspaceName, dir string, serverId, nodeId int, cluster *topology.ClusterRing,
	clusterListener *cluster_listener.ClusterListener,
	replicationFactor int, logFileSizeMb int, logFileCount int) (context.Context, *shard) {

	ctx, cancelFunc := context.WithCancel(context.Background())

	s := &shard{
		keyspace:          keyspaceName,
		id:                nodeId,
		serverId:          serverId,
		db:                rocks.New(dir),
		clusterRing:       cluster,
		clusterListener:   clusterListener,
		replicationFactor: replicationFactor,
		nodeFinishChan:    make(chan bool),
		cancelFunc:        cancelFunc,
	}
	if logFileSizeMb > 0 {
		s.lm = binlog.NewLogManager(dir, nodeId, int64(logFileSizeMb*1024*1024), logFileCount)
		s.lm.Initialze()
	}
	s.nextSegmentKey = []byte(fmt.Sprintf("%d.next.segment", s.id))
	s.nextOffsetKey = []byte(fmt.Sprintf("%d.next.offset", s.id))

	return ctx, s
}

func (s *shard) startWithBootstrapAndFollow(ctx context.Context, mayBootstrap bool) {

	if s.clusterRing != nil && mayBootstrap {
		err := s.bootstrap(ctx)
		if err != nil {
			log.Printf("bootstrap: %v", err)
		}
	}

	s.follow(ctx)

	s.clusterListener.RegisterShardEventProcessor(s)

}

func (s *shard) shutdownNode() {

	s.isShutdown = true

	s.cancelFunc()

	s.clusterListener.RemoveKeyspace(s.keyspace)

	s.clusterListener.UnregisterShardEventProcessor(s)

	close(s.nodeFinishChan)

	if s.lm != nil {
		s.lm.Shutdown()
	}

}

func (s *shard) setCompactionFilterClusterSize(clusterSize int) {

	s.db.SetCompactionForShard(s.id, clusterSize)

}