package swagger

import (
	"bufio"
	"io/ioutil"
	"log"
	"net/http"
)

func TestPost(w http.ResponseWriter, r *http.Request) {
	//w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	//r.ParseForm()
	r.ParseForm()
	stbyte,_:=ioutil.ReadAll(r.Body)
	log.Print(string(stbyte))
	text:=r.FormValue("testText")
	log.Print(text)
	//email := r.FormValue("email")
	//pass := r.FormValue("password")
//	str := r.FormValue("testText")
//	stbyte,_:=ioutil.ReadAll(r.Body)
//	parsedValue, err := url.QueryUnescape(string(stbyte))
//	log.Print(string(stbyte))
	//dec := json.NewDecoder(r.Body)
	//var test Test
	//json.Unmarshal([]byte(parsedValue), &test)
	//if err := dec.Decode(&test); err == io.EOF {
	//	//OK
	//} else if err != nil {
	//	log.Fatal(err)
	//}
	if err := dec.Decode(&test); err == io.EOF {
		//OK
	} else if err != nil {
		log.Fatal(err)
	}

	//_, err = db.Model(&test).Insert()
	//if err != nil {
	//	log.Print(err)
	//	w.WriteHeader(http.StatusInternalServerError)
	//}
	w.WriteHeader(http.StatusOK)
}

func TestPut(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}

func CheckCredentialsTeacherPost(w http.ResponseWriter, r *http.Request) {
	var teachers []Teacher
	scanner := bufio.NewReader(r.Body)

	res, _, _ := scanner.ReadLine()
	login := string(res)

	res, _, _ = scanner.ReadLine()
	pass := string(res)

	err := db.Model(&teachers).Where("login = ? and password = ?", login, pass).Select()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}
	if len(teachers) == 0 {
		w.Write([]byte("no"))
	} else {
		w.Write([]byte("yes"))
	}
	w.WriteHeader(http.StatusOK)

}

func CheckCredentialsPost(w http.ResponseWriter, r *http.Request) {
	var student []Student

	scanner := bufio.NewReader(r.Body)

	res, _, _ := scanner.ReadLine()
	login := string(res)

	res, _, _ = scanner.ReadLine()
	pass := string(res)

	err := db.Model(&student).Where("email = ? and password = ?", login, pass).Select()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}
	if len(student) == 0 {
		w.Write([]byte("no"))
	} else {
		w.Write([]byte("yes"))
	}
	w.WriteHeader(http.StatusOK)

}
