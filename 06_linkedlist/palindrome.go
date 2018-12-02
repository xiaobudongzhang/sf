package _6_linkedlist


/*
1 快慢指针定位中间节点
2 从中间节点对后半部分逆序
3 前后半部分比较，判断是否为回文
4 后半部分逆序复原
---------------------------
1 快慢指针定位中间节点（这里要区分奇偶情况）
1.1 奇数情况，中点位置不需要矫正
1.2 偶数情况，使用偶数定位中点策略，要确定是返回上中位数或下中位数
1.2.1 如果是返回上中位数，后半部分串头取next
1.2.2 如果是返回下中位数，后半部分串头既是当前节点位置，但前半部分串尾要删除掉当前节点
2 从中间节点对后半部分逆序，或者将前半部分逆序
3 一次循环比较，判断是否为回文
4 恢复现场*/
func IsPalindrome(l *LinkedList) bool {
	if l.length < 1 {
		return false
	}
	if l.length == 1 {
		return true
	}
	//1 快慢指针定位中间节点
	stepOld := l.length / 2
	step := stepOld
	if l.length % 2 != 0{
		step++
	}
	midP := l.FindByIndex(step)
	//2 从中间节点对后半部分逆序
	l.Reverse(midP)
	//3 前后半部分比较，判断是否为回文
	prev := l.head.next
	next := midP.next

	for i := uint(0); i < stepOld; i++ {
		if prev.GetValue() != next.GetValue() {
			return false
		}
		prev = prev.GetNext()
		next = next.GetNext()
	}
	//4 后半部分逆序复原
	l.Reverse(midP)
	return true
}
