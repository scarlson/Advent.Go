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
)

func Toting(o *Object) bool {
	// whether or not the user currently holds the object
	return false
}

func Here(o *Object) bool {
	// true if toting or if room.hasitem
	return false
}

func At(o *Object) bool {
	// true if on either side of two placed objecct?
	return false
}

func Liq() bool {
	// object number of liquid in bottle?
	// should be water/21  oil/22 or nothing
	return false
}

func Dark() bool {
	// true if room is dark
	// why is this not a property of room
	/* light rooms =
	0   1   2   3   4   5   6   7   8   9   10
	0   100 115 116 126
	*/

	return false
}

func Pct(n int) bool {
	return rand.Intn(100) > n-1
}

func Yea() bool {
	// 50/50 true/false response
	return rand.Int()%2 == 0
}

func GetRoom(id int) *Room {
	return Rooms[id]
}

func Move(r *Room, a *Action) *Room {
	if _, ok := r.Connection[a.Id]; ok {
		return Rooms[r.Connection[a.Id]]
	}
	return r
}

func GetMessage(id int) string {
	return Msgs[id]
}

func ProcessInput(c string) string {
	i := strings.Split(c, " ")
	if len(i) == 1 {
		a := GetActionFromStr(c)
		return a.AcceptableStr[0]
	}
	return ""
}

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

func Adventure() {
	scanner := bufio.NewScanner(os.Stdin)
	cur_room := Rooms[1]
	log.Println(GetMessage(65))
	for scanner.Scan() {
		t := scanner.Text()
		a := GetActionFromStr(strings.ToUpper(t))
		if a != nil {
			if a.GetMessage() != "" {
				log.Println(a.GetMessage())
			}
			cur_room := Move(cur_room, a)
			log.Println(cur_room.LongDesc)
			for _, o := range cur_room.Objects {
				log.Println(o.Description[000])
			}
		}
	}
	if err := scanner.Err(); err != nil {
		log.Println(os.Stderr, "read error:", err)
	}
}
