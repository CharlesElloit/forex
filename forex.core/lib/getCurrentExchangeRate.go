package forexCore

import (
	"fmt"
	"time"

	"api.forex.com/models"
	"api.forex.com/storage"
	"github.com/google/uuid"
)

func GetCurrentExchangeRate(currency_id uuid.UUID) (*models.Rate, error) {
	/*
	 * Make sure the currency passed in is valid currency
	 */
	if currency_id != uuid.Nil {
		/*
		 * We need to check if the currency exists in the database before we
		 * we can query the data for current exchange rate at any given requested time.
		 */
		var currency models.Currency
		if currencyExists := storage.DB.Model(&models.Currency{}).Where("id = ?", currency_id).Find(&currency); currencyExists.Error != nil {
			return nil, fmt.Errorf("the currency provided doesn't exists")
		}

		/*
		 * Get current date in YYYY-MM-DD format
		 */
		currentDate := time.Now().Format("2006-01-02")

		var rate models.Rate
		result := storage.DB.Model(&models.Rate{}).
			Where("created_date = ? AND currency_id = ? AND is_approved = ?", currentDate, currency_id, true).
			Order("created_time DESC").
			Limit(1).
			Find(&rate)

		// Check for errors or if no records were found
		if result.Error != nil {
			return nil, fmt.Errorf("query error: %w", result.Error)
		}
		if result.RowsAffected == 0 {
			return nil, fmt.Errorf("no approved rate found for %s", currentDate)
		}

		return &rate, nil
	}
	return nil, fmt.Errorf("please provide a valid currency")
}
