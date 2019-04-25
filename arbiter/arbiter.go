package arbiter

import (
	"fmt"
	"repository"
	"shared"
)

//Arbiter - Арбитр. Синхронизирует работу участников с Репозиторием.
type IArbiter interface
{
	//Serve - Начать обслуживание участников
	Serve()
	//Hello Регистрация нового участника
	Hello(p shared.IStateful) (chan shared.State)
	//Завершить обслуживание
	Quit()
	//PutState - Принять комманду от участника на запись в Репозиторий
	Put(msg shared.State) error
	//GetState - Прочитать состоние из Репозития
	Get() (shared.State)
}

//Arbiter - Арбитр. Синхронизирует работу участников с Репозиторием.
type Arbiter struct {
	Participants map[shared.IStateful](chan shared.State)
	repo repository.IRepository
	state shared.State
	input chan shared.State
	quit chan bool
}

//Hello Зарегестировать участника и возвратить ему канал для чтения
func (a *Arbiter) Hello(p shared.IStateful) (chan shared.State) {
	a.Participants[p] =  make(chan shared.State)
	return a.Participants[p]
}

//Завершить обслуживание
func (a *Arbiter) Quit() {
	a.quit <- true
	close(a.input)
	for _,clientChannel := range a.Participants{
		close(clientChannel)
	}
}

//PutState Принять комманду от участника на запись в Репозитий
func (a *Arbiter) Put(state shared.State) (output error) {
	defer func(){
		if rec := recover(); rec != nil {
			output= shared.ErrorOutOfService
		}
	}()
	a.input <- state
	return nil
}

func (a *Arbiter) Get() (shared.State) {
	return a.state
}

//Serve - Начать обслуживание участников
func (a *Arbiter) Serve() {
	a.input = make(chan shared.State)
	a.quit = make(chan bool)
	a.Participants = map[shared.IStateful](chan shared.State){}
	go func(){
		for {
			select {
			case _,ok := <- a.quit:
				if ok {
					fmt.Println("Arbiter's is closing. Unhandled: ",len(a.input), " requests.")
					return
					} else {panic("look for bug")}
			case x, ok := <-a.input:
				if ok {
					err := a.repo.Put(x)
					if (err != nil){
						fmt.Println(err.Error())
					}else {
						a.ToAll(x)
						a.state = x
					}
				} else {
					panic("look for bug")
					return
				}
			default:
				//nothing
			}
		}
	}()
}

//Переслать состояние от одного участника ко всем
func (arbiter *Arbiter) ToAll(state shared.State) {
	for _, ch := range arbiter.Participants {
		ch <- state;
	}
};

//NewArbiter - Создать арбитра.
func NewArbiter(repo repository.IRepository) *Arbiter {
	p := Arbiter{repo: repo}
	return &p
}