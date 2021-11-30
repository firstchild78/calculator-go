package selenium

import (
	"github.com/tebeka/selenium"
)

func setUp(url string) (selenium.WebDriver, selenium.Service) {
	service, err :=
		selenium.NewChromeDriverService("chromedriver.exe", 8888)
	caps := selenium.Capabilities{"browserName": "chrome"}
	wd, err := selenium.NewRemote(caps, "http://127.0.0.1:8888/wd/hub")
	wd.Refresh()
	if err != nil {
		panic(err)
	}
	wd.Get(url)
	return wd, *service
}

func calculate(wd selenium.WebDriver, leftNumber string, rightNumber string, operator string) string {

	inputLeft, err := wd.FindElement(selenium.ByID, "leftNumber")
	if err != nil {
		panic(err)
	}
	err = inputLeft.SendKeys(leftNumber)
	if err != nil {
		panic(err)
	}

	inputRight, err := wd.FindElement(selenium.ByID, "rightNumber")
	if err != nil {
		panic(err)
	}
	err = inputRight.SendKeys(rightNumber)
	selectOp, err := wd.FindElement(selenium.ByID, "operator")
	err = selectOp.SendKeys(operator)
	buttonCalc, err := wd.FindElement(selenium.ByID, "button_calculate")
	err = buttonCalc.Click()
	inputResult, err := wd.FindElement(selenium.ByID, "result")
	result, err := inputResult.GetAttribute("value")

	if err != nil {
		panic(err)
	}

	return result
}
