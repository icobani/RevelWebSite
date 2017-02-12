/******
* B1 Yönetim Sistemleri Yazılım ve Danışmanlık Limited Şirketi
* B1 Digitial
* http://www.b1.com.tr
*
*
*
* Date      : 11/02/2017    
* Time      : 21:48
* Developer : ibrahimcobani
*
*******/
package BusinessLibs

import (
	"strings"
	"regexp"
	"fmt"
)

type HumanParser struct {
	Name     string
	LastName string
	FullName string
}

func (this *HumanParser) Set(val string) {

	fmt.Println(val)

	replace := func(word string) string {
		return strings.Title(word)
	}
	r := regexp.MustCompile(`\w+`)
	val = r.ReplaceAllStringFunc(strings.ToLower(val), replace)
	fmt.Println("...")
	fmt.Println(val)
	this.FullName = val

	names := strings.Fields(val)
	namescount := len(names)

	fmt.Println(namescount, names)

	if namescount == 1 {
		this.Name = val
	}

	if namescount == 2 {
		this.Name = names[0]
		this.LastName = strings.ToUpper(names[1])
	}

	if namescount > 2 {
		for i := 0; i < namescount - 1; i++ {
			if this.Name != ""{
				this.Name += " "
			}
			this.Name += names[i]
		}
		this.LastName = strings.ToUpper(names[namescount])
	}

}
