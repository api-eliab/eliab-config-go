package eliab

import (
	"testing"
)

func TestSqlConnect(t *testing.T) {
	if !dbConnect() {
		t.Error("No se ha conectado a la base de datos!")
	}
}

// func TestQuery(t *testing.T) {
// 	query := fmt.Sprintf("SELECT id FROM users WHERE email = @email AND password = @password")
// 	query, err := getQueryString(
// 		query,
// 		sql.Named("email", "email"),
// 		sql.Named("password", time.Now()),
// 	)
// 	if err != nil {
// 		t.Error(err)
// 		t.Error(query)
// 	}

// 	t.Log(query)
// }
