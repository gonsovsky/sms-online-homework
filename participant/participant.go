package participant

import (
	"arbiter"
	"fmt"
	"github.com/nu7hatch/gouuid"
	"shared"
)

//Participant - Участник.
type Participant struct {
	Id     int
	config *shared.RedisCfg
	arbiter arbiter.IArbiter
	arbiterChannel  chan shared.State
	state shared.State //текущее cостояние
	received int //счетчик полученных собщений
	sent int //счетчик полученных собщений
}

//Put - Отправить комманду Арбитру
func (p *Participant) Put(val uuid.UUID) (error) {
	s := shared.NewState(val, p.Id)
	return p.arbiter.Put(s)
}

//Listen - слушать Арбитра
func (p *Participant) Listen() {
	p.arbiterChannel = p.arbiter.Hello(p)
	go func(){
		for {
			select {
			case x, ok := <-p.arbiterChannel:
				if ok {
					p.received++
					p.state = x
				} else {
					fmt.Println("Участник ",p.Id, " получил от арбитра ",p.received, " состояний")
					return
				}
			default:
				//nothing
			}
		}
	}()
}

//Speek - Отправка запросов к Арбитру
func (p *Participant) Speek() {
	go func() {
		for {
			v,_ := uuid.NewV4()
			err := p.Put(*v)
			if (err != nil){
				fmt.Println("Участник ",p.Id, " успел отправить",p.sent, " состояний.")
				return
			}else {
				p.sent++}
		}
	}()
}

func (a *Participant) Get() (shared.State) {
	return a.state
}

//NewParticipant - Создать участника.
func NewParticipant(config *shared.RedisCfg, id int, arbiter arbiter.IArbiter) *Participant {
	p := Participant{config: config, Id: id, arbiter: arbiter}
	return &p
}
