package ac

import (
	"container/list"
	"fmt"
)

/**
这里的最长可匹配后缀子串，不同于kmp中的最长可匹配后缀子串，
字符串 abc 的后缀子串有两个 bc，c，
我们拿它们与其他模式串匹配，
如果某个后缀子串可以匹配某个模式串的前缀，
那我们就把这个后缀子串叫作可匹配后缀子串
 */
type  AcNode struct {
	data byte
	children [26]*AcNode
	isEndingChar bool //是否是完整的字符串,比如插入he,her 则e也是字符结尾
	length int
	fail *AcNode
}

type Trie struct {
	root *AcNode
}

func NewTrie() *Trie {
	return &Trie{
		root:newAcNode('/'),
	}
}

func newAcNode(data byte) *AcNode  {
	return &AcNode{
		data:data,
		length:-1,
	}
}

func (tree *Trie) Insert(text string) bool{
	node := tree.root
	for _, v :=range text {
		index := v - 'a'
		if node.children[index] == nil {//子类不存在则创建
			node.children[index] = newAcNode(byte(v))
		}
		node = node.children[index]//走到下一节点
	}
	//完整字符串的前缀子串都存默认 false,这样查找的时候，就可以判断是否是完整的子串了
	node.isEndingChar = true
	node.length = len(text)
	return true
}

func (tree *Trie) Find(text string) bool {
	node := tree.root
	for _, v :=range text {
		index := v - 'a'
		if node.children[index] == nil {//子类不存在则创建
			return false
		}
		node = node.children[index]//走到下一节点
	}
	return node.isEndingChar
}

func (tree *Trie)BuildFailurePointer()  {
	queues := list.New()
	tree.root.fail = nil
	queues.PushBack(tree.root)

	for queues.Len() > 0 {//按层构造失败指针
		p := queues.Remove(queues.Front()).(*AcNode)
		for i:=0;i<26;i++{//遍历子节点
			pc := p.children[i]

			if pc == nil {//没有子节点
				continue
			}

			if p == tree.root {//如果时根节点的孩子则失败指针指向根节点
				pc.fail = tree.root
				queues.PushBack(pc)//将子节点加入,按层遍历用
				continue
			}

			//构造失败指针
			/**
			我们假设节点 p 的失败指针指向节点 q，
			我们看节点 p 的子节点 pc 对应的字符，
			是否也可以在节点 q 的子节点中找到。
			如果找到了节点 q 的一个子节点 qc，
			对应的字符跟节点 pc 对应的字符相同，
			则将节点 pc 的失败指针指向节点 qc。
			如果节点 q 中没有子节点的字符等于节点 pc 包含的字符，
			则令 q=q->fail（fail 表示失败指针，
			这里有没有很像 KMP 算法里求 next 的过程？），
			继续上面的查找，直到 q 是 root 为止，
			如果还没有找到相同字符的子节点，
			就让节点 pc 的失败指针指向 root。
			 */
			q := p.fail
			for q != nil  {
				qc := q.children[pc.data - 'a']
				if qc != nil {
					pc.fail = qc
					break
				}
				q = q.fail//如果
			}
			if q == nil {
				pc.fail = tree.root
			}
			queues.PushBack(pc)//将子节点加入,按层遍历用
		}
	}

}

func (tree *Trie)Search(search string)  {
	n := len(search)
	p := tree.root

	for i :=0;i<n ; i++ {
		idx := search[i] - 'a'
		for p.children[idx] == nil && p != tree.root  {//递归查找最长可匹配后缀子串
			p = p.fail
		}

		p = p.children[idx]
		if p == nil {
			p = tree.root
		}

		tmp := p
		for tmp != tree.root  {//从当前字符判断是否是结束符,在递归判断失败指针是否是是
			if tmp.isEndingChar{
				pos := i - tmp.length +1
				fmt.Printf("匹配起始下标:%v,长度:%v\n", pos, tmp.length)
			}
			tmp = tmp.fail
		}
	}
}
