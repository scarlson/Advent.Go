package main

import (
	"bufio"
	"log"
	"math/rand"
	"os"
	"strings"
)

var (
	lastmessage int

	dwarves   []Object
	bottle    *Object
	inventory []*Object
)

// ActSpeak displays the default message for the action
func ActSpeak(a *Action) string {
	return a.GetMessage()
}

// Toting tells us whether or not the user currently holds the object
func Toting(o *Object) bool {
	return false
}

// Here returns true if toting or if room.hasitem
func Here(o *Object) bool {
	return false
}

// At returns true if on either side of two placed objecct?
func At(o *Object) bool {
	return false
}

// Liq returns the flag for what the bottle contents are
func Liq() bool {
	// object number of liquid in bottle?
	// should be water/21  oil/22 or nothing
	return false
}

// Dark returns bool of whether the room has light or not
func Dark() bool {
	return false
}

// Pct returns true if and event happens Pct(n) of times (eg 30% change)
func Pct(n int) bool {
	return rand.Intn(100) > n-1
}

// Yea is 50/50 chance of yay or neigh
func Yea() bool {
	// 50/50 true/false response
	return rand.Int()%2 == 0
}

// GetRoom is a convenience method to return a room from the rooms map
func GetRoom(id int) *Room {
	return Rooms[id]
}

// Move returns the resulting room if the input action is performed on the input room
func Move(r *Room, a *Action) *Room {
	if _, ok := r.Connection[a.ID]; ok {
		return Rooms[r.Connection[a.ID]]
	}
	return r
}

// GetMessage is a convenience method to return a message from the message map
func GetMessage(id int) string {
	return Msgs[id]
}

// ProcessInput evaluates user input and executes the game logic
func ProcessInput(c string) string {
	i := strings.Split(c, " ")
	if len(i) == 1 {
		a := GetActionFromStr(c)
		return a.AcceptableStr[0]
	}
	return ""
}

// GetActionFromStr searches the action map for the action applied by user input string
func GetActionFromStr(c string) *Action {
	//Actions[45].AcceptableStr = []string{`NORTH`, `N`}
	for _, cmd := range Actions {
		for _, str := range cmd.AcceptableStr {
			if len(c) > 5 {
				if str == c[:5] {
					return cmd
				}
			} else {
				if str == c {
					return cmd
				}
			}
		}
	}
	return nil
}

// Adventure kickstarts the game logic
func Adventure() {
	scanner := bufio.NewScanner(os.Stdin)
	curRoom := Rooms[1]
	log.Println(GetMessage(65))
	for scanner.Scan() {
		t := scanner.Text()
		a := GetActionFromStr(strings.ToUpper(t))
		if a != nil {
			if a.GetMessage() != "" {
				log.Println(a.GetMessage())
			}
			curRoom := Move(curRoom, a)
			log.Println(curRoom.LongDesc)
			for _, o := range curRoom.Objects {
				log.Println(o.Description[000])
			}
		}
	}
	if err := scanner.Err(); err != nil {
		log.Println(os.Stderr, "read error:", err)
	}
}
