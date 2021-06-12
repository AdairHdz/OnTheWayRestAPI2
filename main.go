package main

import (		
	"github.com/AdairHdz/OnTheWayRestAPI/BusinessLayer/businessEntities"
	"github.com/AdairHdz/OnTheWayRestAPI/DataLayer/database"
	"github.com/AdairHdz/OnTheWayRestAPI/ServicesLayer/routes"	
	uuid "github.com/satori/go.uuid"
)


func main() {	
	
	DB := database.GetDatabase()
	DB.SetupJoinTable(&businessEntities.PriceRate{}, "WorkingDays", &businessEntities.PriceRateWorkingDay{})
	DB.AutoMigrate(
		&businessEntities.User{},
		&businessEntities.WorkingDay{},
		&businessEntities.State{},
		&businessEntities.City{},
		&businessEntities.ServiceProvider{},	
		&businessEntities.PriceRate{},
		&businessEntities.ServiceRequest{},		
		&businessEntities.Review{},
		&businessEntities.ReviewEvidence{},
		&businessEntities.ServiceRequester{},
		&businessEntities.Address{})	

	var workingDays [7]businessEntities.WorkingDay
	workingDays[0] = businessEntities.WorkingDay{
		ID: 1,
		Name: "Lunes",
	}
	
	workingDays[1] = businessEntities.WorkingDay{		
		ID: 2,
		Name: "Martes",
	}
	workingDays[2] = businessEntities.WorkingDay{
		ID: 3,
		Name: "Miércoles",
	}
	workingDays[3] = businessEntities.WorkingDay{
		ID: 4,
		Name: "Jueves",
	}
	workingDays[4] = businessEntities.WorkingDay{
		ID: 5,
		Name: "Viernes",
	}
	workingDays[5] = businessEntities.WorkingDay{
		ID: 6,
		Name: "Sábado",
	}
	workingDays[6] = businessEntities.WorkingDay{		
		ID: 7,
		Name: "Domingo",
	}

	DB.Save(&workingDays)

	var states [3]businessEntities.State
	veracruzID := uuid.FromStringOrNil("3d5b56f8-0e6c-495d-b010-196f26d87e48")
	jaliscoID := uuid.FromStringOrNil("51fda705-4fc6-481e-bcab-c34ae7144b54")
	nuevoLeonID := uuid.FromStringOrNil("a2064dcf-ebfe-4203-ae5a-f5f8e870bd02")
	states[0] = businessEntities.State{
		ID: veracruzID,
		Name: "Veracruz",
	}

	states[1] = businessEntities.State{
		ID: jaliscoID,
		Name: "Jalisco",
	}

	states[2] = businessEntities.State{
		ID: nuevoLeonID,
		Name: "Nuevo León",
	}

	DB.Save(&states)

	var cities [3]businessEntities.City

	coatepecID := uuid.FromStringOrNil("a61e6b22-da70-4a04-92e6-4d80342d85c4")
	xalapaID := uuid.FromStringOrNil("1bdbee8b-941f-475a-886c-efa902fefc1f")
	xicoID := uuid.FromStringOrNil("b4926200-bb8b-45c0-80cc-f4c40b5fb945")

	cities[0] = businessEntities.City{
		ID: xalapaID,
		Name: "Xalapa",
		StateID: veracruzID,
	}
	
	cities[1] = businessEntities.City{
		ID: coatepecID,
		Name: "Coatepec",
		StateID: veracruzID,
	}

	cities[2] = businessEntities.City{
		ID: xicoID,
		Name: "Xico",
		StateID: veracruzID,
	}

	DB.Save(&cities)		

	routes.StartServer()	
}