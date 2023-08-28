package main

type MyQueue struct {
	in, out []int
}

func Constructor() *MyQueue {
	return &MyQueue{}
}

func (this *MyQueue) push(x int) {
	this.in = append(this.in, x)
}
func (this *MyQueue) in2() {
	if len(this.in) > 0 {
		this.out = append(this.out, this.in[len(this.in)-1])
		this.in = this.in[:len(this.in)-1]
	}
}
func (this *MyQueue) pop() (x int) {
	if len(this.out) == 0 {
		this.in2()
	}
	x = this.out[len(this.out)-1]
	this.out = this.out[:len(this.out)-1]
	return
}

func (this *MyQueue) peek() int {
	if len(this.out) == 0 {
		this.in2()
	}
	return this.out[len(this.out)-1]

}

func (this *MyQueue) empty() bool {
	if len(this.in) == 0 && len(this.out) == 0 {
		return true
	}
	return false
}
