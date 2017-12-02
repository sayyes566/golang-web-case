package main

type Power struct {
	Power string `json:"power"`
	//   bool   `json:"power"`
}

/*==========Person Management=================*/
type Person struct {
	Power string `json:"power"`
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

type Histry_Notice struct {
	HNid          int    `json: "hnid" bson:"hnid"`
	Identity      string //student or teacher
	Person        string //name
	Mail          string //name
	DocName       string
	UpdateTime    string
	Delete_Record bool
}

type Notice_Docs struct {
	NDid          int    `json: "hnid" bson:"hnid"`
	Target        string //student or teacher
	Title         string
	Content       string
	UpdateTime    string
	Delete_Record bool
}
