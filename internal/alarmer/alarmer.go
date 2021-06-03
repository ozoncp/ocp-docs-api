package alarmer

type Alarmer interface {
	Alarm()
	Init()
	Close()
}

type alarmer struct {

}

func New() Alarmer {
	return &alarmer{}
}

func (a * alarmer) Alarm() {

}

func (a * alarmer) Init() {

}

func (a * alarmer) Close() {

}