package swagger

import (
	"github.com/gorilla/mux"
	"net/http"
	"strings"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		var handler http.Handler
		handler = route.HandlerFunc
		handler = authMiddleware(handler, route.Name)
		handler = Logger(handler, route.Name)

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}

	return router

}

var routes = Routes{
	Route{
		"AudioStudentIdGet",
		strings.ToUpper("Get"),
		"/api/audio/{studentId}",
		AudioStudentIdGet,
	},

	Route{
		"AudioStudentIdPost",
		strings.ToUpper("Post"),
		"/api/audio/{studentId}",
		AudioStudentIdPost,
	},

	Route{
		"StudentCreateWithArrayPost",
		strings.ToUpper("Post"),
		"/api/student/createWithArray",
		StudentCreateWithArrayPost,
	},

	Route{
		"StudentPost",
		strings.ToUpper("Post"),
		"/api/student",
		StudentPost,
	},

	Route{
		"StudentPut",
		strings.ToUpper("Put"),
		"/api/student",
		StudentPut,
	},
	Route{
		"StudentDelete",
		strings.ToUpper("Delete"),
		"/api/student",
		StudentDelete,
	},

	Route{
		"StudentsDelete",
		strings.ToUpper("Delete"),
		"/api/students",
		StudentsDelete,
	},

	Route{
		"StudentsGet",
		strings.ToUpper("Get"),
		"/api/students",
		StudentsGet,
	},

	Route{
		"TestPost",
		strings.ToUpper("Post"),
		"/api/test",
		TestPost,
	},

	Route{
		"TestPut",
		strings.ToUpper("Put"),
		"/api/test",
		TestPut,
	},

	Route{
		"QuestionsGet",
		strings.ToUpper("Get"),
		"/api/questions",
		QuestionsGet,
	},

	Route{
		"TeachersGet",
		strings.ToUpper("Get"),
		"/api/teachers",
		TeachersGet,
	},

	Route{
		"TeacherPost",
		strings.ToUpper("Post"),
		"/api/teacher",
		TeacherPost,
	},

	Route{
		"TeacherDelete",
		strings.ToUpper("Delete"),
		"/api/teacher",
		TeacherDelete,
	},

	Route{
		"CheckCredentialsTeacherPost",
		strings.ToUpper("Post"),
		"/api/checkCredentialsTeacher",
		CheckCredentialsTeacherPost,
	},
	Route{
		"CheckCredentialsPost",
		strings.ToUpper("Post"),
		"/api/checkCredentials",
		CheckCredentialsPost,
	},
	Route{
		"LoginPost",
		strings.ToUpper("Post"),
		"/api/login",
		LoginPost,
	},
}
