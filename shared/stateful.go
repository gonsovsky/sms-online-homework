package shared

//Интерфейс и для Арбитра и для участников.
type IStateful interface {
	Get() (State)
}