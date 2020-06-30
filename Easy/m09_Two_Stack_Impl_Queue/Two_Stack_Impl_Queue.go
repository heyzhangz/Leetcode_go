package questionm02_01
/*
	剑指 Offer 09. 用两个栈实现队列
	用两个栈实现一个队列。队列的声明如下，请实现它的两个函数 appendTail 和 deleteHead ，分别完成在队列尾部插入整数和在队列头部删除整数的功能。(若队列中没有元素，deleteHead 操作返回 -1 )
 */

/*
	思路: 一个临时栈存出的顺序？
	题解:
        Go 有个List的包
*/


type CQueue struct {
    mainstack []int
    tempstack []int
    length int
}


func Constructor() CQueue {

    obj := CQueue{mainstack: []int{}, tempstack: []int{}, length: 0}
    return obj
}


func (this *CQueue) AppendTail(value int)  {

    this.mainstack = append(this.mainstack, value)
    this.length++
}


func (this *CQueue) DeleteHead() int {

    if this.length == 0 {
        return -1
    }

    // 这里不用每次都倒腾一遍，当第二个栈为空的时候再倒腾就行
    if len(this.tempstack) <= 0 {
        for len(this.mainstack) > 0 {
            a := this.mainstack[len(this.mainstack) - 1]
            this.mainstack = this.mainstack[ : len(this.mainstack) - 1]
            this.tempstack = append(this.tempstack, a)
        }
    }

    del := this.tempstack[len(this.tempstack) - 1]
    this.tempstack = this.tempstack[ : len(this.tempstack) - 1]

    //for len(this.tempstack) > 0 {
    //    a := this.tempstack[len(this.tempstack) - 1]
    //    this.tempstack = this.tempstack[ : len(this.tempstack) - 1]
    //    this.mainstack = append(this.mainstack, a)
    //}

    this.length--
    return del
}


/**
 * Your CQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AppendTail(value);
 * param_2 := obj.DeleteHead();
 */