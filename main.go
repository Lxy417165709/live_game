package main


// ------------------------ MyMap ------------------------
type MyMap struct{
	Cap int
	Nodes []*Node
	mutex Mutex
}

func NewMap(cap int) *MyMap{
	return &MyMap{
		Cap:cap,
		Nodes : make([]*Node,cap)
	}
}

func (m *MyMap) RangeAll() []int{
	m.mutex.Lock()
	defer m.mutex.UnLock()

	keys := make([]int,0)
	for _, nodeHead := range m.Nodes {
		curHead := nodeHead
		for curHead != nil{
			keys = append(keys,curHead.Key)
			curHead = curHead.Next
		}
	}
	sort.Ints(keys)
	return keys
}


func (m *MyMap) Insert(key,value int) {
	m.mutex.Lock()
	defer m.mutex.UnLock()

	insertIndex := key%m.Cap
	head := m.Nodes[insertIndex]
	insertNdoe :=  &Node{
		Key:key,
		Value:value,
	}
	if head == nil{
		m.Nodes[insertIndex] = insertNdoe
		return
	}
	for head.Next!=nil{
		head = head.Next
	}
	head.Next = insertNdoe
}


// ------------------------ 底层结构 ------------------------
type Node struct{
	Key int
	Value interface{}
	Next *Node
}

func main() {

}