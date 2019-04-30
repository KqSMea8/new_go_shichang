package data

import "xinxin/service/model"

// import multi days basic data
// api/excel/UploadFile
func ImportData(details []model.BasicDailyData){
	
	// insert into basic, replace all old records
	
	InsertChangedDate(tx,details)
	NotifyReporter()
}

// api/data/ListDataByDateRange
func ListDataByDateRange() {
}

// api/data/DeleteData
func DeleteData(details []model.BasicDailyData) {
	
	// delete from basic
	
	InsertChangedDate(tx,details)
}

// api/data/UpdateData
func UpdateData(details []model.BasicDailyData) {

	// replace to basic
	
	InsertChangedDate(tx,details)
	NotifyReporter()
}