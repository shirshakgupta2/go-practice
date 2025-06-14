package main

func main() {
	
		robo := &RobotDog{Name: "Bolt"}
		performActions(robo)

	vehicles := []Vehicle{&Car{"Tesla"}, &Bike{"Yamaha"}}

	for _, v := range vehicles {
		v.Start()

		// Type assertion: special behavior for Car
		if car, ok := v.(*Car); ok {
			car.PlayMusic()
		}
	}
}
