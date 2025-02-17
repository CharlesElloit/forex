/**
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements. See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership. The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License. You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the License is distributed on an
 * "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 * KIND, either express or implied. See the License for the
 * specific language governing permissions and limitations
 * under the License.
 */
package forexRate

import (
	"strings"

	"api.forex.com/forex.rates/common"
	"api.forex.com/models"
	"api.forex.com/storage"
	"api.forex.com/utilities"
	"github.com/google/uuid"
	"github.com/kataras/iris/v12"
)

func Add(ctx iris.Context) {
	//check if the user performing this operation has adequate permissions.
	var input common.AddRateInput
	if err := ctx.ReadJSON(&input); err != nil {
		utilities.HandleValidationError(err, ctx)
		return
	}

	/*
	* Avoid adding the same rates on the same day as the provious added rate.
	* baiscally because it doesn't make sense to add a new rate which is the
	* same as the previous one.
	 */
	var newRate models.Rate
	if rateExists, err := common.GetAndHandleRateExists(&newRate, input.SellingRate, input.BuyingRate, input.AverageRate, input.CurrencyID); err != nil {
		utilities.CreateInternalServerError(ctx)
		return
	} else if rateExists {
		utilities.CreateError(iris.StatusConflict, "Conflict", "Rate already exists", ctx)
		return
	}

	/*
	 * TODO:
	 * Don't allow new rate to be added when the pervious added ra3320te is not yet passed
	 * and also it shouldn't allow adding new rate after a mid rate has been added for
	 * that particular currency.
	 */
	sequence, err := utilities.GenerateSquenceNumber(&models.Rate{})
	if err != nil {
		utilities.CreateInternalServerError(ctx)
		return
	}

	// If we reach here meaning the rate doesn't exists and we'll need to add it.
	// TODO: replace the createdbyID with the currently login user id.
	newRate = models.Rate{
		Srno:                sequence,
		SellingRate:         input.SellingRate,
		SellingTransferRate: input.SellingTransferRate,
		BuyingRate:          input.BuyingRate,
		BuyingTransferRate:  input.BuyingTransferRate,
		AverageRate:         input.AverageRate,
		MidRate:             input.MidRate,
		RateMode:            strings.ToUpper(input.RateMode),
		CurrencyID:          input.CurrencyID,
		CreatedbyID:         uuid.MustParse("84052d72-bf7e-4767-8a74-13c19446d1ff"),
	}

	storage.DB.Create(&newRate)
	ctx.JSON(iris.Map{
		"ID":          newRate.ID,
		"SellingRate": newRate.SellingRate,
		"BuyingRate":  newRate.BuyingRate,
		"AverageRate": newRate.AverageRate,
		"MidRate":     newRate.MidRate,
		"RateMode":    newRate.RateMode,
		"CurrencyID":  newRate.CurrencyID,
		"CreatedDate": newRate.CreatedDate,
		"CreatedTime": newRate.CreatedTime,
		"CreatedbyID": newRate.CreatedbyID,
	})
}
