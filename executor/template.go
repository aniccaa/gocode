package main

type TemplateInterface interface {
	Step1()
	Step2()
}

type AbstractClazz struct {
	templateImpl TemplateInterface
}

func (c *AbstractClazz) Exec() {
	c.templateImpl.Step1()
	c.templateImpl.Step2()
}

type ConcreteClazz struct {

}

func (c *ConcreteClazz) Step1()  {

}

func (c *ConcreteClazz) Step2()  {

}

func main()  {
	executor := AbstractClazz{templateImpl:new(ConcreteClazz)}
	executor.Exec()
}
