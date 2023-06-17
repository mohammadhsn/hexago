package hexago

import (
	"fmt"
	"github.com/google/uuid"
	"reflect"
	"testing"
	"time"
)

const (
	DoSthCmd         = "DoSth"
	DoSth2Cmd        = "DoSth2"
	SthHappenedEvent = "sth-happened"
)

var (
	busAudit []string
	cmdBus   CmdBus
)

type User struct {
	Identifier uuid.UUID
	events     []Event
}

func (u *User) DoSth() {
	u.events = append(u.events, NewSthHappened("Sth"))
}

func (u *User) Id() uuid.UUID {
	return u.Identifier
}

func (u *User) Pull() []Event {
	e := u.events
	u.events = make([]Event, 0)
	return e
}

type DoSth struct {
	SomeField    string
	AnOtherField int
}

type StaffRepo struct {
	seen []AggRoot
}

func (r *StaffRepo) Persist(u User) {
	r.seen = append(r.seen, &u)
}

func (r *StaffRepo) Seen() []AggRoot {
	return r.seen
}

type DoSthHandler struct {
	Repo *StaffRepo
}

func (d DoSthHandler) Handle(_ Cmd) (string, error) {
	defer func() {
		busAudit = append(busAudit, fmt.Sprintf("%T", d))
	}()

	u := User{Identifier: NewId()}
	u.DoSth()
	d.Repo.Persist(u)
	return u.Identifier.String(), nil
}

func NewDoSth() Cmd {
	return NewCmd(DoSthCmd, DoSth{
		SomeField:    "foo",
		AnOtherField: 10,
	})
}

type DoSth2 struct {
	A string
	B int
}

type DoSth2Handler struct {
}

func (c DoSth2Handler) Handle(_ Cmd) (string, error) {
	defer func() {
		busAudit = append(busAudit, fmt.Sprintf("%T", c))
	}()

	return "", nil
}

type SthHappened struct {
	Reason string
}

func NewSthHappened(reason string) Event {
	return Event{
		Identifier: NewId(),
		Name:       SthHappenedEvent,
		At:         time.Now(),
		Payload:    SthHappened{Reason: reason},
	}
}

type SthHappenedHandler struct {
}

func (s SthHappenedHandler) Handle(_ Event) {
	defer func() {
		busAudit = append(busAudit, fmt.Sprintf("%T", s))
	}()
}

type SthHappenedHandler2 struct {
}

func (s SthHappenedHandler2) Handle(_ Event) {
	defer func() {
		busAudit = append(busAudit, fmt.Sprintf("%T", s))
	}()

	cmdBus.Handle(NewCmd(DoSth2Cmd, DoSth2{A: "foo", B: 10}))
}

func createBus() {
	cmdBag := NewCmdBag()
	cmdBag.Add(DoSthCmd, func() CmdHandler {
		return DoSthHandler{Repo: &StaffRepo{}}
	})
	cmdBag.Add(DoSth2Cmd, func() CmdHandler {
		return DoSth2Handler{}
	})

	eventBag := NewEventBag()
	eventBag.Add(SthHappenedEvent, func() EventHandler {
		return SthHappenedHandler{}
	})
	eventBag.Add(SthHappenedEvent, func() EventHandler {
		return SthHappenedHandler2{}
	})

	cmdBus = &InMemoryBus{cmdBag, eventBag, make([]*Event, 0)}
}

func TestInMemoryBus_Handle(t *testing.T) {
	createBus()
	if res := cmdBus.Handle(NewDoSth()); res == "" {
		t.Error("did not expect empty string")
	}

	if len(busAudit) != 4 {
		t.Errorf("expected len to be 4, got %d", len(busAudit))
	}
}

func TestHasRepoField(t *testing.T) {
	ds := DoSth{SomeField: "foo", AnOtherField: 10}

	if hasRepoField(reflect.ValueOf(ds)) {
		t.Error("expected false, got true")
	}

	h := DoSthHandler{Repo: &StaffRepo{}}

	if !hasRepoField(reflect.ValueOf(h)) {
		t.Error("expected false, got true")
	}
}
