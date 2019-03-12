package DocParser

import (
	Model "../models"
	"log"
	"os/exec"
	"regexp"
	"strconv"
	"strings"
)

func ParseQuestionsText(str string) []Model.Question {
	var re = regexp.MustCompile(`(?m)(?P<question>^\s*\((?P<id>\d{1,2})\).*)[\n]*\s*(?P<optionA>a\).*)
[\n]*\s*(?P<optionB>b\).*)[\n]*\s*(?P<optionC>c\).*)[\n]*\s*(?P<optionD>d\).*)`)

	res := re.FindAllStringSubmatch(str, -1)
	log.Print("матчей:", len(res))
	questions := make([]Model.Question, len(res))
	for i := range res {
		log.Print("субматчей", len(res[i]))
		id, _ := strconv.Atoi(res[i][2])
		questions[i] = Model.Question{Question: strings.Replace(strings.TrimSpace(res[i][1]), "\t", "", -1), Id: id, OptionA: strings.TrimSpace(res[i][3]),
			OptionB: strings.TrimSpace(res[i][4]), OptionC: strings.TrimSpace(res[i][5]), OptionD: strings.TrimSpace(res[i][6])}
	}
	return questions
}

func ParseAnswers(str string) []Model.Answer {
	var re = regexp.MustCompile(`\((?P<id>\d{1,2})\)\s*(?P<answer>[abcd])`)

	res := re.FindAllStringSubmatch(str, -1)
	answers := make([]Model.Answer, len(res))
	for i := range res {
		id, _ := strconv.Atoi(res[i][1])
		answers[i] = Model.Answer{Id: id, Answer: res[i][2]}
	}
	return answers
}

func GetReadingText(str string) string {
	idx1 := strings.Index(str, "Read the text below. For questions 21 to 25, choose the best answer (a, b, c or d).")
	idx2 := strings.Index(str, "(21)")

	res := str[idx1:idx2]

	return res
}

func GetTestFromDocx(questionsPath , answersPath string) *Model.Test {
	//TODO: DELETE CREATE ISSUES
	cmd := exec.Command("libreoffice", "--headless", "--cat", "txt:Text (encoded):UTF8", questionsPath)
	str, err := cmd.Output()
	if err != nil {
		log.Print(err)
	}
	questions := ParseQuestionsText(string(str))
	readingText := GetReadingText(string(str))
	cmd = exec.Command("libreoffice", "--headless", "--cat", "txt:Text (encoded):UTF8", answersPath)
	str, err = cmd.Output()
	if err != nil {
		log.Print(err)
	}
	answers := ParseAnswers(string(str))

	test:= CreateTest(questions,answers,readingText)

	return test
}
func CreateTest(questions []Model.Question,answers []Model.Answer, readingText string) *Model.Test{
	var test Model.Test
	test.BaseQuestions = make([]Model.Question,60)
	k:=0
	for i:=0;i<len(questions);i++ {
		if(i>=20 && i<=24){
			continue
		}
		test.BaseQuestions[k]=questions[i]
		test.BaseQuestions[k].Id=k+1
		k++
	}
	test.ReadingQuestions = new(Model.Reading)
	test.ReadingQuestions.Question =  readingText

	test.ReadingQuestions.BaseQuestion = make([]Model.Question,5)
	k=0
	for i:=20;i<25;i++ {
		test.ReadingQuestions.BaseQuestion[k]=questions[i]
		test.ReadingQuestions.BaseQuestion[k].Id=k+1
		k++
	}

	test.Answers = new(Model.AnswerContainer)
	test.Answers.Reading = make([]Model.Answer,5)
	test.Answers.Base = make([]Model.Answer,60)
	test.Answers.Writing = "";

	k=0
	for i:=0;i<len(answers);i++ {
		if(i>=20 && i<=24){
			continue
		}
		test.Answers.Base[k]=answers[i]
		test.Answers.Base[k].Id=k+1
		k++
	}

	k=0
	for i:=20;i<25;i++ {
		test.Answers.Reading[k]=answers[i]
		test.Answers.Reading[k].Id=k+1
		k++
	}

	return &test

}


