package main

import (
	"fmt"
	"strings"
	"sync"
)

// KeyValue 通用键值对
type KeyValue struct {
	Key   string
	Value int
}

// ==================== 配置 ====================
const (
	mapWorkerNum    = 3 // 3个Map节点
	reduceWorkerNum = 2 // 2个Reduce节点
)

// ==================== 业务自定义函数 ====================
// Map：一行文本拆单词，输出 <word,1>
func mapFunc(line string) []KeyValue {
	var kvs []KeyValue
	words := strings.Fields(line)
	for _, w := range words {
		kvs = append(kvs, KeyValue{Key: w, Value: 1})
	}
	return kvs
}

// Reduce：同key累加
func reduceFunc(key string, vals []int) int {
	sum := 0
	for _, v := range vals {
		sum += v
	}
	return sum
}

// ==================== 分布式MR核心 ====================

// 分区：把key分到指定reduce节点
func partition(key string) int {
	return len(key) % reduceWorkerNum
}

// 分布式Map阶段：多worker并行处理
func runMap(splitLines [][]string, mapOutChan chan<- KeyValue) {
	var wg sync.WaitGroup
	wg.Add(mapWorkerNum)

	for i := 0; i < mapWorkerNum; i++ {
		idx := i
		go func() {
			defer wg.Done()
			// 当前worker处理自己的数据分片
			for _, line := range splitLines[idx] {
				kvs := mapFunc(line)
				for _, kv := range kvs {
					mapOutChan <- kv
				}
			}
		}()
	}

	// 所有Map跑完关闭通道
	go func() {
		wg.Wait()
		close(mapOutChan)
	}()
}

// Shuffle：按分区把相同key聚合，分发给Reduce
func runShuffle(mapOutChan <-chan KeyValue) []map[string][]int {
	// 每个Reduce对应一个分组map
	reduceGroups := make([]map[string][]int, reduceWorkerNum)
	for i := range reduceGroups {
		reduceGroups[i] = make(map[string][]int)
	}

	// 路由KV到对应Reduce分组
	for kv := range mapOutChan {
		rid := partition(kv.Key)
		reduceGroups[rid][kv.Key] = append(reduceGroups[rid][kv.Key], kv.Value)
	}
	return reduceGroups
}

// 分布式Reduce阶段：多worker并行聚合
func runReduce(reduceGroups []map[string][]int) map[string]int {
	result := make(map[string]int)
	var mu sync.Mutex
	var wg sync.WaitGroup

	wg.Add(reduceWorkerNum)
	for i := 0; i < reduceWorkerNum; i++ {
		idx := i
		go func() {
			defer wg.Done()
			// 处理当前Reduce分区的所有key
			for k, vals := range reduceGroups[idx] {
				res := reduceFunc(k, vals)
				mu.Lock()
				result[k] = res
				mu.Unlock()
			}
		}()
	}
	wg.Wait()
	return result
}

// 入口：分布式MapReduce总调度
func DistributedMapReduce(data []string) map[string]int {
	// 1. 数据分片：分给每个Map Worker
	split := make([][]string, mapWorkerNum)
	for i, line := range data {
		idx := i % mapWorkerNum
		split[idx] = append(split[idx], line)
	}

	mapOutChan := make(chan KeyValue, 100)

	// 2. 并行Map
	runMap(split, mapOutChan)

	// 3. Shuffle 分区分组
	reduceGroups := runShuffle(mapOutChan)

	// 4. 并行Reduce
	return runReduce(reduceGroups)
}

// ==================== 测试运行 ====================
func main() {
	// 原始海量数据
	data := []string{
		"hello go mapreduce",
		"hello go runtime",
		"go mapreduce golang",
		"hello distributed go",
		"mapreduce go design",
	}

	res := DistributedMapReduce(data)

	// 打印结果
	for k, v := range res {
		fmt.Printf("%-15s => %d\n", k, v)
	}
}
