package main

import "fmt"

type Subject interface {
	Request() string
}

type RealSubject struct{}

func (rs *RealSubject) Request() string {
	return "RealSubject: Handling request"
}

type Proxy struct {
	realSubject *RealSubject
}

func NewProxy() *Proxy {
	return &Proxy{}
}

func (p *Proxy) Request() string {
	if p.realSubject == nil {
		fmt.Println("Proxy: Creating RealSubject instance")
		p.realSubject = &RealSubject{}
	}

	fmt.Println("Proxy: Checking access before forwarding the request to RealSubject")
	return p.realSubject.Request()
}

func main() {
	proxy := NewProxy()

	result := proxy.Request()
	fmt.Println(result)
}
