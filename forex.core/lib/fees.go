func CreateTransaction(userID string, fromCurrency, toCurrency string, amount float64) (Transaction, error)
func GetTransactionByID(txID string) (Transaction, error)
func ListUserTransactions(userID string) ([]Transaction, error)
func CancelTransaction(txID string) error
func CalculateTransactionFee(amount float64, currency string) (float64, error)