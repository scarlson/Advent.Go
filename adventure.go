package main

import (
	"bufio"
	"log"
	"os"
	"strings"
)

func Toting(o *Object) bool {
	return false
}

func Here(o *Object) bool {
	return false
}

func At(o *Object) bool {
	return false
}

func Liq(o *Object) bool {
	return false
}

func Dark(r *Room) bool {
	return false
}

func Pct(n int) bool {
	return false
}

func Yea() bool {
	return false
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
	log.Println(cur_room.ShortDesc)
	for scanner.Scan() {
		t := scanner.Text()
		a := GetActionFromStr(strings.ToUpper(t))
		if a != nil {
			cur_room := Move(cur_room, a)
			log.Println(cur_room.ShortDesc)
		}
	}
	if err := scanner.Err(); err != nil {
		log.Println(os.Stderr, "read error:", err)
	}
}
