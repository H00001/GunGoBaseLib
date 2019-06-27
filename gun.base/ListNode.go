package gun_base

type listNode struct {
	next  *listNode
	value interface{}
	prev  *listNode
}

type LinkedList struct {
	head      *listNode
	emptyNode *listNode
}

func (this *LinkedList) addLast(value interface{}) {
	var node = new(listNode)
	node.value = value
	this.emptyNode.next = node
	this.emptyNode = this.emptyNode.next

}
