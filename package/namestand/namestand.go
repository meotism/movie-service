package namestand

import (
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
	"strings"
)

func Check(str string) string {
	re1, err := regexp.Compile(`[^a-zA-Z\s]+`)

	if err != nil {
		log.Fatal(err)
	}

	re2, err := regexp.Compile(`\s\s+`)

	if err != nil {
		log.Fatal(err)
	}

	re3, err := regexp.Compile(`^ | $`)

	if err != nil {
		log.Fatal(err)
	}

	match := re1.ReplaceAllString(str, "")     // remove all digit and non-word
	match2 := re2.ReplaceAllString(match, " ") // remove double space
	match3 := re3.ReplaceAllString(match2, "") // remove space in first and end
	match3 = strings.ToLower(match3)
	strArr := strings.Split(match3, "")
	strArr[0] = strings.ToUpper(strArr[0])

	for i := 0; i < len(strArr); i++ {
		if strArr[i] == " " {
			strArr[i+1] = strings.ToUpper(strArr[i+1])
		}
	}

	return strings.Replace(strings.Join(strArr[:], ","), ",", "", -1)
}

func RemoveDoubleSpace(str string) string {
	re2, err := regexp.Compile(`\s\s+`)

	if err != nil {
		log.Fatal(err)
	}

	str = re2.ReplaceAllString(str, " ") // remove double space

	if str == " " {
		str = ""
	}

	return str
}

func FormatText(str string, isHaveNumber bool, isLowerCase bool) string {

	var regexStr = `[^a-zA-Z\s]+`

	if isHaveNumber {
		regexStr = `[^a-zA-Z0-9\s]+`
	}

	re1, err := regexp.Compile(regexStr)

	if err != nil {
		log.Fatal(err)
	}

	re3, err := regexp.Compile(`^ | $`)

	if err != nil {
		log.Fatal(err)
	}

	match := re1.ReplaceAllString(str, "")    // remove non-word
	match2 := re3.ReplaceAllString(match, "") // remove space in first and end

	if isLowerCase {
		match2 = strings.ToLower(match2) //Lowercase all
	}

	return match2
}

func IsDate(str string) string {

	var regexStr = `/^(\d{4})-(\d{2})-(\d{2}) (\d{2}):(\d{2}):(\d{2})$/`
	re1, err := regexp.Compile(regexStr)

	if err != nil {
		log.Fatal(err)
	}

	match := re1.FindString(str) // remove non-word

	return match
}

func IsContainBadWord(str string) bool {
	client := &http.Client{}
	url := "https://checkbadwordapi.herokuapp.com/check/" + str
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("accept", "application/json")
	req.Header.Set("api-key", "Y2hlY2tiYWR3b3JkYXBpa2V5")
	res, _ := client.Do(req)

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatalln(err)
	}
	bodyStr := string(body)

	isBad := strings.Contains(bodyStr, `"is_bad":true`)

	return isBad
}
