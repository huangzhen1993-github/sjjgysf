package main
//反转一个单链表。
//输入: 1->2->3->4->5->NULL
//分析：
//假设存在链表 1 → 2 → 3 → Ø，我们想要把它改成 Ø ← 1 ← 2 ← 3。
//在遍历列表时，将当前节点的 next 指针改为指向前一个元素。由于节点没有引用其上一个节点，因此必须事先存储其前一个元素。在更改引用之前，还需要另一个指针来存储下一个节点。不要忘记在最后返回新的头引用！

type listNode struct {
	Next *listNode
}

//迭代
func ReverseList(head *listNode) *listNode {
	var pre *listNode
	cur := head
	for cur != nil{
		tem := cur.Next
		cur.Next = pre
		pre = cur
		cur = tem
	}
	return  pre
}

//递归
func ReverseList1(head *listNode) *listNode{
	if head == nil || head.Next == nil{
		return head
	}
	l := ReverseList1(head.Next)
	head.Next.Next =head
	head = nil
	return l
}

//递归
func reverseList2(head *listNode) *listNode {
	return helper(nil, head)
}

func helper(prev, cur *listNode) *listNode {
	if cur == nil {
		return prev
	}
	next := cur.Next
	cur.Next = prev
	return helper(cur, next)
}

