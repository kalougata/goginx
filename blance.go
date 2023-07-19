package goginx

import (
	"hash/crc32"
	"sort"
	"strconv"
)

// 负载均衡器
//TODO 使用哈希一致性算法建构，参考https://juejin.cn/post/7134656152452726792

// 哈希环建构
type hashRing struct {
	ring  []int          //哈希环
	nodes map[int]string //节点哈希映射到节点名称
}

// 新建哈希节点，replicas为每个真实节点对应的虚拟节点数
func (location *location) addNode(engine *Engine) {
	upstream := engine.upstream[location.upstream]
	for _, node := range upstream {
		for i := 0; i < engine.replicas; i++ {
			hashValue := int(hash([]byte(strconv.Itoa(i) + node)))
			location.hashRing.ring = append(location.hashRing.ring, hashValue)
			location.hashRing.nodes[hashValue] = node
		}
	}
	sort.Ints(location.hashRing.ring)
}

// 计算哈希值
// TODO gpt写的
func hash(data []byte) uint32 {
	return crc32.ChecksumIEEE(data)
}
