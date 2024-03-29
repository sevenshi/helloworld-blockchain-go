package netcore

/*
 @author king 409060350@qq.com
*/

import (
	"helloworld-blockchain-go/dto"
	"helloworld-blockchain-go/netcore-client/client"
	"helloworld-blockchain-go/netcore/configuration"
	"helloworld-blockchain-go/netcore/service"
	"helloworld-blockchain-go/util/SystemUtil"
	"helloworld-blockchain-go/util/ThreadUtil"
)

type NodeBroadcaster struct {
	netCoreConfiguration *configuration.NetCoreConfiguration
	nodeService          *service.NodeService
}

func NewNodeBroadcaster(netCoreConfiguration *configuration.NetCoreConfiguration, nodeService *service.NodeService) *NodeBroadcaster {
	var nodeBroadcaster NodeBroadcaster
	nodeBroadcaster.netCoreConfiguration = netCoreConfiguration
	nodeBroadcaster.nodeService = nodeService
	return &nodeBroadcaster
}

func (b *NodeBroadcaster) start() {
	defer func() {
		if e := recover(); e != nil {
			SystemUtil.ErrorExit("在区块链网络中广播自己出现异常", e)
		}
	}()
	for {
		b.broadcastNode()
		ThreadUtil.MillisecondSleep(b.netCoreConfiguration.GetNodeBroadcastTimeInterval())
	}
}

func (b *NodeBroadcaster) broadcastNode() {
	nodes := b.nodeService.QueryAllNodes()
	if nodes == nil || len(nodes) == 0 {
		return
	}

	for _, node := range nodes {
		nodeClient := client.NewNodeClient(node.Ip)
		var pingRequest dto.PingRequest
		nodeClient.PingNode(pingRequest)
	}
}
