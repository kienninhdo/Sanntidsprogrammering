package main

import "fmt"
//var elevator Elevator, static
//var outputDevice ElevOutputDevice, static

// den skal gi ut en void __attribute__((constructor))?
func fsm_init(){
	elevator := elevator_uninitialized()

	/*
    con_load("elevator.con",
        con_val("doorOpenDuration_s", &elevator.config.doorOpenDuration_s, "%lf")
        con_enum("clearRequestVariant", &elevator.config.clearRequestVariant,
            con_match(CV_All)
            con_match(CV_InDirn)
        )
    )
	*/

	outputDevice :=elevio_getOutputDevice()

}

func setAllLights(es Elevator){ //Elevator defineres i requests.go, static

	for floor := 0; floor < N_FLOORS; floor++ { //N_FLOORS defineres annet sted
		for btn := 0; btn < N_BUTTONS; btn++ { 
			outputDevice.requestButtonLight(floor, btn, es.requests[floor][btn]) //ikke gjort om fra C, outputdevice fra elevio_outputdevice
		}
	}
}

func fsm_onInitBetweenFloors() {
	outputDevice.MotorDirection(D_Down)
	elevator.dirn := D_Down
	elevator.behaviour := EB_Moving
}

func fsm_onRequestButtonPress(btn_floor int, btn_type Button) {
	fmt.Printf("\n\n%s(%d, %s)\n", "FYLL INN FUNKSJON", btnFloor, btnType) //hvilken funksjon skal fylles inn
	elevator_print(elevator)

	switch elevator.behaviour {
	case EB_DoorOpen:
		if requests_shouldClearImmediately(elevator, btn_floor, btn_type) {
			timer_start(elevator.config.doorOpenDuration_s)
		} else {
			elevator.requests[btn_floor][btn_type] = 1
		}
	case EB_Moving:
		elevator.requests[btn_floor][btn_type] = 1
	
	case EB_Idle:
		elevator.requests[btn_floor][btn_type] = 1
		var pair DirnBehaviourPair := requests_chooseDirection(elevator)
		elevator.dirn := pair.dirn
		elevator.behaviour = pair.behaviour //skal være en del av structen
		switch pair.behaviour {
			case EB_DoorOpen:
				outputDevice.doorLight(1)
				timer_start(elevator.config.doorOpenDuration_s) //hmmm
				elevator := request_clearAtCurrentFloor(elevator)
			case EB_DoorMoving:
				outputDevice.MotorDirection(elevator.dirn)
			case EB_Idle:
				//do nothing
		}
	}

	setAllLights(elevator)
	fmt.Printf("\nNew state: \n")
	elevator_print(print)

}

func fsm_onFloorArrival(newFloor int) {
	fmt.Printf("\n\n%s(%d)\n", "FYLL INN FUNKSJON", newFloor)
	elevator_print(elevator)
	elevator.floor := newFloor
	outputDevice.floorIndicator(elevator.floor)

	switch elevator.behaviour{
		case EB_Moving:
			if(requests_shouldStop(elevator))

	}
}
