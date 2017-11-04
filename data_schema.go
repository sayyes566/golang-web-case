package main

import (
	// "fmt"

	"fmt"
	"log"
	"reflect"
	"strconv"
	"strings"
	"time"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

var (
	con_ip      = "127.0.0.1"
	con_db      = "test"
	table_count = 0
)

const (
	con_layout = "Mon Jan _2 15:04:05 2006"
)

// type Model struct {
// }

/*==========Person Management=================*/
type Person struct {
	/*
		student/ teacher/ assistant
	*/
	Uid                     int    `json: "uid" bson:"uid"`
	Name                    string `json: "Name" bson:"Name"`             //[uniqle]
	Phone_Home              string `json: "phone_home" bson:"phone_home"` // Nil
	Phone_Cell              string `json: "phone_cell" bson:"phone_cell"` //Nil
	EMAIL                   string `json: "email" bson:"email"`           // Nil
	Identity                string `json: "identity" bson:"identity"`     // student/ teacher/ assistant
	Open_Date               string `json: "open_date" bson:"open_date"`
	Classes                 string `json: "classes" bson:"classes"`     //[A,B..] or nil
	Salary_HR_Int           int    `json: "salary_hr" bson:"salary_hr"` //[same as Salary table] for Teacher and Assistant only
	Salary_HR               string `json: "salary_hr" bson:"salary_hr"`
	Student_Parents_Contect string `json: "student_parents_contect" bson:"student_parents_contect"`
	Note                    string `json: "note" bson:"note"`
	UpdateTime              string `json: "updatetime" bson:"updatetime"`         // Array timestamp
	Delete_Record           bool   `json: "uidelete_record" bson:"delete_record"` // 1 = delete but is not in the trash
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
	Cid          int    `json: "cid" bson:"cid"`
	Class_Name   string `json: "class_name" bson:"class_name"`     //chinese name
	Class        string `json: "class" bson:"class"`               //[uniqle] Name of class ex: English A, English B
	Day          string `json: "day" bson:"day"`                   // Fixed day ex: friday
	Time_HR      string `json: "time_hr" bson:"time_hr"`           // Fixed time ex: 1430 -> 14:30
	Duration_HR  string `json: "duration_hr" bson:"duration_hr"`   // Fixed duration ex: 2.5 -> 2.5 hours
	Open_Date    string `json: "open_date" bson:"open_date"`       // when is this class's first day ?
	Teacher      string `json: "teacher" bson:"teacher"`           // Name of a teacher
	Num_Students int    `json: "num_students" bson:"num_students"` // number of students
	Course_Type  string `json: "course_type" bson:"course_type"`   // ex: math or english
	//student
	Students              string `json: "students" bson:"students"`                           //a,b,c
	Cost_Each_Student_Str string `json: "cost_each_student_str" bson:"cost_each_student_str"` // Name of students: [..,..]
	Cost_Each_Student     int    `json: "cost_each_student" bson:"cost_each_student"`
	Charge_Times_Int      int    `json: "charge_times" bson:"charge_times"` // cost of a class each student
	Charge_Times          string
	UpdateTime            string `json: "updatetime" bson:"updatetime"`       // Array timestamp
	Delete_Record         bool   `json: "delete_record" bson:"delete_record"` // 1 = delete but is not in the trash

}

/*==========Payment Management=================*/
type Student_Payment struct {
	/*
		A Student : Class -> 1 : N
	*/
	SCid             int    `json: "scid" bson:"scid"`
	Student_Class    string // for unique key: student name + class name
	Student          string
	Class            string
	Date_Has_Paid    string
	Date_Next_Pay    string // When does a student have to pay next time?
	Date_Start_Study string
	Times_Study_Pay  int    // How many study times does a student have to pay?
	Pay_Total        int    //A student has paid ? NTD.
	Pay_Next         int    // Next time a student need to pay ? NTD.
	UpdateTime       string // Array timestamp
	Delete_Record    bool   // 1 = delete but is not in the trash
}

type Assistant_Salary struct {
	ASid            int `json: "asid" bson:"asid"`
	Assistant       string
	HR_Work_No_Gain int //How long (work Hours)  hasn't gaiven to the assistant.
	HR_Work_Total   int
	//@problem: forget to check
	HR_Salary     int
	UpdateTime    string // Array timestamp
	Delete_Record bool   // 1 = delete but is not in the trash
}

type Teacher_Salary struct {
	/*
		A Teacher : Class -> 1 : N
	*/
	TCid            int    `json: "tcid" bson:"tcid"`
	Teacher_Class   string // for unique key
	Teacher         string
	Class           string
	Date_Has_Earn   string //YYYYMMdd
	Date_Next_Earn  string // YYYYMMdd
	Date_Start_Work string // YYYYMMdd
	Times_Work_Earn int    //  how many times of teach will give a salary?
	Earn_Total      int    //NTD
	Earn_Next       int    //NTD
	HR_Salary       int    // salary of each hr
	Earn_Type       string // by HR_Salary or Earn_Next
	UpdateTime      string // Array timestamp
	Delete_Record   bool   // 1 = delete but is not in the trash

}

/*==========Time Management=================*/
type Person_Class struct {
	/*
		attend record
			One day a Teacher or a Student : Class -> 1 : N
	*/
	PCid          int    `json: "pcid" bson:"pcid"`
	Identity      string //student or teacher
	Name          string
	Class         string
	Date          string //YYYYMMdd
	Time          string //HHmm
	Cost_Other    int    //ex printed paper
	Has_Make_UP   bool   // if this person want to make up a class than show 1 equeal yes
	Date_Make_UP  string // YYYMMdd
	Time_Make_UP  string
	Attend        bool   // 1 yes 0 no
	UpdateTime    string // Array timestamp
	Delete_Record bool   // 1 = delete but is not in the trash

}
type Assistant_Date struct {
	// work record
	ADid             int `json: "adid" bson:"adid"`
	Assistant        string
	Date_Attend      string
	Punch_Start_Time string //recently check in to work
	Punch_End_Time   string //recently check in to work
	UpdateTime       string // Array timestamp
	Delete_Record    bool   // 1 = delete but is not in the trash
}

type Histry_Payment struct {
	HPid          int    `json: "hpid" bson:"hpid"`
	Identity      string //student or teacher
	Person        string //name
	InOut         bool   // student = in = 1 , teacher, assistant = out = 0
	Money         bool
	Date          string
	UpdateTime    string
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

func Data_set_index(c *mgo.Collection, arr_keys []string, Unique, DropDups, Background, Sparse bool) error {
	index := mgo.Index{
		// Key:        []string{"teacher_class"},
		Key:        arr_keys,
		Unique:     Unique,
		DropDups:   DropDups,
		Background: Background,
		Sparse:     Sparse,
	}
	err := c.EnsureIndex(index)
	if err != nil {
		log.Fatal("Data_set_index: " + err.Error())
		//panic(err)
	}
	return err
}

func Data_connect_mgo(serverIP string, db string, collection string) (*mgo.Collection, *mgo.Session) {
	session, err := mgo.Dial(serverIP)
	if err != nil {
		log.Fatal("Data_connect_mgo: " + err.Error())
		//panic(err)
	}

	session.SetMode(mgo.Monotonic, true)
	c := session.DB(db).C(collection)
	return c, session
}

// func Data_find_person_one_field(c *mgo.Collection, M bson.M, one_field string) []Person {
// 	person := []Person{}
// 	//err := c.Find(M).One(&result)
// 	err := c.Find(M).Select(bson.M{one_field: 1}).All(&person)
// 	if err != nil {
// 		log.Fatal("Data_find_person: " + err.Error())
// 	}
// 	table_count = len(person)
// 	return person
// }
func Data_find_person(c *mgo.Collection, M bson.M) []Person {
	person := []Person{}
	//err := c.Find(M).One(&result)
	err := c.Find(M).All(&person)
	if err != nil {
		log.Fatal("Data_find_person: " + err.Error())
	}
	table_count = len(person)
	return person
}

func Data_find_class(c *mgo.Collection, M bson.M) []Class {
	class_s := []Class{}
	err := c.Find(M).All(&class_s)
	if err != nil {
		log.Fatal("Data_find_class: " + err.Error())
	}
	table_count = len(class_s)
	return class_s
}

func Data_remove_Person(s []Person, identity string) string {
	//p := Person{}
	fmt.Println("===dmp")
	fmt.Println(s)

	c, session := Data_connect_mgo(con_ip, con_db, "people")
	defer session.Close()

	for _, ss := range s {
		fmt.Println(&ss.Name)
		fmt.Println(ss.Name)
		fmt.Println("identity")
		fmt.Println(identity)
		//suport to remove multiple records
		colQuerier := bson.M{
			strings.ToLower("Name"): ss.Name,
		}
		update_field := bson.M{
			strings.ToLower("Identity"):      identity,
			strings.ToLower("UpdateTime"):    time.Now().Format(con_layout),
			strings.ToLower("Delete_Record"): true,
		}
		change := bson.M{"$set": update_field}
		fmt.Println(colQuerier)
		fmt.Println(update_field)
		fmt.Println(change)
		err := c.Update(colQuerier, change)
		if err != nil {
			fmt.Println(err.Error())
			log.Fatal("Data_remove_Person: " + err.Error())
		}

	}

	return "success"
}

func Data_remove_Class(s []Class) string {
	//p := Person{}
	fmt.Println("===dmc")
	fmt.Println(s)

	c, session := Data_connect_mgo(con_ip, con_db, "class")
	defer session.Close()
	for _, ss := range s {
		//suport to remove multiple records
		colQuerier := bson.M{
			strings.ToLower("Class"): ss.Class,
		}
		update_field := bson.M{
			strings.ToLower("UpdateTime"):    time.Now().Format(con_layout),
			strings.ToLower("Delete_Record"): true,
		}
		change := bson.M{"$set": update_field}
		fmt.Println(colQuerier)
		fmt.Println(update_field)
		fmt.Println(change)
		err := c.Update(colQuerier, change)
		if err != nil {
			fmt.Println(err.Error())
			log.Fatal("Data_remove_Class: " + err.Error())
		}

	}

	return "success"
}

func Data_update_Class(s Class) string {
	cost_int := 0
	//charge_times := 0
	c, session := Data_connect_mgo(con_ip, con_db, "class")
	defer session.Close()
	if s.Cost_Each_Student_Str != "" {
		cost, err_conv := strconv.Atoi(s.Cost_Each_Student_Str)
		if err_conv != nil {
			cost = 0
			//log.Fatal("Data_find_Person_One: err_conv" + err.Error())
		}
		cost_int = cost
	}

	// if s.Charge_Times != "" {
	// 	intt, err_conv := strconv.Atoi(s.Charge_Times)
	// 	if err_conv != nil {
	// 		intt = 0
	// 		//log.Fatal("Data_find_Person_One: err_conv" + err.Error())
	// 	}
	// 	charge_times = cost
	// }
	//[]string{"Kristen"}

	// Update
	colQuerier := bson.M{
		strings.ToLower("Class"): s.Class,
	}
	update_field := bson.M{
		//"Uid": index_last,
		//Name:                    s.Name,
		strings.ToLower("Class_Name"):            s.Class_Name,
		strings.ToLower("Day"):                   s.Day,
		strings.ToLower("Time_HR"):               s.Time_HR,
		strings.ToLower("Duration_HR"):           s.Duration_HR,
		strings.ToLower("Open_Date"):             s.Open_Date,
		strings.ToLower("Teacher"):               s.Teacher,
		strings.ToLower("Course_Type"):           s.Course_Type,
		strings.ToLower("Students"):              s.Students,
		strings.ToLower("Num_Students"):          len(strings.Split(s.Students, ",")),
		strings.ToLower("Charge_Times"):          s.Charge_Times,
		strings.ToLower("Cost_Each_Student_Str"): s.Cost_Each_Student_Str,
		strings.ToLower("Cost_Each_Student"):     cost_int,
		strings.ToLower("UpdateTime"):            time.Now().Format(con_layout),
		strings.ToLower("Delete_Record"):         false,
	}
	change := bson.M{"$set": update_field}
	fmt.Println(colQuerier)
	fmt.Println(update_field)
	fmt.Println(change)
	err := c.Update(colQuerier, change)
	if err != nil {
		fmt.Println(err.Error())
		return ("Warning: System Error. DUC")
		//log.Fatal("Data_update_Person: " + err.Error())
	}
	return "success"
}

func Data_update_Person(s Person, identity string) string {
	//p := Person{}
	salary_int := 0
	c, session := Data_connect_mgo(con_ip, con_db, "people")
	defer session.Close()
	if s.Salary_HR != "" {
		sa, err_conv := strconv.Atoi(s.Salary_HR)
		if err_conv != nil {
			sa = 0
			//log.Fatal("Data_find_Person_One: err_conv" + err.Error())
		}
		salary_int = sa
	}
	// Update
	colQuerier := bson.M{
		strings.ToLower("Name"): s.Name,
	}
	update_field := bson.M{
		//"Uid": index_last,
		//Name:                    s.Name,
		strings.ToLower("Phone_Home"):              s.Phone_Home,
		strings.ToLower("Phone_Cell"):              s.Phone_Cell,
		strings.ToLower("EMAIL"):                   s.EMAIL,
		strings.ToLower("Identity"):                identity,
		strings.ToLower("Open_Date"):               s.Open_Date,
		strings.ToLower("Classes"):                 s.Classes,
		strings.ToLower("Salary_HR"):               s.Salary_HR,
		strings.ToLower("Salary_HR_Int"):           salary_int,
		strings.ToLower("Student_Parents_Contect"): s.Student_Parents_Contect,
		strings.ToLower("Note"):                    s.Note,
		strings.ToLower("UpdateTime"):              time.Now().Format(con_layout),
		strings.ToLower("Delete_Record"):           false,
	}
	change := bson.M{"$set": update_field}
	fmt.Println(colQuerier)
	fmt.Println(update_field)
	fmt.Println(change)
	err := c.Update(colQuerier, change)
	if err != nil {
		fmt.Println(err.Error())
		return ("Warning: System Error. DUP")
		//log.Fatal("Data_update_Person: " + err.Error())
	}
	return "success"
}
func Data_insert_Person(s Person, identity string) string {
	salary_int := 0
	fmt.Println(s.Name)
	p := Person{}
	c, session := Data_connect_mgo(con_ip, con_db, "people")
	defer session.Close()
	p = Data_find_Person_Last_Uid("-uid")
	index_last := p.Uid
	// fmt.Println("index_last")
	// fmt.Println(index_last)
	//index, err := strconv.Atoi(last_index)
	//https://github.com/fatih/structs
	if s.Salary_HR != "" {
		sa, err_conv := strconv.Atoi(s.Salary_HR)
		if err_conv != nil {
			sa = 0
			//log.Fatal("Data_find_Person_One: err_conv" + err.Error())
		}
		salary_int = sa
	}

	index_last += 1
	// fmt.Println(index_last)
	err := c.Insert(
		&Person{
			Uid:                     index_last,
			Name:                    s.Name,
			Phone_Home:              s.Phone_Home,
			Phone_Cell:              s.Phone_Cell,
			EMAIL:                   s.EMAIL,
			Identity:                identity,
			Open_Date:               s.Open_Date,
			Classes:                 s.Classes,
			Salary_HR:               s.Salary_HR,
			Salary_HR_Int:           salary_int,
			Student_Parents_Contect: s.Student_Parents_Contect,
			Note:          s.Note,
			UpdateTime:    time.Now().Format(con_layout),
			Delete_Record: false,
		})
	if err != nil {
		str_error := err.Error()
		if strings.Contains(str_error, "duplicate key") {
			return ("Warning: Username has already been taken. Please input another Name! (ex: John11)")
		} else {
			fmt.Println(str_error)
			//log.Fatal("Data_insert_Person: " + str_error)
			return "fail"
		}

	}
	return "success"
	//m := structs.Map(server)

}

func Data_insert_Class(s Class) string {

	cost_int := 0
	charge_times := 0
	//fmt.Println(s.Name)
	p := Class{}
	c, session := Data_connect_mgo(con_ip, con_db, "class")
	defer session.Close()
	p = Data_find_Class_Last_Cid("-cid")
	index_last := p.Cid
	if s.Cost_Each_Student_Str != "" {
		cost, err_conv := strconv.Atoi(s.Cost_Each_Student_Str)
		if err_conv != nil {
			cost = 0
			//log.Fatal("Data_find_Person_One: err_conv" + err.Error())
		}
		cost_int = cost
	}
	if s.Charge_Times != "" {
		times, err_conv := strconv.Atoi(s.Charge_Times)
		if err_conv != nil {
			times = 0
			//log.Fatal("Data_find_Person_One: err_conv" + err.Error())
		}
		charge_times = times
	}

	if s.Cost_Each_Student_Str != "" {
		cost, err_conv := strconv.Atoi(s.Cost_Each_Student_Str)
		if err_conv != nil {
			cost = 0
			//log.Fatal("Data_find_Person_One: err_conv" + err.Error())
		}
		cost_int = cost
	}
	index_last += 1
	// fmt.Println(index_last)
	err := c.Insert(
		&Class{
			Cid:          index_last,
			Class:        s.Class,
			Class_Name:   s.Class_Name,
			Day:          s.Day,
			Time_HR:      s.Time_HR,
			Duration_HR:  s.Duration_HR,
			Open_Date:    s.Open_Date,
			Teacher:      s.Teacher,
			Num_Students: len(strings.Split(s.Students, ",")),
			Course_Type:  s.Course_Type,
			//student
			Students:              s.Students,
			Charge_Times_Int:      charge_times,
			Charge_Times:          s.Charge_Times,
			Cost_Each_Student_Str: s.Cost_Each_Student_Str,
			Cost_Each_Student:     cost_int,
			UpdateTime:            time.Now().Format(con_layout),
			Delete_Record:         false,
		})
	if err != nil {
		str_error := err.Error()
		if strings.Contains(str_error, "duplicate key") {
			return ("Warning: Class name has already been taken. Please input another Name! (ex: ENG12)")
		} else {
			fmt.Println(str_error)
			//log.Fatal("Data_insert_Person: " + str_error)
			return "fail"
		}

	}
	return "success"
	//m := structs.Map(server)

}

func Data_find_Person_One(condition_identity string, sort_string string) Person {
	c, session := Data_connect_mgo(con_ip, con_db, "people")
	defer session.Close()
	//p := Person{}
	//fmt.Println(reflect.TypeOf(p))
	M := bson.M{"identity": condition_identity}
	// Query One
	result := Person{}
	err := c.Find(M).Sort(sort_string).One(&result) //Desc
	if err != nil {
		result.Uid = 0
		return result
	}
	return result
}

func Data_find_Person_Last_Uid(sort_string string) Person {
	c, session := Data_connect_mgo(con_ip, con_db, "people")
	defer session.Close()
	M := bson.M{}
	// Query One
	result := Person{}
	err := c.Find(M).Sort(sort_string).One(&result) //Desc
	if err != nil {
		result.Uid = 0
		return result
	}
	return result
}

func Data_find_Class_Last_Cid(sort_string string) Class {
	c, session := Data_connect_mgo(con_ip, con_db, "class")
	defer session.Close()
	M := bson.M{}
	// Query One
	result := Class{}
	err := c.Find(M).Sort(sort_string).One(&result) //Desc
	if err != nil {
		result.Cid = 0
		return result
	}
	return result
}

func Data_find_Person_Name_List(identity string) string {
	str_res := ""
	c, session := Data_connect_mgo(con_ip, con_db, "people")
	defer session.Close()
	M := bson.M{"identity": identity, "delete_record": false}
	person := []Person{}
	err := c.Find(M).Select(bson.M{"name": 1}).All(&person)
	if err != nil {
		log.Fatal("Data_find_Person_Name_List: " + err.Error())
	}
	for _, v := range person {
		//fmt.Println(v.Name)
		str_res += v.Name + ","
	}
	//fmt.Println(str_res[:len(str_res)-1])
	return str_res[:len(str_res)-1]
}

func Data_find_Person_list(identity string) []Person {
	c, session := Data_connect_mgo(con_ip, con_db, "people")
	defer session.Close()
	M := bson.M{"identity": identity, "delete_record": false}
	query_res := Data_find_person(c, M)
	fmt.Println(reflect.TypeOf(query_res))
	return query_res
}

func Data_find_Class_list() []Class {
	c, session := Data_connect_mgo(con_ip, con_db, "class")
	defer session.Close()
	M := bson.M{"delete_record": false}
	query_res := Data_find_class(c, M)
	fmt.Println(reflect.TypeOf(query_res))
	return query_res
}

func Data_create_Person(m bson.M) string {
	c, session := Data_connect_mgo(con_ip, con_db, "people")
	defer session.Close()
	if err := c.Insert(m); err != nil {
		fmt.Println("no")

		return "fail"
	}
	fmt.Println("ya")
	return "success"
}

// func PrintFields() int {
// 	return 0
// }

func PrintFields(b Person, query_res []Person) {
	//b := Person{}
	val := reflect.ValueOf(b)
	j := 0
	for j < len(query_res) {
		fmt.Println(query_res[j])
		for i := 0; i < val.Type().NumField(); i++ {
			fmt.Println(val.Type().Field(i).Name)

		}
		j += 1
	}
}

//str := `{"page": 1, "fruits": ["apple", "peach"]}`

// func main() {
// 	const layout = "Mon Jan _2 15:04:05 2006"

// 	// session, err := mgo.Dial("127.0.0.1")
// 	// if err != nil {
// 	//         panic(err)
// 	// }
// 	// defer session.Close()

// 	// Optional. Switch the session to a monotonic behavior.
// 	// session.SetMode(mgo.Monotonic, true)
// 	// c := session.DB("test").C("assistant_date")
// 	// // Index
// 	// index := mgo.Index{
// 	// 	Key:        []string{"teacher_class"},
// 	// 	Unique:     true,
// 	// 	DropDups:   true,
// 	// 	Background: true,
// 	// 	Sparse:     true,
// 	// }
// 	// err = c.EnsureIndex(index)
// 	// if err != nil {
// 	// 	panic(err)
// 	// }

// 	// p := new(Assistant_Date)
// 	// p.Assistant = "John"
// 	// p.Date_Attend = "20170909"
// 	// p.Punch_Start_Time = "0800"
// 	// p.Punch_End_Time = "1700"
// 	// p.UpdateTime = time.Now().Format(layout)

// 	// err = c.Insert(p)
// 	// if err != nil {
// 	// 	log.Fatal(err)
// 	// }

// 	// result := Person{}
// 	// err = c.Find(bson.M{"Name": "Eslie"}).One(&result)
// 	// if err != nil {
// 	//         log.Fatal(err)
// 	// }

// 	// fmt.Println("Phone:", result.Phone_Cell)
// }
