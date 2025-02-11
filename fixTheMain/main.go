package main

import "github.com/01-edu/z01"

type DoorState bool

const (
	OPEN  DoorState = true
	CLOSE DoorState = false
)

type Door struct {
	state DoorState
}

func PrintStr(s string) {
	for _, r := range s {
		z01.PrintRune(r)
	}
	z01.PrintRune('\n')
}

func OpenDoor(ptrDoor *Door) {
	PrintStr("Door Opening...")
	ptrDoor.state = OPEN
}

func CloseDoor(ptrDoor *Door) bool {
	PrintStr("Door Closing...")
	ptrDoor.state = CLOSE
	return true
}

func IsDoorOpen(ptrDoor *Door) bool {
	PrintStr("Is the Door opened ?")
	return ptrDoor.state == OPEN
}

func IsDoorClose(ptrDoor *Door) bool {
	PrintStr("Is the Door closed ?")
	return ptrDoor.state == CLOSE
}

func main() {
	door := &Door{}

	OpenDoor(door)
	if IsDoorClose(door) {
		OpenDoor(door)
	}
	if IsDoorOpen(door) {
		CloseDoor(door)
	}
	if door.state == OPEN {
		CloseDoor(door)
	}
}
