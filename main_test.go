package KeDuBack

import (
	"testing"
)

func TestHerror_handling(t *testing.T) {
	var MongoURL string
	var SECRET string
	if herror_handling(&MongoURL, &SECRET) == 1 {
		t.Errorf("Error: herror_handling")
	}
}

func TestConnectDB(t *testing.T) {
	var MongoURL string
	SECRET := "secret"
	if herror_handling(&MongoURL, &SECRET) == 1 {
		t.Errorf("Error: herror_handling")
	}
	if connectDB(MongoURL) == nil {
		t.Errorf("Error: connectDB")
	}
}

func TestStartServer(t *testing.T) {
	startServer()
}

func TestMain(m *testing.M) {
	m.Run()
}
