package core

import (
	"fmt"
	"github.com/MiracleOS-Team/abg/v2/settings"
)

var abg *Abg

type Abg struct {
	Cnf *settings.Config
}

func NewApx(cnf *settings.Config) *Abg {
	apx = &Abg{
		Cnf: cnf,
	}

	err := apx.EssentialChecks()
	if err != nil {
		// localisation features aren't available at this stage, so this error can't be translated
		fmt.Println("ERROR: Unable to find apx configuration files")
		return nil
	}

	return apx
}

func NewStandardApx() *Abg {
	cnf, err := settings.GetApxDefaultConfig()
	if err != nil {
		panic(err)
	}

	apx = &Abg{
		Cnf: cnf,
	}

	err = apx.EssentialChecks()
	if err != nil {
		// localisation features aren't available at this stage, so this error can't be translated
		fmt.Println("ERROR: Unable to find apx configuration files")
		return nil
	}
	return apx
}
