package main

import "fmt"

type SubsystemA struct {
}

func (s *SubsystemA) OperationA1() {
	fmt.Println("SubsystemA: OperationA1")
}

func (s *SubsystemA) OperationA2() {
	fmt.Println("SubsystemA: OperationA2")
}

type SubsystemB struct {
}

func (s *SubsystemB) OperationB1() {
	fmt.Println("SubsystemB: OperationB1")
}

func (s *SubsystemB) OperationB2() {
	fmt.Println("SubsystemB: OperationB2")
}

type Facade struct {
	subsystemA *SubsystemA
	subsystemB *SubsystemB
}

func NewFacade() *Facade {
	return &Facade{
		subsystemA: &SubsystemA{},
		subsystemB: &SubsystemB{},
	}
}

func (f *Facade) Operation() {
	fmt.Println("Facade: Operation")
	f.subsystemA.OperationA1()
	f.subsystemA.OperationA2()
	f.subsystemB.OperationB1()
	f.subsystemB.OperationB2()
}

func main() {
	facade := NewFacade()
	facade.Operation()
}
