package main

/*
gmail
api -> web application
https://console.developers.google.com/apis/credentials?project=api-project-662368221041
https://developers.google.com/gmail/api/v1/reference/users/messages/send
calendar
https://support.google.com/calendar/answer/41207?hl=zh-Hant
https://developers.google.com/google-apps/calendar/quickstart/js
https://console.developers.google.com/apis/credentials?project=api-project-662368221041
https://developers.google.com/google-apps/calendar/v3/reference/events/insert#examples
*/
import (
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
	con_layout  = "Mon Jan _2 15:04:05 2006"
)

func index_set(c *mgo.Collection, arr_keys []string, Unique, DropDups, Background, Sparse bool) error {
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
		log.Fatal("index_set: " + err.Error())
		//panic(err)
	}
	return err
}

/*=========================================
			Connect to Mongo DB
=========================================*/
func mgo_connect(serverIP string, db string, collection string) (*mgo.Collection, *mgo.Session) {
	session, err := mgo.Dial(serverIP)
	if err != nil {
		log.Fatal("mgo_connect: " + err.Error())
		//panic(err)
	}

	session.SetMode(mgo.Monotonic, true)
	c := session.DB(db).C(collection)
	return c, session
}

/*=========================================
				Find
=========================================*/
func person_find(c *mgo.Collection, M bson.M) []Person {
	person := []Person{}
	//err := c.Find(M).One(&result)
	err := c.Find(M).All(&person)
	if err != nil {
		log.Fatal("person_find: " + err.Error())
	}
	table_count = len(person)
	return person
}

func class_find(c *mgo.Collection, M bson.M) []Class {
	class_s := []Class{}
	err := c.Find(M).All(&class_s)
	if err != nil {
		log.Fatal("class_find: " + err.Error())
	}
	table_count = len(class_s)
	return class_s
}

func account_teacher_find(c *mgo.Collection, M bson.M) []Teacher_Salary {
	s := []Teacher_Salary{}
	err := c.Find(M).All(&s)
	if err != nil {
		log.Fatal("class_find: " + err.Error())
	}
	table_count = len(s)
	return s
}

func account_student_find(c *mgo.Collection, M bson.M) []Student_Payment {
	s := []Student_Payment{}
	err := c.Find(M).All(&s)
	if err != nil {
		log.Fatal("class_find: " + err.Error())
	}
	table_count = len(s)
	return s
}

func account_assistant_find(c *mgo.Collection, M bson.M) []Assistant_Salary {
	s := []Assistant_Salary{}
	err := c.Find(M).All(&s)
	if err != nil {
		log.Fatal("class_find: " + err.Error())
	}
	table_count = len(s)
	return s
}

func notice_find(c *mgo.Collection, M bson.M) []Notice_Docs {
	s := []Notice_Docs{}
	err := c.Find(M).All(&s)
	if err != nil {
		log.Fatal("notice_find: " + err.Error())
	}
	table_count = len(s)
	fmt.Println("notice")
	fmt.Println(table_count)
	fmt.Println(s)
	return s
}

