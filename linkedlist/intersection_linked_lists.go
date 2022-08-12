package main
//题目：
//给你两个单链表的头节点 headA 和 headB ，请你找出并返回两个单链表相交的起始节点。如果两个链表不存在相交节点，返回 null

//暴力解法
func GetIntersectionNode(headA,headB *listNode) *listNode {
	var _intersect *listNode
	for headA != nil {
		for headB != nil{
			if headA == headB{
				_intersect = _intersect
			}
			headB = headB.Next
		}
		headA = headA.Next
	}
	return _intersect
}

//使用一个哈希表判重
func GetIntersectionNode1(headA,headB *listNode) *listNode {
	s := make(map[*listNode]bool)
	for headA != nil {
		s[headA], headA = true, headA.Next
	}
	for headB != nil {
		if s[headB] {
			return headB
		}
		headB = headB.Next
	}
	return nil
}

// 双指针法
func getIntersectionNode3(headA, headB *listNode) *listNode {
	p, q := headA, headB
	for p != q {
		if p == nil {
			p = headB //当A链表遍历完成时，p赋值成B链表，保证循环退出
		} else {
			p = p.Next
		}
		if q == nil {
			q = headA
		} else {
			q = q.Next
		}
	}
	return p
}