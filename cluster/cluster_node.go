/*
 * @Author: SpenserCai
 * @Date: 2023-08-16 15:25:34
 * @version:
 * @LastEditors: SpenserCai
 * @LastEditTime: 2023-08-16 15:33:30
 * @Description: file content
 */
package cluster

import (
	"sd-webui-discord/queue"

	webui "github.com/SpenserCai/sd-webui-go"
)

type ClusterNode struct {
	Name         string
	ActionQueue  *queue.ActionQueue
	StableClient *webui.StableDiffInterface
}

func NewClusterNode(name string, actionQueue *queue.ActionQueue, stableClient *webui.StableDiffInterface) *ClusterNode {
	return &ClusterNode{
		Name:         name,
		ActionQueue:  actionQueue,
		StableClient: stableClient,
	}
}
