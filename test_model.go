package main

import (
	"fmt"
	"log"
	"reflect"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Person struct {
	Name                    string   //[uniqle]
	Phone_Home              string   // Nil
	Phone_Cell              string   //Nil
	EMAIL                   string   // Nil
	Identity                string   // student/ teacher/ assistant
	Classes                 []string //[A,B..] or nil
	Salary_HR               int      //[same as Salary table] for Teacher and Assistant only
	Student_Parents_Contect string
	Note                    string
	UpdateTime              string // Array timestamp
	Delete_Record           bool
}

func Data_connect_mgo(serverIP string, db string, collection string) (*mgo.Collection, *mgo.Session) {
	session, err := mgo.Dial(serverIP)
	if err != nil {
		panic(err)
	}

	session.SetMode(mgo.Monotonic, true)
	c := session.DB(db).C(collection)
	return c, session
}

func Data_person_find(c *mgo.Collection, M bson.M) Person {
	result := Person{}
	err := c.Find(M).One(&result)
	if err != nil {
		log.Fatal(err)
	}
	return result
}

func Data_find_collection() {
	c, session := Data_connect_mgo("127.0.0.1", "test", "people")
	defer session.Close()
	result := Person{}
	fmt.Println(reflect.TypeOf(result))
	M := bson.M{"name": "Eslie"}
	// fmt.Println(reflect.TypeOf(c))
	// fmt.Println(reflect.TypeOf(M))
	res := Data_person_find(c, M)
	fmt.Println("Phone:", res.Phone_Cell)

}

func main() {
	Data_find_collection()
}
