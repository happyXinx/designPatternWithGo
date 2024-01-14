package adapter

type Target interface {
	Request() string
}

type Adaptee interface {
	SpecificRequest() string
}

func NewAdatpee() Adaptee {
	return &adapteeImpl{}
}

type adapteeImpl struct {
}

func (r *adapteeImpl) SpecificRequest() string {
	return "adapteeimpl"
}

type adapter struct {
	Adaptee
}

func NewAdapter(adaptee Adaptee) Target {
	return &adapter{adaptee}
}

func (r *adapter) Request() string {
	return r.SpecificRequest()
}
