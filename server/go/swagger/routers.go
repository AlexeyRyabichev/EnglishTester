package swagger

import (
	api "../apis"
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
		handler = api.AuthMiddleware(handler, route.Name)
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
		api.AudioStudentIdGet,
	},

	Route{
		"AudioStudentIdPost",
		strings.ToUpper("Post"),
		"/api/audio/{studentId}",
		api.AudioStudentIdPost,
	},

	Route{
		"StudentCreateWithArrayPost",
		strings.ToUpper("Post"),
		"/api/student/createWithArray",
		api.StudentCreateWithArrayPost,
	},

	Route{
		"StudentPost",
		strings.ToUpper("Post"),
		"/api/student",
		api.StudentPost,
	},

	Route{
		"StudentPut",
		strings.ToUpper("Put"),
		"/api/student",
		api.StudentPut,
	},
	Route{
		"StudentDelete",
		strings.ToUpper("Delete"),
		"/api/student",
		api.StudentDelete,
	},

	Route{
		"StudentsDelete",
		strings.ToUpper("Delete"),
		"/api/students",
		api.StudentsDelete,
	},

	Route{
		"StudentsGet",
		strings.ToUpper("Get"),
		"/api/students",
		api.StudentsGet,
	},

	Route{
		"TestPost",
		strings.ToUpper("Post"),
		"/api/test",
		api.TestPost,
	},

	Route{
		"TestPut",
		strings.ToUpper("Put"),
		"/api/test",
		api.TestPut,
	},

	Route{
		"TestsGet",
		strings.ToUpper("Get"),
		"/api/tests",
		api.TestsGet,
	},

	Route{
		"TestGet",
		strings.ToUpper("Get"),
		"/api/test/{testId}",
		api.TestGet,
	},

	Route{
		"QuestionsGet",
		strings.ToUpper("Get"),
		"/api/questions",
		api.QuestionsGet,
	},

	Route{
		"TeachersGet",
		strings.ToUpper("Get"),
		"/api/teachers",
		api.TeachersGet,
	},

	Route{
		"TeacherPost",
		strings.ToUpper("Post"),
		"/api/teacher",
		api.TeacherPost,
	},

	Route{
		"TeacherDelete",
		strings.ToUpper("Delete"),
		"/api/teacher",
		api.TeacherDelete,
	},

	Route{
		"CheckCredentialsTeacherPost",
		strings.ToUpper("Post"),
		"/api/checkCredentialsTeacher",
		api.CheckCredentialsTeacherPost,
	},
	Route{
		"CheckCredentialsPost",
		strings.ToUpper("Post"),
		"/api/checkCredentials",
		api.CheckCredentialsPost,
	},
	Route{
		"LoginPost",
		strings.ToUpper("Post"),
		"/api/login",
		api.LoginPost,
	},

	Route{
		"AnswersPost",
		strings.ToUpper("Post"),
		"/api/answer",
		api.AnswersPost,
	},
}