/*=========================================
				Find One
=========================================*/
func person_find_one(condition_identity string, sort_string string) Person {
	c, session := mgo_connect(con_ip, con_db, "people")
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

func person_find_Last_Uid(sort_string string) Person {
	c, session := mgo_connect(con_ip, con_db, "people")
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

func class_find_last_cid(sort_string string) Class {
	c, session := mgo_connect(con_ip, con_db, "class")
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

func notice_find_last_id(sort_string string) Notice_Docs {
	c, session := mgo_connect(con_ip, con_db, "notice_docs")
	defer session.Close()
	M := bson.M{}
	// Query One
	result := Notice_Docs{}
	err := c.Find(M).Sort(sort_string).One(&result) //Desc
	if err != nil {
		result.NDid = 0
		return result
	}
	return result
}

func person_find_Name_List(identity string) string {
	str_res := ""
	c, session := mgo_connect(con_ip, con_db, "people")
	defer session.Close()
	M := bson.M{"identity": identity, "delete_record": false}
	person := []Person{}
	err := c.Find(M).Select(bson.M{"name": 1}).All(&person)
	if err != nil {
		log.Fatal("person_find_Name_List: " + err.Error())
	}
	for _, v := range person {
		//fmt.Println(v.Name)
		str_res += v.Name + ","
	}
	//fmt.Println(str_res[:len(str_res)-1])
	return str_res[:len(str_res)-1]
}

/*=========================================
				Find List
=========================================*/
func person_find_list(identity string) []Person {
	c, session := mgo_connect(con_ip, con_db, "people")
	defer session.Close()
	M := bson.M{"identity": identity, "delete_record": false}
	query_res := person_find(c, M)
	fmt.Println(reflect.TypeOf(query_res))
	return query_res
}

func class_find_list() []Class {
	c, session := mgo_connect(con_ip, con_db, "class")
	defer session.Close()
	M := bson.M{"delete_record": false}
	query_res := class_find(c, M)
	fmt.Println(reflect.TypeOf(query_res))
	return query_res
}

func account_find_student_list() []Student_Payment {
	c, session := mgo_connect(con_ip, con_db, "student_payment")
	defer session.Close()
	M := bson.M{"delete_record": false}
	query_res := account_student_find(c, M)
	fmt.Println(reflect.TypeOf(query_res))
	return query_res
}

func account_find_teacher_list() []Teacher_Salary {
	c, session := mgo_connect(con_ip, con_db, "teacher_salary")
	defer session.Close()
	M := bson.M{"delete_record": false}
	query_res := account_teacher_find(c, M)
	fmt.Println(reflect.TypeOf(query_res))
	return query_res
}

func account_find_assistant_list() []Assistant_Salary {
	c, session := mgo_connect(con_ip, con_db, "assistant_salary")
	defer session.Close()
	M := bson.M{"delete_record": false}
	query_res := account_assistant_find(c, M)
	fmt.Println(reflect.TypeOf(query_res))
	return query_res
}

func notice_find_list() []Notice_Docs {
	c, session := mgo_connect(con_ip, con_db, "notice_docs")
	defer session.Close()
	M := bson.M{"delete_record": false}
	query_res := notice_find(c, M)
	fmt.Println(reflect.TypeOf(query_res))
	return query_res
}

/*=========================================
				Remove
=========================================*/
func person_remove(s []Person, identity string) string {
	//p := Person{}
	fmt.Println("===dmp")
	fmt.Println(s)

	c, session := mgo_connect(con_ip, con_db, "people")
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
			log.Fatal("person_remove: " + err.Error())
		}

	}

	return "success"
}

func class_remove(s []Class) string {
	//p := Person{}
	fmt.Println("===dmc")
	fmt.Println(s)

	c, session := mgo_connect(con_ip, con_db, "class")
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
			log.Fatal("class_remove: " + err.Error())
		}

	}

	return "success"
}

func account_teacher_remove(s []Teacher_Salary) string {
	// //p := Person{}
	// fmt.Println("===dmc")
	// fmt.Println(s)

	// c, session := mgo_connect(con_ip, con_db, "class")
	// defer session.Close()
	// for _, ss := range s {
	// 	//suport to remove multiple records
	// 	colQuerier := bson.M{
	// 		strings.ToLower("Class"): ss.Class,
	// 	}
	// 	update_field := bson.M{
	// 		strings.ToLower("UpdateTime"):    time.Now().Format(con_layout),
	// 		strings.ToLower("Delete_Record"): true,
	// 	}
	// 	change := bson.M{"$set": update_field}
	// 	fmt.Println(colQuerier)
	// 	fmt.Println(update_field)
	// 	fmt.Println(change)
	// 	err := c.Update(colQuerier, change)
	// 	if err != nil {
	// 		fmt.Println(err.Error())
	// 		log.Fatal("class_remove: " + err.Error())
	// 	}

	// }

	return "success"
}
func account_student_remove(s []Student_Payment) string {
	// //p := Person{}
	// fmt.Println("===dmc")
	// fmt.Println(s)

	// c, session := mgo_connect(con_ip, con_db, "class")
	// defer session.Close()
	// for _, ss := range s {
	// 	//suport to remove multiple records
	// 	colQuerier := bson.M{
	// 		strings.ToLower("Class"): ss.Class,
	// 	}
	// 	update_field := bson.M{
	// 		strings.ToLower("UpdateTime"):    time.Now().Format(con_layout),
	// 		strings.ToLower("Delete_Record"): true,
	// 	}
	// 	change := bson.M{"$set": update_field}
	// 	fmt.Println(colQuerier)
	// 	fmt.Println(update_field)
	// 	fmt.Println(change)
	// 	err := c.Update(colQuerier, change)
	// 	if err != nil {
	// 		fmt.Println(err.Error())
	// 		log.Fatal("class_remove: " + err.Error())
	// 	}

	// }

	return "success"
}
func account_assistant_remove(s []Assistant_Salary) string {
	// //p := Person{}
	// fmt.Println("===dmc")
	// fmt.Println(s)

	// c, session := mgo_connect(con_ip, con_db, "class")
	// defer session.Close()
	// for _, ss := range s {
	// 	//suport to remove multiple records
	// 	colQuerier := bson.M{
	// 		strings.ToLower("Class"): ss.Class,
	// 	}
	// 	update_field := bson.M{
	// 		strings.ToLower("UpdateTime"):    time.Now().Format(con_layout),
	// 		strings.ToLower("Delete_Record"): true,
	// 	}
	// 	change := bson.M{"$set": update_field}
	// 	fmt.Println(colQuerier)
	// 	fmt.Println(update_field)
	// 	fmt.Println(change)
	// 	err := c.Update(colQuerier, change)
	// 	if err != nil {
	// 		fmt.Println(err.Error())
	// 		log.Fatal("class_remove: " + err.Error())
	// 	}

	// }

	return "success"
}

