package repo

import (
	"bookingoto-try/features/user"
	"bookingoto-try/helpers"
	"bookingoto-try/models"

	"github.com/jinzhu/gorm"
)

type UserRepoImpl struct {
	DB *gorm.DB
}

func NewUserRepo(DB *gorm.DB) user.UserRepo {
	return &UserRepoImpl{DB}
}

func (r *UserRepoImpl) ReadAllUser() ([]map[string]interface{}, error) {
	result := make([]map[string]interface{}, 0)
	customerMap := make(map[int]map[string]interface{})

	sql := "SELECT c.cst_id, c.cst_name, c.cst_dob, c.cst_phone_num, c.cst_email, n.nationality_name, f.fl_relation, f.fl_name, f.fl_dob FROM customers AS c JOIN nationalities AS n ON c.nationality_id = n.nationality_id LEFT JOIN family_lists AS f ON c.cst_id = f.cst_id"
	rows, err := r.DB.Debug().Raw(sql).Rows()
	if err != nil {
		helpers.Logger("error", "In Server: "+err.Error())
		return nil, err
	}

	for rows.Next() {
		var rest models.Customer
		var restFamilytest models.Familylist
		var restNationality models.Nationality
		r.DB.Debug().ScanRows(rows, &rest)
		r.DB.Debug().ScanRows(rows, &restFamilytest)
		r.DB.Debug().ScanRows(rows, &restNationality)

		customerID := rest.CstId
		keluargaMember := make(map[string]interface{})
		keluargaMember["hubungan"] = restFamilytest.FlRelation
		keluargaMember["nama"] = restFamilytest.FlName
		keluargaMember["tanggal_lahir"] = restFamilytest.FlDob

		if _, exists := customerMap[customerID]; !exists {
			customerMap[customerID] = map[string]interface{}{
				"cst_id":          rest.CstId,
				"nama":            rest.CstName,
				"tanggal_lahir":   rest.CstDob.String(),
				"telepon":         rest.CstPhoneNum,
				"kewarganegaraan": restNationality.NationalityName,
				"email":           rest.CstEmail,
				"keluarga":        []map[string]interface{}{keluargaMember},
			}
		} else {
			customerMap[customerID]["keluarga"] = append(customerMap[customerID]["keluarga"].([]map[string]interface{}), keluargaMember)
		}
	}

	for _, customerData := range customerMap {
		result = append(result, customerData)
	}

	return result, nil
}

func (r *UserRepoImpl) ReadSingleId(cstId string) (map[string]interface{}, error) {
	responseData := make(map[string]interface{})
	var keluarga []map[string]interface{}

	sql := "SELECT c.cst_id, c.cst_name, c.cst_dob, c.cst_phone_num, c.cst_email, n.nationality_name, f.fl_relation, f.fl_name, f.fl_dob FROM customers AS c JOIN nationalities AS n ON c.nationality_id = n.nationality_id LEFT JOIN family_lists AS f ON c.cst_id = f.cst_id WHERE c.cst_id = ?"
	row, err := r.DB.Debug().Raw(sql, cstId).Rows()
	if err != nil {
		helpers.Logger("error", "In Server: "+err.Error())
		return nil, err
	}

	for row.Next() {
		var rest models.Customer
		var restFamilytest models.Familylist
		var restNationality models.Nationality
		r.DB.Debug().ScanRows(row, &rest)
		r.DB.Debug().ScanRows(row, &restFamilytest)
		r.DB.Debug().ScanRows(row, &restNationality)

		keluargaMember := make(map[string]interface{})
		keluargaMember["hubungan"] = restFamilytest.FlRelation
		keluargaMember["nama"] = restFamilytest.FlName
		keluargaMember["tanggal_lahir"] = restFamilytest.FlDob

		keluarga = append(keluarga, keluargaMember)

		responseData["nama"] = rest.CstName
		responseData["tanggal_lahir"] = rest.CstDob.String()
		responseData["telepon"] = rest.CstPhoneNum
		responseData["kewarganegaraan"] = restNationality.NationalityName
		responseData["email"] = rest.CstEmail

	}

	responseData["keluarga"] = keluarga

	return responseData, nil
}
