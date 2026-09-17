type MinStack struct {
	data *Data
}

type Data struct {
	value int
	min   int
	next  *Data
}

func Constructor() MinStack {
	return MinStack{}
}

func (this *MinStack) Push(val int) {
	if this.data == nil {
		this.data = &Data{
			value: val,
			min:   val,
		}
		return
	}

	currentMin := this.data.min

	if val < currentMin {
		currentMin = val
	}

	newData := &Data{
		value: val,
		min:   currentMin,
		next:  this.data,
	}

	this.data = newData
}

func (this *MinStack) Pop() {
	if this.data != nil {
		this.data = this.data.next
	}
}

func (this *MinStack) Top() int {
	return this.data.value
}

func (this *MinStack) GetMin() int {
	return this.data.min
}