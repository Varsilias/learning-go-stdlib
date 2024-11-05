package main

import (
	"fmt"
)

type Person struct {
	Name string
	Age  int
	City string
}

// TODO: Implement String() method for Person
// Should return formatted string with all fields

func (p Person) String() string {
	return fmt.Sprintf("My name is %s and I am %d years old, I live in %v", p.Name, p.Age, p.City)
}

type ValidationError struct {
	Field        string
	ErrorMessage string
}

// TODO: Implement Error() method
// Should return properly formatted error message

func (v ValidationError) Error() error {
	return fmt.Errorf("Error occurred in field %s the Type is %T, Full error is %v", v.Field, v.Field, v.ErrorMessage)
}

// func quiz() {
// 	var dayOfWeek string
// 	// TODO: Ask user questions using fmt.Print
// 	fmt.Print("What day is it?: ")

// 	// Read answers using fmt.Scan
// 	fmt.Scan(&dayOfWeek)

// 	// Format and display results
// 	fmt.Println(dayOfWeek)
// }

type TableData struct {
	Headers []string
	Rows    [][]string
}

// TODO: Implement function to print formatted table
// Use fmt.Printf with width specifiers

// Width Specifiers

// Width is the number of characters to output, if the number is more than
// the original string character, it will add extra spaces up to the surplus number
// to the original characters. The spaces will be to the left by default, if you put a "-"
// sign before the "WIDTH Specifier", the spaces will be added to the right

/* E.g fmt.Printf("|%10v|", "Daniel") // |    Daniel|
   E.g fmt.Printf("|%-10v|", "Daniel") // |Daniel    |
*/

// Precision Specifier
// This is the part that comes after the DOT in the formatter.
// The general idea behind "PRECISION Specifyier" is that it limits the number of character
// to output from the original character, if the number is larger than the original character
// length, it will print the original character as is
// But how it is done is dependent on the DATA TYPE

/*
	E.g fmt.Printf("|%.2s|", "DANIEL") // |DA|
	E.g fmt.Printf("|%.5s|", "DANIEL") // |DANIE|
	E.g fmt.Printf("|%.3f|", 3.1423455666) // |3.142|
*/

func (t TableData) Print() {

	// fmt.Printf("|%-10s", t.Headers[0])
	// fmt.Printf("|%-10s", t.Headers[1])
	// fmt.Printf("|%-10s", t.Headers[2])
	// fmt.Printf("|%-10s\n", t.Headers[3])

	for i := 0; i < len(t.Headers); i++ {
		if i == len(t.Headers)-1 {
			fmt.Printf("|%-30s\n", t.Headers[i])
			break
		}
		fmt.Printf("|%-30s", t.Headers[i])
	}

	for i := 0; i < len(t.Rows); i++ {
		for j := 0; j < len(t.Rows[i]); j++ {
			if j == len(t.Rows[i])-1 {
				fmt.Printf("|%-30s\n", t.Rows[i][j])
				continue
			}
			fmt.Printf("|%-30s", t.Rows[i][j])
		}
	}
}

func main() {
	// var name string
	// fmt.Print("Enter your name: ")
	// fmt.Scanln(&name)

	// fmt.Println((name))

	// name := "Daniel"
	// age := 25
	// shouldLockIn := true

	// type Animal struct {
	// 	Sound     string
	// 	Type      string
	// 	Dimension int64
	// }

	// %v is to print the value in its default format
	// what is the default format for most Go types
	// default format is just the value as exactly defined

	// fmt.Printf("Name: %v \n", name)
	// fmt.Printf("Name Type: %T \n", name)
	// fmt.Printf("Age: %v \n", age)
	// fmt.Printf("Boolean: %v \n", shouldLockIn)
	// fmt.Printf("Struct: %v \n", Animal{Sound: "bark", Type: "Mammal", Dimension: 600})
	// fmt.Printf("Struct With Field Names: %+v", Animal{Sound: "bark", Type: "Mammal", Dimension: 600})

	// var test []byte

	// fmt.Println(fmt.Append(test, "hello"))

	// const name, id = "bueller", 17
	// err := fmt.Errorf("user %q (id %d) not found", name, id)
	// fmt.Println(err.Error())

	// const name, age = "Kim", 22
	// s := fmt.Sprint(name, " is ", age, " years old.\n")

	// io.WriteString(os.Stdout, s) // Ignoring error for simplicity.

	// TODO: Create variables of different types (int, float, string, bool)
	// Format them using different verbs (%v, %T, %d, %f, %s)
	// Print them in different ways

	// name := "Daniel"
	// age := 25
	// percentage := 32.687562028
	// shouldLockIn := true

	// // Strings
	// fmt.Printf("String in default format: %v\n", name)
	// fmt.Printf("String as Type: %T\n", name)
	// fmt.Printf("String as String Slice %s\n", name)

	// // Ints
	// fmt.Printf("Int in default format: %v\n", age)
	// fmt.Printf("Int as Type: %v\n", age)
	// fmt.Printf("Int in Decimal format: %d\n", age)

	// //Bool
	// fmt.Printf("Bool in default format: %v\n", shouldLockIn)
	// fmt.Printf("Bool as Type: %T\n", shouldLockIn)
	// fmt.Printf("Bool as Boolean format: %t\n", shouldLockIn)

	// //Float
	// fmt.Printf("Float in default format: %v\n", percentage)
	// fmt.Printf("Float as Type: %T\n", percentage)
	// fmt.Printf("Float in Float format: %f\n", percentage)

	// TODO: Use the implemented String Method
	// var me = Person{Name: "Daniel Okoronkwo", Age: 25, City: "Lagos"}
	// sopuu := Person{Name: "Sopuluchukwu Nnacheta", Age: 28, City: "Lagos"}
	// shazzar := Person{Name: "Daniel Oguejiofor", Age: 25, City: "Ibadan"}

	// fmt.Println(me.String())
	// fmt.Println(sopuu.String())
	// fmt.Println(shazzar.String())

	//TODO: Use the implemented Error method

	// var stringError = ValidationError{Field: "String", ErrorMessage: "Could not validate string"}
	// floatError := ValidationError{"Float", "Could not validate float"}

	// fmt.Println(stringError.Error().Error())
	// fmt.Println(floatError.Error())

	// quiz()

	// fmt.Printf("|%10v|\n", "Daniel")
	// fmt.Printf("|%-10v|\n", "Daniel")

	// fmt.Printf("|%.10s|\n", "DANIEL")
	// fmt.Printf("|%.3f|", 3.1423455666)

	var tableData = TableData{
		Headers: []string{"Name", "Age", "City", "Account Balance"},
		Rows: [][]string{
			{"Daniel Okoronkwo", "25", "Lagos", "1,000,000"},
			{"Sopuluchukwu Nnacheta", "28", "Lagos", "1,500,000"},
			{"Daniel Oguejiofor", "25", "Ibadan", "2,000,000"},
			{"Cordelia Ukpai", "26", "Abakiliki", "2,500,000"},
			{"Elias Emmanuel", "30", "Abuja", "3,000,000"},
			{"Promise Nnacheta", "35", "Lagos", "3,500,000"},
		},
	}
	tableData.Print()
}
