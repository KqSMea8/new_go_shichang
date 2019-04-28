package api

import "github.com/gin-gonic/gin"

// One vendor with all goods price
func CreateVendorPriceList(ctx *gin.Context) {
}

// One vendor with all goods price
func UpdateVendorPriceList(ctx *gin.Context) {
}

// update specified vendors with specified goods price
func UpdateAllPrice(ctx *gin.Context) {
	// will not update the price with 0
	// if param isAll = "all" then update all vendors
}