package engine

import (
	"log"

	"github.com/PicoOrg/gopoc/model"
	cve_2022_46463 "github.com/PicoOrg/gopoc/poc/Harbor/CVE-2022-46463"
)

var (
	pocs map[string]func(...model.Option) model.IPoc
)

func init() {
	pocs = map[string]func(...model.Option) model.IPoc{
		"CVE-2022-46463": cve_2022_46463.NewPoc,
	}
}

func Run(options ...model.Option) {
	for _, poc := range pocs {
		p := poc(options...)
		if p.Verify() {
			log.Printf("success: %+v", p)
		}
	}
}
