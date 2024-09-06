package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// TreeNode 定义二叉树节点，使用 string 类型
type TreeNode struct {
	Val   string    `json:"val"`   // 节点值类型为 string
	Left  *TreeNode `json:"left"`  // 左子树
	Right *TreeNode `json:"right"` // 右子树
}

// BuildTree 根据给定的 BFS 序列构建二叉树
// 处理空字符串 ("") 和 nil，生成空节点
func BuildTree(bfs []string) *TreeNode {
	if len(bfs) == 0 {
		return nil
	}

	nodes := make([]*TreeNode, len(bfs))

	// 构造每个节点
	for i, val := range bfs {
		if val == "" { // 处理空节点
			nodes[i] = nil
		} else {
			nodes[i] = &TreeNode{Val: val}
		}
	}

	// 建立父节点与左右子节点的连接
	for i := 0; i < len(bfs); i++ {
		if nodes[i] == nil {
			continue
		}
		leftIndex := i*2 + 1
		rightIndex := i*2 + 2

		if leftIndex < len(bfs) {
			nodes[i].Left = nodes[leftIndex]
		}
		if rightIndex < len(bfs) {
			nodes[i].Right = nodes[rightIndex]
		}
	}

	return nodes[0]
}

// GenerateTreeHandler 处理生成二叉树的 API 请求
func GenerateTreeHandler(c *gin.Context) {
	var req struct {
		BFS []string `json:"bfs"` // 使用 string 类型处理 BFS 序列
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	root := BuildTree(req.BFS)
	c.JSON(http.StatusOK, root)
}
