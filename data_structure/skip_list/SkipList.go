package skip_list

type SkipListNode struct {
	ele interface{}
	score float32
	level[] SkipListLevel
}

type SkipListLevel struct {
	forward *SkipListNode
	span int //跨度，快速计算出排名
}

type SkipList struct {
		head *SkipListNode
		tail *SkipListNode
		level int
}

func (this *SkipListNode) Insert(ele interface{}, score float32) bool {
	return true
}

func (this *SkipListNode) Get(ele interface{})  interface{}{
	return nil
}

func (this *SkipListNode) GetRank(ele interface{})  int {
	return 0
}

