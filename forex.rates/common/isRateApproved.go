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
package common

import (
	forexCore "api.forex.com/forex.core/common"
	"api.forex.com/models"
	"api.forex.com/storage"
	"api.forex.com/utilities"
	"github.com/google/uuid"
	"github.com/kataras/iris/v12"
)

func IsRateApproved(rateType forexCore.EnumType, currency_id uuid.UUID, ctx iris.Context) bool {
	rateMode := forexCore.EnumType(rateType)
	if !rateMode.IsValid() {
		utilities.CreateError(iris.StatusBadRequest, "Incorrect rate mode", "Incorrect rate mode", ctx)
	}
	//make a query to the database checking if the rate exists and approved.
	var rate models.Rate
	query := storage.DB.Model(&models.Rate{}).
		Where("created_date = ? AND currency_id = ? AND rate_mode = ? AND is_approved = ?",
			forexCore.CurrentDate, currency_id, rateMode, false).
		Order("created_time DESC").Limit(1).Find(&rate)

	return query.RowsAffected > 0
}
