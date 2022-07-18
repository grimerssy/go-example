package id

import (
	"math"

	"github.com/pjebs/optimus-go"
)

type ConfigOptimus struct {
	Prime  uint64
	Random uint64
}

type Optimus struct {
	core optimus.Optimus
}

func NewOptimus(cfg ConfigOptimus) *Optimus {
	optimus.MAX_INT = math.MaxInt64

	return &Optimus{
		core: optimus.NewCalculated(cfg.Prime, cfg.Random),
	}
}

func (o *Optimus) ObfuscateId(id int64) (int64, error) {
	return int64(o.core.Encode(uint64(id))), nil
}

func (o *Optimus) DeobfuscateId(obfuscated int64) (int64, error) {
	return int64(o.core.Decode(uint64(obfuscated))), nil
}
