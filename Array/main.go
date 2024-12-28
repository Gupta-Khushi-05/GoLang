package main

import "fmt"

func main(){
	hobby := [3]string{"art", "paint", "craft"}
  fmt.Println(hobby)
	fmt.Println(hobby[0])
	fmt.Println(hobby[1:])
  slice := hobby[:2] 
	fmt.Println(slice) 
	fmt.Println(hobby[0:2])
	fmt.Println(cap(slice))
	fmt.Println(slice[1:3])

	var goal = []string{"success"}
	fmt.Println(goal)
	goal = append(goal, "happy")
	fmt.Println(goal)
	goal[1] = "paint"
	goal = append(goal, "fly")
	fmt.Println(goal)


	type product struct{
		title string
		id int
		price float64
	}

	userProduct := []product{{"pen", 2, 2.5}, {"chart", 4, 5.6}}
  fmt.Println(userProduct)
  newProduct := product{"pencil", 7, 2.6}
	userProduct = append(userProduct, newProduct)
	fmt.Println(userProduct)



}