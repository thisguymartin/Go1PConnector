package pkg

import (
	"log"
	"os"

	"github.com/1Password/connect-sdk-go/connect"
)

func OnePasswordConnect(uuid string, vaultUUID string) connect.Client {
	var client, err = connect.NewClientFromEnvironment()
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	return client
}

func GetAllVaules(passwordConnect connect.Client) {

	vaults, error := passwordConnect.GetVaults()

	if error != nil {
		log.Fatal(error)
		os.Exit(1)
	}

	for _, v := range vaults {
		print(v.ID)
		print(v.Name)
	}

}
