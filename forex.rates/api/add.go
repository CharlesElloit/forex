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

	forexCoreCommon "api.forex.com/forex.core/common"
	forexCore "api.forex.com/forex.core/lib"
	forexCurrency "api.forex.com/forex.currency/helper"
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

	// check if the currency is acitve then proceed otherwise don't
	active, activeErr := forexCurrency.IsCurrencyActive(input.CurrencyID)
	if activeErr != nil {
		utilities.CreateInternalServerError(ctx)
		return
	}
	if !active {
		utilities.CreateError(iris.StatusBadRequest, "Inactive currency", "inactive currency", ctx)
		return
	}

	/*
	* Avoid adding the same rates on the same day as the provious added rate.
	* baiscally because it doesn't make sense to add a new rate which is the
	* same as the previous one.
	 */
	var newRate models.Rate
	if rateExists, err := common.GetAndHandleRateExists(
		&newRate, input.SellingRate, input.BuyingRate, input.CurrencyID,
	); err != nil {
		utilities.CreateInternalServerError(ctx)
		return
	} else if rateExists {
		utilities.CreateError(iris.StatusConflict, "Conflict", "Rate already exists", ctx)
		return
	}

	/*
	 * Don't allow new rate to be added when the pervious added rate is not yet passed and
	 * TODO: also it shouldn't allow adding new rate after a mid rate has been added and approved for
	 * that particular currency.
	 */
	if common.IsRateApproved(forexCoreCommon.EnumType(input.RateMode), input.CurrencyID, ctx) {
		utilities.CreateError(
			iris.StatusNotAcceptable, "Rate insertion error", "either previous rate not passed or mid rate already exists", ctx,
		)
		return
	}

	sequence, err := utilities.GenerateSquenceNumber(&models.Rate{})
	if err != nil {
		utilities.CreateInternalServerError(ctx)
		return
	}

	/*
	 * Check if the average rate is passed in, if not calculate the average rate and store
	 * that instead of storing zero
	 * TODO: Clarity check: if the rate can be less than zero.
	 */
	if input.AverageRate == 0 {
		input.AverageRate = forexCore.CalculateAveragegRate(input.SellingRate, input.BuyingRate)
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
		"id":           newRate.ID,
		"srno":         newRate.Srno,
		"selling_rate": newRate.SellingRate,
		"buying_rate":  newRate.BuyingRate,
		"average_rate": newRate.AverageRate,
		"mid_rate":     newRate.MidRate,
		"rate_mode":    newRate.RateMode,
		"currency_id":  newRate.CurrencyID,
		"created_date": newRate.CreatedDate,
		"created_time": newRate.CreatedTime,
		"createdby_id": newRate.CreatedbyID,
	})
}
