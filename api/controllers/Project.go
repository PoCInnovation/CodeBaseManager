package controllers

import (
	"cbm-api/database"
	"cbm-api/models"
)

func FindProject(db database.Database, name string) *models.Project {
	project := models.Project{
		Name: name,
	}
	result := db.DB.First(&project)
	if result.Error != nil {
		return nil
	}
	return &project
}

//func (s *Server) CreateProject(w http.ResponseWriter, r *http.Request) {
//	body, err := ioutil.ReadAll(r.Body)
//	if err != nil {
//		responses.ERROR(w, http.StatusUnprocessableEntity, err)
//	}
//	proj := models.Project{}
//	err = json.Unmarshal(body, &proj)
//	if err != nil {
//		responses.ERROR(w, http.StatusUnprocessableEntity, err)
//		return
//	}
//	userCreated, err := proj.Save(s.DB.DB)
//	if err != nil {
//		responses.ERROR(w, http.StatusInternalServerError,
//			err)
//		return
//	}
//	w.Header().Set("Location", fmt.Sprintf("%s%s/%d", r.Host, r.RequestURI, userCreated.ID))
//	responses.JSON(w, http.StatusCreated, userCreated)
//}
//
//func (s *Server) GetUser(w http.ResponseWriter, r *http.Request) {
//	vars := mux.Vars(r)
//	projName := vars["name"]
//	if projName == "" {
//		responses.ERROR(w, http.StatusBadRequest, errors.New("no project name"))
//		return
//	}
//	proj := models.Project{}
//	projGotten, err := proj.FindProjectByName(s.DB.DB, projName)
//	if err != nil {
//		responses.ERROR(w, http.StatusBadRequest, err)
//		return
//	}
//	responses.JSON(w, http.StatusOK, projGotten)
//}
