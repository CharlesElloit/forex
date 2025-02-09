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
	"github.com/google/uuid"
)

type AddRateInput struct {
	SellingRate         float32   `json:"selling_rate"`
	SellingTransferRate float32   `json:"selling_transfer_rate"`
	BuyingRate          float32   `json:"buying_rate"`
	BuyingTransferRate  float32   `json:"buying_transfer_rate"`
	AverageRate         float32   `json:"average_rate"`
	MidRate             float32   `json:"mid_rate"`
	RateMode            string    `json:"rate_mode" validate:"required"`
	CurrencyID          uuid.UUID `json:"currency_id" validate:"required"`
}
