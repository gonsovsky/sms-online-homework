package shared

import (
	"encoding/json"
	"fmt"
	uuid "github.com/nu7hatch/gouuid"
	"time"
)

//State - состояние от участника к Арбитру и от Арбитра ко всем участникам
type State struct {
	RequestTime   time.Time `json:"requestTime"`
	Value	      uuid.UUID `json:"value"`
	ParticipantId int       `json:"participantId"`
}

//ToJSON convert struct to Json byte array
func (state *State) ToJSON() []byte {
	output, err := json.Marshal(state)
	if err != nil {
		panic(err)
	}
	return output
}

//FromJSON Fill in with data from byte[]
func FromJSON(str string) (State) {
	bytes := ([]byte)(str)
	state := State{}
	json.Unmarshal(bytes, &state)
	return state
}

//NewState return new State struct
func NewState(val uuid.UUID, partId int) State {
	return State{Value: val, ParticipantId: partId, RequestTime: time.Now().UTC()}
}

func (s *State) Equals(a State) (bool){
	return s.Value == a.Value && s.RequestTime == a.RequestTime && s.ParticipantId == a.ParticipantId
}

func AllEquals(states []State) (bool) {
	for i:=0; i<= len(states)-1-1; i++{
		if states[i].Equals(states[i+1]) == false{
			return false
		}
	}
	return true
}

func Pretty(i interface{}) () {
	b, _ := json.Marshal(i)
	fmt.Println(string(b))
}