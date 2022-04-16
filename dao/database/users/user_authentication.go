package databaseUser

import (
	"context"
	"errors"
	"log"

	database "github.com/akromibn37/hospitalityCollaboration/dao/database"
	apimodel "github.com/akromibn37/hospitalityCollaboration/models/user"
	"github.com/akromibn37/hospitalityCollaboration/pkg/util"
)

type TUserAuthentication struct {
	UserId      string `json:"user_id"`
	UserName    string `json:"user_name"`
	PassWord    string `json:"password"`
	UserType    string `json:"user_type"`
	DisplayName string `json:"display_name"`
	Active      bool   `json:"active"`
	CreateBy    string `json:"create_by"`
	CreateDate  string `json:"create_date"`
	UpdateBy    string `json:"update_by"`
	UpdateDate  string `json:"update_date"`
}

//Login
func Login(req *apimodel.UserLoginRequest) (res *TUserAuthentication, err error) {
	db := database.GetDB()
	statement := "select * from t_user_authentication where user_name=? and password=? and active=1"
	resp := TUserAuthentication{}

	err = db.QueryRow(statement, req.Username, req.PassWord).Scan(&resp.UserId, &resp.UserName, &resp.PassWord, &resp.UserType, &resp.DisplayName, &resp.Active, &resp.CreateBy, &resp.CreateDate, &resp.UpdateBy, &resp.UpdateDate)
	if err != nil {
		log.Print("err:", err)
		return nil, err
	}
	return &resp, nil
}

//Login
func GetByUserId(req *apimodel.UserGetByUserIdRequest) (res *TUserAuthentication, err error) {
	db := database.GetDB()
	statement := "select * from t_user_authentication where user_id=?"
	resp := TUserAuthentication{}

	err = db.QueryRow(statement, req.UserId).Scan(&resp.UserId, &resp.UserName, &resp.PassWord, &resp.UserType, &resp.DisplayName, &resp.Active, &resp.CreateBy, &resp.CreateDate, &resp.UpdateBy, &resp.UpdateDate)
	if err != nil {
		log.Print("err:", err)
		return nil, err
	}
	return &resp, nil
}

//Add New User
func CreateBasicUser(req *apimodel.UserCreateRequest, userid string) (err error) {
	db := database.GetDB()
	ctx := context.Background()
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	statement := "select count(*) from t_user_authentication where user_name=?"
	result := db.QueryRow(statement, req.Username)
	var count int
	err = result.Scan(&count)
	if err != nil {
		return err
	}
	if count != 0 {
		return errors.New("Username repeated")
	}
	statement = "INSERT INTO t_user_authentication (user_id,user_name,password,user_type,display_name,active,create_by,create_date,update_by,update_date) VALUES (?,?,?,?,?,?,?,?,?,?);"
	_, err = tx.Query(statement, userid, req.Username, req.PassWord, req.UserType, req.DisplayName, true, req.CreateBy, util.TimeNow(), req.UpdateBy, util.TimeNow())
	if err != nil {
		tx.Rollback()
		return err
	}
	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}

//GetPassword
func GetPassword(userId string) (res string, err error) {
	db := database.GetDB()
	statement := "select password from t_user_authentication where user_id=?"
	resp := TUserAuthentication{}

	err = db.QueryRow(statement, userId).Scan(&resp.PassWord)
	if err != nil {
		log.Print("err:", err)
		return "", err
	}
	return resp.PassWord, nil
}

//GetAllProfile
func GetAllData() (res []*TUserAuthentication, err error) {
	db := database.GetDB()
	statement := "select * from t_user_authentication"
	resp := make([]*TUserAuthentication, 0)

	results, err := db.Query(statement)
	if err != nil {
		log.Print("err:", err)
		return nil, err
	}
	for results.Next() {
		tmpResp := new(TUserAuthentication)
		err = results.Scan(&tmpResp.UserId, &tmpResp.UserName, &tmpResp.PassWord, &tmpResp.UserType, &tmpResp.DisplayName, &tmpResp.Active, &tmpResp.CreateBy, &tmpResp.CreateDate, &tmpResp.UpdateBy, &tmpResp.UpdateDate)
		if err != nil {
			log.Print("err:", err)
			return nil, err
		}
		resp = append(resp, tmpResp)
	}
	return resp, nil
}

//GetData
func GetData(userId string) (res *TUserAuthentication, err error) {
	db := database.GetDB()
	statement := "select user_type,display_name,active from t_user_authentication where user_id=?"
	resp := TUserAuthentication{}

	err = db.QueryRow(statement, userId).Scan(&resp.UserType, &resp.DisplayName, &resp.Active)
	if err != nil {
		log.Print("err:", err)
		return nil, err
	}
	return &resp, nil
}

//UpdatePassword
func UpdatePassword(req *apimodel.UserUpdatePasswordRequest) (err error) {
	db := database.GetDB()
	ctx := context.Background()
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	statement := "UPDATE t_user_authentication SET password=?,update_date=? WHERE user_id=?;"
	_, err = tx.Query(statement, req.NewPassword, util.TimeNow(), req.UserId)
	if err != nil {
		tx.Rollback()
		return err
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}

//UpdateData
func UpdateData(req *apimodel.UserUpdateDataRequest) (err error) {
	db := database.GetDB()
	ctx := context.Background()
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	statement := "UPDATE t_user_authentication SET user_type=?,display_name=?,active=?,update_date=? WHERE user_id=?;"
	_, err = tx.Query(statement, req.UserType, req.DisplayName, req.Active, util.TimeNow(), req.UserId)
	if err != nil {
		tx.Rollback()
		return err
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}
