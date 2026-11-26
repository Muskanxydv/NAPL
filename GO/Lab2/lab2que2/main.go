package main

import (
	"fmt"
	"slices"
)

func main() {
	//SLICES
	//Taking a slice of Marks of students from the user
	var marks []int
	fmt.Print("Enter the number of students: ")
	var n int
	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		var mark int
		fmt.Printf("Enter the marks of student %d: ", i+1)
		fmt.Scan(&mark)
		marks = append(marks, mark)
	}

	//adding a new element to the slice
	marks = append(marks, 90)

	//printing the updated slice
	fmt.Println("Updated Marks after adding a new element:", marks)

	//removing an element from the slice by its index (deleting the element at index 1)
	marks = slices.Delete(marks, 1, 2)
	fmt.Println("Updated Marks after removing the element at index 1:", marks)

	//updating an element in the slice (changing the element at index 3 to 77)
	marks[3] = 77
	fmt.Println("Updated Marks after changing the element at index 3 to 77:", marks)

	//MAPS
	//Declaring a map by user
	stu_marks := map[string]int{}
	fmt.Print("Enter the number of students: ")
	var m int
	fmt.Scan(&m)
	for i := 0; i < m; i++ {
		var name string
		var mark int
		fmt.Printf("Enter the name of student %d: ", i+1)
		fmt.Scan(&name)
		fmt.Printf("Enter the marks of student %d: ", i+1)
		fmt.Scan(&mark)
		stu_marks[name] = mark
	}
	fmt.Println("Student Marks:", stu_marks)

	//inserting a new pair into the map
	stu_marks["Yuvraj"] = 90
	fmt.Println("Updated Student Marks after adding a new student:", stu_marks)

	//deleting a pair from the map (deleting "Anju")
	delete(stu_marks, "Anju")
	fmt.Println("Updated Student Marks after removing a student (Anju):", stu_marks)

	//using the lookup operation
	if marks, ok := stu_marks["Kartik"]; ok {
		fmt.Println("Marks of Kartik:", marks)
	}

	//using lookup operation for a key that does not exist to show it returns false
	v := stu_marks["Manshi"]
	v, ok := stu_marks["Manshi"]
	fmt.Println(v, ok)

}
