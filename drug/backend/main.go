package main
import (
	"context"
	"log"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/mmildd_s/app/controllers"
	"github.com/mmildd_s/app/ent"
)
type Doctors struct {
	Doctor []Doctor
}
type Doctor struct {
	DoctorEmail  	string
	DoctorPassword 	string
	DoctorName		string
	DoctorTel		string
}
type Patients struct {
	Patient []Patient
}

type Patient struct {
	PatientName  string
}

type Medicines struct {
	Medicine []Medicine
}

type Medicine struct {
	MedicineName string
}

type Manners struct {
	Manner []Manner
}

type Manner struct {
	MannerName string
}

// @title SUT SA Example API Playlist Vidoe
// @version 1.0
// @description This is a sample server for SUT SE 2563
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:8080
// @BasePath /api/v1
// @securityDefinitions.basic BasicAuth
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
// @securitydefinitions.oauth2.application OAuth2Application
// @tokenUrl https://example.com/oauth/token
// @scope.write Grants write access
// @scope.admin Grants read and write access to administrative information
// @securitydefinitions.oauth2.implicit OAuth2Implicit
// @authorizationUrl https://example.com/oauth/authorize
// @scope.write Grants write access
// @scope.admin Grants read and write access to administrative information
// @securitydefinitions.oauth2.password OAuth2Password
// @tokenUrl https://example.com/oauth/token
// @scope.read Grants read access
// @scope.write Grants write access
// @scope.admin Grants read and write access to administrative information
// @securitydefinitions.oauth2.accessCode OAuth2AccessCode
// @tokenUrl https://example.com/oauth/token
// @authorizationUrl https://example.com/oauth/authorize
// @scope.admin Grants read and write access to administrative information
func main() {
	router := gin.Default()
	router.Use(cors.Default())
	client, err := ent.Open("sqlite3", "file:ent.db?cache=shared&_fk=1")
	if err != nil {
		log.Fatalf("fail to open sqlite3: %v", err)
	}
	defer client.Close()
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
	v1 := router.Group("/api/v1")
	controllers.NewPatientController(v1, client)
	controllers.NewDoctorController(v1, client)
	controllers.NewMannerController(v1, client)
	controllers.NewDrugAllergyController(v1, client)
	controllers.NewMedicineController(v1, client)
	// Set Doctors Data
	doctors := Doctors{
		Doctor: []Doctor{
			Doctor{"dhetporn@gmail.com","124443","Dhetporn Krongsin","0989185789" },
			Doctor{"mild121014@gmail.com","1222224","Ploymanee Chuyat","0888880988" },
		},
	}
	for _, u := range doctors.Doctor {
		client.Doctor.
			Create().
			SetDoctorEmail(u.DoctorEmail).
			SetDoctorPassword(u.DoctorPassword).
			SetDoctorName(u.DoctorName).
			SetDoctorTel(u.DoctorTel).
			Save(context.Background())
	}
	// Set Patients Data
	patients := Patients{
		Patient: []Patient{
			Patient{"สมสมร สายใจ"},
			Patient{"วายใจ สายสม"},
			Patient{"ยนรพี สนา"},
			Patient{"แอนนา เบล"},
		},
	}
	for _, p := range patients.Patient {
		client.Patient.
			Create().
			SetPatientName(p.PatientName).
			Save(context.Background())
	}
	// Set Medicines Data
	medicines := Medicines{
		Medicine: []Medicine{
			Medicine{"ยาปฎิชีวนะ"},
			Medicine{"ยาแก้อักเสบ"},
			Medicine{"ยานอนหลับ"},
			Medicine{"ยาแก้แพ้"},
		},
	}
	for _, me := range medicines.Medicine {
		client.Medicine.
			Create().
			SetMedicineName(me.MedicineName).
			Save(context.Background())
	}

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.Run()
}