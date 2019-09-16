package main

type Executor interface {
	Exec()
}

type AbstractClass struct {
}

func (c *AbstractClass) Step1() {

}

func (c *AbstractClass) Step2() {

}

func (c *AbstractClass) Exec() {
	c.Step1()
	c.Step2()
}

type ConcreteClass struct {
	AbstractClass
}

func (c *ConcreteClass) Step1()  {
	
}

func (c *ConcreteClass) Step2()  {
	
}

func main()  {
	executor := ConcreteClass{}

	executor.Exec()
}