func notice_remove(s []Notice_Docs) string {
	//p := Person{}
	fmt.Println("===dmc")
	fmt.Println(s)

	c, session := mgo_connect(con_ip, con_db, "notice_docs")
	defer session.Close()
	for _, ss := range s {
		//suport to remove multiple records
		colQuerier := bson.M{
			strings.ToLower("NDid"): ss.NDid,
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
			log.Fatal("notice_remove: " + err.Error())
		}

	}

	return "success"
}

/*=========================================
				Update
=========================================*/
func class_update(s Class) string {
	cost_int := 0
	//charge_times := 0
	c, session := mgo_connect(con_ip, con_db, "class")
	defer session.Close()
	if s.Cost_Each_Student_Str != "" {
		cost, err_conv := strconv.Atoi(s.Cost_Each_Student_Str)
		if err_conv != nil {
			cost = 0
			//log.Fatal("person_find_one: err_conv" + err.Error())
		}
		cost_int = cost
	}

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
		//log.Fatal("person_update: " + err.Error())
	}
	return "success"
}

func person_update(s Person, identity string) string {
	//p := Person{}
	salary_int := 0
	c, session := mgo_connect(con_ip, con_db, "people")
	defer session.Close()
	if s.Salary_HR != "" {
		sa, err_conv := strconv.Atoi(s.Salary_HR)
		if err_conv != nil {
			sa = 0
			//log.Fatal("person_find_one: err_conv" + err.Error())
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
		//log.Fatal("person_update: " + err.Error())
	}
	return "success"
}

func account_student_update(s Student_Payment) string {

	return "success"
}
func account_assistant_update(s Assistant_Salary) string {

	return "success"
}
func account_teacher_update(s Teacher_Salary) string {

	return "success"
}

func notice_update(s Notice_Docs) string {
	c, session := mgo_connect(con_ip, con_db, "notice_docs")
	defer session.Close()

	// Update
	colQuerier := bson.M{
		strings.ToLower("NDid"): s.NDid,
	}
	update_field := bson.M{
		strings.ToLower("Target"):        s.Target,
		strings.ToLower("Title"):         s.Title,
		strings.ToLower("Content"):       s.Content,
		strings.ToLower("UpdateTime"):    time.Now().Format(con_layout),
		strings.ToLower("Delete_Record"): false,
	}
	change := bson.M{"$set": update_field}
	fmt.Println(colQuerier)
	fmt.Println(update_field)
	fmt.Println(change)
	err := c.Update(colQuerier, change)
	if err != nil {
		fmt.Println(err.Error())
		return ("Warning: System Error. DUP")
		//log.Fatal("person_update: " + err.Error())
	}
	return "success"
}

/*=========================================
				Insert
=========================================*/
func person_insert(s Person, identity string) string {
	salary_int := 0
	fmt.Println(s.Name)
	p := Person{}
	c, session := mgo_connect(con_ip, con_db, "people")
	defer session.Close()
	p = person_find_Last_Uid("-uid")
	index_last := p.Uid
	// fmt.Println("index_last")
	// fmt.Println(index_last)
	//index, err := strconv.Atoi(last_index)
	//https://github.com/fatih/structs
	if s.Salary_HR != "" {
		sa, err_conv := strconv.Atoi(s.Salary_HR)
		if err_conv != nil {
			sa = 0
			//log.Fatal("person_find_one: err_conv" + err.Error())
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
			//log.Fatal("person_insert: " + str_error)
			return "fail"
		}

	}
	return "success"
	//m := structs.Map(server)

}

func class_insert(s Class) string {

	cost_int := 0
	charge_times := 0
	//fmt.Println(s.Name)
	p := Class{}
	c, session := mgo_connect(con_ip, con_db, "class")
	defer session.Close()
	p = class_find_last_cid("-cid")
	index_last := p.Cid
	if s.Cost_Each_Student_Str != "" {
		cost, err_conv := strconv.Atoi(s.Cost_Each_Student_Str)
		if err_conv != nil {
			cost = 0
			//log.Fatal("person_find_one: err_conv" + err.Error())
		}
		cost_int = cost
	}
	if s.Charge_Times != "" {
		times, err_conv := strconv.Atoi(s.Charge_Times)
		if err_conv != nil {
			times = 0
			//log.Fatal("person_find_one: err_conv" + err.Error())
		}
		charge_times = times
	}

	if s.Cost_Each_Student_Str != "" {
		cost, err_conv := strconv.Atoi(s.Cost_Each_Student_Str)
		if err_conv != nil {
			cost = 0
			//log.Fatal("person_find_one: err_conv" + err.Error())
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
			//log.Fatal("person_insert: " + str_error)
			return "fail"
		}

	}
	return "success"
	//m := structs.Map(server)

}

func account_student_insert(s Student_Payment) string {
	return "success"

}

func account_teacher_insert(s Teacher_Salary) string {
	return "success"

}
func account_assistant_insert(s Assistant_Salary) string {
	return "success"

}

func notice_insert(s Notice_Docs) string {
	// cost_int := 0
	// charge_times := 0
	//fmt.Println(s.Name)
	p := Notice_Docs{}
	c, session := mgo_connect(con_ip, con_db, "notice_docs")
	defer session.Close()
	p = notice_find_last_id("-ndid") //<<<<<<<<===
	index_last := p.NDid
	// if s.Cost_Each_Student_Str != "" {
	// 	cost, err_conv := strconv.Atoi(s.Cost_Each_Student_Str)
	// 	if err_conv != nil {
	// 		cost = 0
	// 		//log.Fatal("person_find_one: err_conv" + err.Error())
	// 	}
	// 	cost_int = cost
	// }
	// if s.Charge_Times != "" {
	// 	times, err_conv := strconv.Atoi(s.Charge_Times)
	// 	if err_conv != nil {
	// 		times = 0
	// 		//log.Fatal("person_find_one: err_conv" + err.Error())
	// 	}
	// 	charge_times = times
	// }

	// if s.Cost_Each_Student_Str != "" {
	// 	cost, err_conv := strconv.Atoi(s.Cost_Each_Student_Str)
	// 	if err_conv != nil {
	// 		cost = 0
	// 		//log.Fatal("person_find_one: err_conv" + err.Error())
	// 	}
	// 	cost_int = cost
	// }
	index_last += 1
	// fmt.Println(index_last)
	err := c.Insert(
		&Notice_Docs{
			NDid:          index_last,
			Target:        s.Target,
			Title:         s.Title,
			Content:       s.Content,
			UpdateTime:    time.Now().Format(con_layout),
			Delete_Record: false,
		})
	if err != nil {
		str_error := err.Error()
		if strings.Contains(str_error, "duplicate key") {
			return ("Warning: Class name has already been taken. Please input another Name! (ex: ENG12)")
		} else {
			fmt.Println(str_error)
			//log.Fatal("person_insert: " + str_error)
			return "fail"
		}

	}
	return "success"
}

func person_create(m bson.M) string {
	c, session := mgo_connect(con_ip, con_db, "people")
	defer session.Close()
	if err := c.Insert(m); err != nil {
		fmt.Println("no")

		return "fail"
	}
	fmt.Println("ya")
	return "success"
}

/*=========================================
				Other Functions
=========================================*/
func fields_print(b Person, query_res []Person) {
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
