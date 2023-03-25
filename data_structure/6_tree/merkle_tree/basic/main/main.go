package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

// MerkleNode 表示 Merkle 树中的一个节点
type MerkleNode struct {
	left  *MerkleNode // 左子节点
	right *MerkleNode // 右子节点
	hash  string      // 当前节点的哈希值
}

// MerkleTree 表示一个 Merkle 树
type MerkleTree struct {
	root *MerkleNode // 根节点
}

// NewMerkleNode 创建一个新的 Merkle 节点
func NewMerkleNode(left, right *MerkleNode, hash string) *MerkleNode {
	return &MerkleNode{left, right, hash}
}

// NewMerkleTree 创建一个新的 Merkle 树
func NewMerkleTree(data [][]byte) *MerkleTree {
	var nodes []*MerkleNode

	// 如果数据中的元素个数是奇数，则复制最后一个元素
	if len(data)%2 != 0 {
		data = append(data, data[len(data)-1])
	}

	// 创建叶子节点
	for _, datum := range data {
		hash := sha256.Sum256(datum)
		hashString := hex.EncodeToString(hash[:])
		node := NewMerkleNode(nil, nil, hashString)
		nodes = append(nodes, node)
	}

	// 逐层构建 Merkle 树
	for i := 0; i < len(data)/2; i++ {
		var level []*MerkleNode

		for j := 0; j < len(nodes); j += 2 {
			left := nodes[j]
			right := nodes[j+1]
			hash := sha256.Sum256([]byte(left.hash + right.hash))
			hashString := hex.EncodeToString(hash[:])
			node := NewMerkleNode(left, right, hashString)
			level = append(level, node)
		}

		nodes = level
	}

	return &MerkleTree{nodes[0]}
}

// GetProof 获取指定数据元素的证明
func (tree *MerkleTree) GetProof(data []byte) []string {
	var proof []string
	hash := sha256.Sum256(data)
	hashString := hex.EncodeToString(hash[:])
	node := tree.root

	for node.left != nil && node.right != nil {
		if node.left.hash == hashString {
			proof = append(proof, node.right.hash)
			hashString = node.left.hash + node.right.hash
			hash = sha256.Sum256([]byte(hashString))
			hashString = hex.EncodeToString(hash[:])
			node = node.left
		} else {
			proof = append(proof, node.left.hash)
			hashString = node.left.hash + node.right.hash
			hash = sha256.Sum256([]byte(hashString))
			hashString = hex.EncodeToString(hash[:])
			node = node.right
		}
	}

	return proof
}

// VerifyProof 验证指定数据元素的证明是否正确
func VerifyProof(rootHash string, data []byte, proof []string) bool {
	hash := sha256.Sum256(data)
	hashString := hex.EncodeToString(hash[:])

	for _, h := range proof {
		if h != hashString {

			// 如果证明中的哈希值与当前数据元素的哈希值不一致，则证明无效
			left := hashString + h
			right := h + hashString
			hash = sha256.Sum256([]byte(left))
			hashString = hex.EncodeToString(hash[:])
			if h == right {
				hash = sha256.Sum256([]byte(right))
				hashString = hex.EncodeToString(hash[:])
			}
		}
	}

	// 最终验证根哈希值是否与给定的根哈希值一致
	return hashString == rootHash
}

// Insert 向 Merkle 树中插入一个新的数据元素
func (tree *MerkleTree) Insert(data []byte) {
	hash := sha256.Sum256(data)
	hashString := hex.EncodeToString(hash[:])
	node := tree.root
	// 找到插入位置
	for node.left != nil && node.right != nil {
		if node.left.hash == hashString {
			hashString = node.left.hash + node.right.hash
			hash = sha256.Sum256([]byte(hashString))
			hashString = hex.EncodeToString(hash[:])
			node = node.left
		} else {
			hashString = node.left.hash + node.right.hash
			hash = sha256.Sum256([]byte(hashString))
			hashString = hex.EncodeToString(hash[:])
			node = node.right
		}
	}

	// 插入新的叶子节点
	newNode := NewMerkleNode(nil, nil, hashString)
	if node.left == nil {
		node.left = newNode
	} else {
		node.right = newNode
	}

	// 更新父节点的哈希值
	for node != nil {
		if node.left == nil || node.right == nil {
			break
		}

		hashString = node.left.hash + node.right.hash
		hash = sha256.Sum256([]byte(hashString))
		node.hash = hex.EncodeToString(hash[:])
		node = node.left
	}
}

// Delete 从 Merkle 树中删除一个数据元素
func (tree *MerkleTree) Delete(data []byte) bool {
	hash := sha256.Sum256(data)
	hashString := hex.EncodeToString(hash[:])
	node := tree.root
	var parent *MerkleNode
	// 找到要删除的节点
	for node.left != nil && node.right != nil {
		if node.left.hash == hashString {
			parent = node
			node = node.left
		} else if node.right.hash == hashString {
			parent = node
			node = node.right
		} else {
			return false
		}
	}

	// 删除叶子节点
	if node.left == nil && node.right == nil {
		if parent.left == node {
			parent.left = nil
		} else {
			parent.right = nil
		}
	} else {
		return false
	}

	// 更新父节点的哈希值
	for parent != nil {
		if parent.left == nil || parent.right == nil {
			break
		}

		hashString = parent.left.hash + parent.right.hash
		hash = sha256.Sum256([]byte(hashString))
		parent.hash = hex.EncodeToString(hash[:])
		parent = parent.left
	}

	return true
}
func main() {
	// 创建 Merkle 树
	data := [][]byte{
		[]byte("hello"),
		[]byte("world"),
		[]byte("test"),
	}
	tree := NewMerkleTree(data)
	// 向 Merkle 树中插入一个新的数据元素
	tree.Insert([]byte("new data"))

	// 生成证明
	proof := tree.GetProof([]byte("hello"))

	// 验证证明
	if VerifyProof(tree.root.hash, []byte("hello"), proof) {
		fmt.Println("Proof is valid")
	} else {
		fmt.Println("Proof is invalid")
	}

	// 从 Merkle 树中删除一个数据元素
	if tree.Delete([]byte("world")) {
		fmt.Println("Data deleted successfully")
	} else {
		fmt.Println("Data not found")
	}
}
