package arraysslicesmaps

import (
	"fmt"
)

type CourseRatingsMap map[string]float32

func (t CourseRatingsMap) ToString(){
	fmt.Println(t)
}

func Maps(){
	websites := map[string]string{
		"Brave" : "https://brave.com",
		"Amazon Web Services" : "https://aws.com",
	}

	fmt.Println(websites)
	fmt.Println(websites["Brave"])

	websites["Telegram"] = "https://telegram.org"
	fmt.Println(websites)
	websites["Amazon Web Services"] = "https://overridden.com"
	fmt.Println(websites)
	delete(websites, "Amazon Web Services")
	fmt.Println(websites)
}


func UsingMake(){
	//this is just memory efficient from GO pov
	userNames := make([]string, 3, 10) //data type and length of the slice and capacity of the slice
	userNames = append(userNames, "Oscar") //we append data after these 3 element defined above
	userNames = append(userNames, "Kimi")
	fmt.Println(userNames)
	userNames[0] = "Max"
	fmt.Println(userNames)



	courses := make(CourseRatingsMap, 3) //3 - length of map

	courses["ReactJS"] = 4.3
	courses["GO"] = 3.5
	courses["Kubernets"] = 3.9
	courses.ToString()


	for index, userName := range userNames {
		fmt.Println(index, userName)
	}

	for key, courseRating := range courses {
		fmt.Println(key, courseRating)
	}
}