package main

import (
	"arbiter"
	"fmt"
	"participant"
	"repository"
	"shared"
	"time"
)

func main() {
	//Репозиторий
	repo := repository.NewRedis(shared.RedisConfig())
	err := repo.Open()
	if (err != nil) {
		panic(err)
	}

	//Арбитр
	arbt := arbiter.NewArbiter(repo)
	arbt.Serve()

	//Участники
	cnt := shared.AppConfig().Participants;
	for i := 1; i <= cnt; i++ {
		p := participant.NewParticipant(shared.RedisConfig(), i, arbt)
		p.Listen()
		p.Speek()
	}

	//остановить, проверить состояние
	//Репозиторий = Арбитр = Все Участники
	time.Sleep(5* time.Second)
	arbt.Quit()
	allStates := []shared.State{
		arbt.Get(),
		repo.Get(),
	}
	for p, _ := range arbt.Participants {
		allStates = append(allStates, p.Get())
	}
	if shared.AllEquals(allStates) == false {
		fmt.Println("Где-то ошибка:", allStates)
	} else {
		fmt.Println("Стороны синхронизированы:")
		shared.Pretty(allStates) //вывести на экран
	}
}

