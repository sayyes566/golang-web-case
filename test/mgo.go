package main

import (
        "fmt"
		"log"
        "gopkg.in/mgo.v2"
		"gopkg.in/mgo.v2/bson"
	//	"time"
)

/*==========Person Management=================*/
type Person struct {
	/*
	student/ teacher/ assistant
	*/
	Name string //[uniqle]
	Phone_Home string // Nil
	Phone_Cell string //Nil
	EMAIL string // Nil
	Identity string // student/ teacher/ assistant
	Classes []string //[A,B..] or nil
	Salary_HR int //[same as Salary table] for Teacher and Assistant only
	Student_Parents_Contect string
	Note string
	UpdateTime  string // Array timestamp
	Delete_Record bool // 1 = delete but is not in the trash
}



// {
	
// 	person["Name"] 
// 	person["Phone_Home"]  
// 	person["Phone_Cell"] 
// 	person["EMAIL"]
// 	person["Identity"] 
// 	person["Classes"] 
// 	person["Salary_HR"] 
// 	person["Student_Parents_Contect"] 
// 	person["Note"] 
// }



/*==========Class Management=================*/
type Class struct {
	Class string //[uniqle] Name of class ex: English A, English B
	Date string // Fixed date ex: friday
	Time_HR string // Fixed time ex: 1430 -> 14:30
	Duration_HR string // Fixed duration ex: 2.5 -> 2.5 hours
	Open_Date string // when is this class's first day ?
	Teacher string // Name of a teacher
	Num_Students int // number of students
	Course_Type string // ex: math or english
	//student
	Students []string // Name of students: [..,..]
	Pay_Student int // cost of a class each student
	UpdateTime  string // Array timestamp
	Delete_Record bool // 1 = delete but is not in the trash
	
}

/*==========Payment Management=================*/
type Student_Payment struct{
	/*
		A Student : Class -> 1 : N
	*/
	Student_Class string // for unique key: student name + class name
	Student string
	Class string 
	Date_Has_Paid string
	Date_Next_Pay string // When does a student have to pay next time?
	Date_Start_Study string
	Times_Study_Pay int // How many study times does a student have to pay?
	Pay_Total int //A student has paid ? NTD.
	Pay_Next int // Next time a student need to pay ? NTD.
	UpdateTime string // Array timestamp
	Delete_Record bool // 1 = delete but is not in the trash
}

type Assistant_Salary struct{
	Assistant string
	HR_Work_No_Gain int //How long (work Hours)  hasn't gaiven to the assistant.
	HR_Work_Total int
	//@problem: forget to check 
	HR_Salary int
	UpdateTime string // Array timestamp
	Delete_Record bool // 1 = delete but is not in the trash
}

type Teacher_Salary struct{
	/*
		A Teacher : Class -> 1 : N
	*/
	Teacher_Class string // for unique key
	Teacher string
	Class string
	Date_Has_Earn string //YYYYMMdd
	Date_Next_Earn string  // YYYYMMdd
	Date_Start_Work string // YYYYMMdd
	Times_Work_Earn int //  how many times of teach will give a salary?
	Earn_Total int //NTD
	Earn_Next int //NTD
	HR_Salary int // salary of each hr
	Earn_Type string // by HR_Salary or Earn_Next
	UpdateTime string // Array timestamp
	Delete_Record bool // 1 = delete but is not in the trash
	
	
}

/*==========Time Management=================*/
type Person_Class struct {
	/*
		One day a Teacher or a Student : Class -> 1 : N
	*/

	Identity string //student or teacher
	Name string
	Class string
	Date string //YYYYMMdd
	Time string //HHmm
	Cost_Other int //ex printed paper
	Has_Make_UP bool // if this person want to make up a class than show 1 equeal yes 
	Date_Make_UP string // YYYMMdd
	Time_Make_UP string
	Attend bool // 1 yes 0 no
	UpdateTime string // Array timestamp
	Delete_Record bool // 1 = delete but is not in the trash

}
type Assistant_Date struct{
	Assistant string
	Date_Attend string
	Punch_Start_Time string //recently check in to work 
	Punch_End_Time string //recently check in to work 
	UpdateTime string // Array timestamp
	Delete_Record bool // 1 = delete but is not in the trash
}

type Histry_Payment struct{
	Identity string //student or teacher
	Person string //name
	InOut bool // student = in = 1 , teacher, assistant = out = 0
	Money bool
	Date string 
	UpdateTime string
	Delete_Record bool
}
/*

		c := session.DB("test").C("people")
		// Index
		index := mgo.Index{
			Key:        []string{"name"},
			Unique:     true,
			DropDups:   true,
			Background: true,
			Sparse:     true,
		}
		err = c.EnsureIndex(index)
		if err != nil {
			panic(err)
		}

		p := new(Person)
		p.Name = "Eslie"
		p.Phone_Home = "02123456789"
		p.Phone_Cell = "0937818336"
		p.EMAIL = "Eslie@gmail.com"
		p.Identity = "Teacher"
		p.Classes = []string{"JP_GOGO4", "JP_GOGO5"}
		p.Salary_HR = 500
		//p.Note = "Normal!"
		//p.Student_Parents_Contect = "09123456789"
		p.UpdateTime = time.Now().Format(layout)

		err = c.Insert(p)
        if err != nil {
                log.Fatal(err)
        }

        result := Person{}
        err = c.Find(bson.M{"Name": "Eslie"}).One(&result)
        if err != nil {
                log.Fatal(err)
        }
*/


