package _8_heap


func (this *Heap)Init(arr []int)  {
	for i:=0;i<len(arr) ;i++  {
		this.data[i+1] = arr[i]
	}
	this.count = len(arr)
}

func (this *Heap)BuildHeap() {
	for i := this.count / 2; i >= 1;i-- {
		this.Heap(this.count, i)
	}
}

func (this *Heap)Sort() {
	for i:=this.count;i >1 ; i-- {
		this.data[i],this.data[1] = this.data[1],this.data[i]
		this.Heap(i-1, 1)
	}
}
