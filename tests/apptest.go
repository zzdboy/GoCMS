package tests

import "time"
import "github.com/revel/revel"

type AppTest struct {
	revel.TestSuite
}

func (t *AppTest) Before() {
	println("Start time:" + time.Now().Format("2006-01-02 15:04:04"))
}

func (t AppTest) TestThatIndexPageWorks() {
	t.Get("/")
	t.AssertOk()
	t.AssertContentType("text/html")
}

func (t *AppTest) After() {
	println("End time:" + time.Now().Format("2006-01-02 15:04:04"))
}
