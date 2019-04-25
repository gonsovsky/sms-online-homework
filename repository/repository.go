package repository

import "shared"

//IRepository - Репозиторий
type IRepository interface
{
	Open() error
	Get() (shared.State)
	Put(state shared.State) error
	Close()
}
