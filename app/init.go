package app

import (

	"github.com/revel/revel"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	"fmt"

	"os"
	"strings"
	"reflect"
	"time"
)

var DB *gorm.DB

func InitDB() {
	cnnString := fmt.Sprintf("host=%s user=%s password='%s' dbname=%s sslmode=disable","localhost", "postgres", "azura777", "aexpense")

	var err error
	DB, err = gorm.Open("postgres", cnnString)
	if err != nil {
		revel.INFO.Println("DB Error", err)
	}
	revel.INFO.Println("DB Connected")
	//AutoMigrate()
}

func init() {
	// Filters is the default set of global filters.
	revel.Filters = []revel.Filter{
		revel.PanicFilter,             // Recover from panics and display an error page instead.
		revel.RouterFilter,            // Use the routing table to select the right Action
		revel.FilterConfiguringFilter, // A hook for adding or removing per-Action filters.
		revel.ParamsFilter,            // Parse parameters into Controller.Params.
		revel.SessionFilter,           // Restore and write the session cookie.
		revel.FlashFilter,             // Restore and write the flash cookie.
		revel.ValidationFilter,        // Restore kept validation errors and save new ones from cookie.
		revel.I18nFilter,              // Resolve the requested language
		HeaderFilter,                  // Add some security based headers
		revel.InterceptorFilter,       // Run interceptors around the action.
		revel.CompressFilter,          // Compress the result.
		revel.ActionInvoker,           // Invoke the action.
	}

	// register startup functions with OnAppStart
	// ( order dependent )
	// revel.OnAppStart(InitDB)
	// revel.OnAppStart(FillCache)
	//revel.OnAppStart(func() {
	//	jobs.Schedule("@midnight",MyJob())
	//	jobs.Every(8 * time.Hour,MyJob())
	//	Asynch olarak bir job çalıştır.
	//	jobs.Now(MyJob{})
	//  5 dakika sonra çalıştır.
	//  jobs.In(5 * time.minute,MyJob{})
	//})

	revel.OnAppStart(func() {
		InitDB()
	})

}

// TODO turn this into revel.HeaderFilter
// should probably also have a filter for CSRF
// not sure if it can go in the same filter or not
var HeaderFilter = func(c *revel.Controller, fc []revel.Filter) {
	// Add some common security headers
	c.Response.Out.Header().Add("X-Frame-Options", "SAMEORIGIN")
	c.Response.Out.Header().Add("X-XSS-Protection", "1; mode=block")
	c.Response.Out.Header().Add("X-Content-Type-Options", "nosniff")

	fc[0](c, fc[1:]) // Execute the next filter stage.
}


func WriteLines(lines []string, path string) (err error) {
	var (
		file *os.File
	)

	if file, err = os.Create(path); err != nil {
		return
	}
	defer file.Close()

	//writer := bufio.NewWriter(file)
	for _,item := range lines {
		//fmt.Println(item)
		_, err := file.WriteString(strings.TrimSpace(item) + "\n");
		//file.Write([]byte(item));
		if err != nil {
			//fmt.Println("debug")
			fmt.Println(err)
			break
		}
	}
	/*content := strings.Join(lines, "\n")
	_, err = writer.WriteString(content)*/
	return
}


func MakeCaptionML(this interface{}) {
	var ModelLangs_TRK []string
	var ModelLangs_ENU []string

	t := reflect.TypeOf(this)

	ModelLangs_TRK = append(ModelLangs_TRK, "")
	ModelLangs_TRK = append(ModelLangs_TRK, "# --------------------------------------------")
	ModelLangs_TRK = append(ModelLangs_TRK, fmt.Sprintf("# [%s] table defination for Turkish", t.Name()))
	ModelLangs_TRK = append(ModelLangs_TRK, fmt.Sprintf("# Dev : Ibrahim COBANI %s", time.Now().Local().Format("2006-01-02 15:04:05")))
	ModelLangs_TRK = append(ModelLangs_TRK, "# --------------------------------------------")
	ModelLangs_TRK = append(ModelLangs_TRK, "")
	ModelLangs_TRK = append(ModelLangs_TRK, "")

	ModelLangs_ENU = append(ModelLangs_ENU, "")
	ModelLangs_ENU = append(ModelLangs_ENU, "# --------------------------------------------")
	ModelLangs_ENU = append(ModelLangs_ENU, fmt.Sprintf("# [%s] table defination for English", t.Name()))
	ModelLangs_ENU = append(ModelLangs_ENU, fmt.Sprintf("# Dev : Ibrahim COBANI %s", time.Now().Local().Format("2006-01-02 15:04:05")))
	ModelLangs_ENU = append(ModelLangs_ENU, "# --------------------------------------------")
	ModelLangs_ENU = append(ModelLangs_ENU, "")
	ModelLangs_ENU = append(ModelLangs_ENU, "")


	for i := 0; i < t.NumField(); i++ {
		tag := t.Field(i).Tag.Get("CaptionML")
		if tag != "" {
			var CaptionMLs []string
			CaptionMLs = strings.Split(tag, ";")
			for _, CaptionML := range CaptionMLs {
				var MLTags []string
				MLTags = strings.Split(CaptionML, "=")
				switch strings.TrimSpace(MLTags[0]) {
				case "trk", "TRK":
					ModelLangs_TRK = append(ModelLangs_TRK, fmt.Sprintf("%s.%s = %s\n", t.Name(), t.Field(i).Name, MLTags[1]))
					break
				case "enu", "ENU":
					ModelLangs_ENU = append(ModelLangs_ENU, fmt.Sprintf("%s.%s = %s\n", t.Name(), t.Field(i).Name, MLTags[1]))
					break
				}
			}
		}
	}

	WriteLines(ModelLangs_TRK, "messages/m_" + t.Name() + ".tr")
	WriteLines(ModelLangs_ENU, "messages/m_" + t.Name() + ".en")
}