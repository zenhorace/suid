package suid

import (
	"crypto/rand"
	"math"
	"math/big"
	"time"

	mr "math/rand"
)

type SUID struct {
	base int
}

// New creates an SUID instance which can be used to generate short UUIDs.
func New() SUID {
	return SUID{base: 62}
}

// NewCaseInsensitive returns a SUID instance which will generate
// case insensitive IDs. By default, IDs "ABC1" and "abc1" would be
// considered different. But, if "ABC1" == "abc1" for your purposes,
// use this version of SUID, which will only use lower case letters.
func NewCaseInsensitive() SUID {
	return SUID{base: 36}
}

func (s SUID) New128Bit() (string, error) {
	max := &big.Int{}
	max.Exp(big.NewInt(2), big.NewInt(128), nil)
	return suid(s.base, max)
}

func (s SUID) New64Bit() (string, error) {
	max := &big.Int{}
	max.SetUint64(math.MaxUint64)
	return suid(s.base, max)
}

func (s SUID) TimeBasedID() string {
	val := big.Int{}
	val.Add(big.NewInt(time.Now().UnixNano()), big.NewInt(mr.Int63()))
	return val.Text(s.base)
}

func (s SUID) New64BitWithFallback() string {
	id, err := s.New64Bit()
	if err != nil {
		return s.TimeBasedID()
	}
	return id
}

func suid(base int, max *big.Int) (string, error) {
	r, err := rand.Int(rand.Reader, max)
	if err != nil {
		return "", err
	}
	return r.Text(base), nil
}
