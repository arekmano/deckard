package transaction

import (
	"io"
	"net/http"
)

func HttpTransaction(req *http.Request) Transaction {
	return func(context *TransactionContext) error {
		res, err := http.DefaultClient.Do(req)
		if err != nil {
			return err
		}
		_, err = io.Copy(context.TransactionWriter, res.Body)
		return err
	}
}
