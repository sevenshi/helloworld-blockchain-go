package netcore

/*
 @author king 409060350@qq.com
*/

import (
	"helloworld-blockchain-go/dto"
	"helloworld-blockchain-go/netcore-client/client"
	"helloworld-blockchain-go/netcore/configuration"
	"helloworld-blockchain-go/netcore/service"
	"helloworld-blockchain-go/util/LogUtil"
	"helloworld-blockchain-go/util/SystemUtil"
	"helloworld-blockchain-go/util/ThreadUtil"
)

type NodeCleaner struct {
	netCoreConfiguration *configuration.NetCoreConfiguration
	nodeService          *service.NodeService
}

func NewNodeCleaner(netCoreConfiguration *configuration.NetCoreConfiguration, nodeService *service.NodeService) *NodeCleaner {
	var nodeCleaner NodeCleaner
	nodeCleaner.netCoreConfiguration = netCoreConfiguration
	nodeCleaner.nodeService = nodeService
	return &nodeCleaner
}

func (b *NodeCleaner) start() {
	defer func() {
		if e := recover(); e != nil {
			SystemUtil.ErrorExit("清理死亡节点出现异常", e)
		}
	}()
	for {
		b.cleanDeadNodes()
		ThreadUtil.MillisecondSleep(b.netCoreConfiguration.GetNodeCleanTimeInterval())
	}
}

func (b *NodeCleaner) cleanDeadNodes() {
	nodes := b.nodeService.QueryAllNodes()
	if nodes == nil || len(nodes) == 0 {
		return
	}

	for _, node := range nodes {
		nodeClient := client.NewNodeClient(node.Ip)
		var pingRequest dto.PingRequest
		pingResponse := nodeClient.PingNode(pingRequest)
		if pingResponse == nil {
			b.nodeService.DeleteNode(node.Ip)
			LogUtil.Debug("节点清理器发现死亡节点[" + node.Ip + "]，已在节点数据库中将该节点删除了。")
		}
	}
}
