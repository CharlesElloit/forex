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
	"api.forex.com/models"
	"api.forex.com/storage"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func GetAndHandleRateExists(
	rate *models.Rate,
	selling_rate float32,
	buying_rate float32,
	average_rate float32,
	currency_id uuid.UUID,
) (bool, error) {
	var rateExistsQuery *gorm.DB
	if rateExistsQuery = storage.DB.
		Where("selling_rate = ? AND buying_rate = ? AND average_rate = ? AND currency_id = ? AND created_date = current_date",
			selling_rate, buying_rate, average_rate, currency_id).
		Find(&rate); rateExistsQuery.Error != nil {
		return false, rateExistsQuery.Error
	}
	if rateExistsQuery.RowsAffected > 0 {
		return true, nil
	}
	return false, nil
}