/*
index := mgo.Index{
			Key:        []string{"class"},
			Unique:     true,
			DropDups:   true,
			Background: true,
			Sparse:     true,
		}
		err = c.EnsureIndex(index)
		if err != nil {
			panic(err)
		}

		p := new(Class)
		p.Class = "ENG"
		p.Date = "Saturday"
		p.Time_HR = "1530"
		p.Duration_HR = "1" 
		p.Open_Date = "20150305"
		p.Teacher = "Bill"
		p.Num_Students = 1
		p.Course_Type = "English"
		//student
		p.Students = []string{"Kristen"}
		p.Pay_Student = 900
		p.UpdateTime  = time.Now().Format(layout)
*/
/*
c := session.DB("test").C("student_payment")
		// Index
		index := mgo.Index{
			Key:        []string{"student_class"},
			Unique:     true,
			DropDups:   true,
			Background: true,
			Sparse:     true,
		}
		err = c.EnsureIndex(index)
		if err != nil {
			panic(err)
		}

		p := new(Student_Payment)
		p.Student_Class = "KristenENG"
		p.Class = "ENG"
		p.Student = "Kristen"
		p.Date_Has_Paid = "0708"
		p.Date_Next_Pay = "0816"
		p.Date_Start_Study = "20170401"
		p.Times_Study_Pay = 11 
		p.Pay_Total = 6600
		p.Pay_Next = 3300
		
		p.UpdateTime  = time.Now().Format(layout)
*/
/*
c := session.DB("test").C("assistant_salary")
		// Index
		index := mgo.Index{
			Key:        []string{"assistant"},
			Unique:     true,
			DropDups:   true,
			Background: true,
			Sparse:     true,
		}
		err = c.EnsureIndex(index)
		if err != nil {
			panic(err)
		}

		p := new(Assistant_Salary)
		p.Assistant = "John"
		p.HR_Work_No_Gain = 20
		p.HR_Work_Total = 50
		p.HR_Salary = 115
		p.UpdateTime  = time.Now().Format(layout)
*/


/*
c := session.DB("test").C("teacher_salary")
		// Index
		index := mgo.Index{
			Key:        []string{"teacher_class"},
			Unique:     true,
			DropDups:   true,
			Background: true,
			Sparse:     true,
		}
		err = c.EnsureIndex(index)
		if err != nil {
			panic(err)
		}

		p := new(Teacher_Salary)
		p.Teacher_Class = "Eslie_JP_GOGO4"
		p.Teacher = "Eslie"
		p.Class = "JP_GOGO4"
		p.Date_Has_Earn = "20170909"
		p.Date_Next_Earn = "20170930"
		p.Date_Start_Work = "20150101"
		p.Times_Work_Earn = 4
		p.Earn_Total = 33000
		p.Earn_Next = 1200
		p.HR_Salary = 300
		p.Earn_Type = "Earn_Next"
		p.UpdateTime  = time.Now().Format(layout)

*/

/*
c := session.DB("test").C("person_class")
p := new(Person_Class)

		p.Identity = "Teacher"
		p.Name = "Bill"
		p.Class = "ENG"
		p.Date = "20170916"
		p.Time = "1030"
		//p.Cost_Other = 0
		p.Attend = true
		p.UpdateTime  = time.Now().Format(layout)
		
*/
/*
c := session.DB("test").C("assistant_date")
	p := new(Assistant_Date)

		p.Assistant = "John"
		p.Date_Attend = "20170909"
		p.Punch_Start_Time = "0800"
		p.Punch_End_Time = "1700"
		
		p.UpdateTime  = time.Now().Format(layout)
*/
func main() {
		const layout = "Mon Jan _2 15:04:05 2006"
        session, err := mgo.Dial("127.0.0.1")
        if err != nil {
                panic(err)
        }
        defer session.Close()

        // Optional. Switch the session to a monotonic behavior.
        session.SetMode(mgo.Monotonic, true)
		c := session.DB("test").C("people")
		// // Index
		// index := mgo.Index{
		// 	Key:        []string{"teacher_class"},
		// 	Unique:     true,
		// 	DropDups:   true,
		// 	Background: true,
		// 	Sparse:     true,
		// }
		// err = c.EnsureIndex(index)
		// if err != nil {
		// 	panic(err)
		// }

		// p := new(Assistant_Date)
		// p.Assistant = "John"
		// p.Date_Attend = "20170909"
		// p.Punch_Start_Time = "0800"
		// p.Punch_End_Time = "1700"
		// p.UpdateTime  = time.Now().Format(layout)
		
		
		
		// err = c.Insert(p)
        // if err != nil {
        //         log.Fatal(err)
        // }

        result := Person{}
        err = c.Find(bson.M{"name": "Eslie"}).One(&result)
        if err != nil {
                log.Fatal(err)
        }
		fmt.Println("Name:", result.Name)
         fmt.Println("Phone:", result.Phone_Cell)
}