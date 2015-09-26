package model

import (
	"bytes"
	"encoding/base32"
	"github.com/macedo/category_service-go/Godeps/_workspace/src/code.google.com/p/go-uuid/uuid"
)

var encoding = base32.NewEncoding("ybndrfg8ejkmcpqxot1uwisza345h769")

func NewId() string {
	var b bytes.Buffer
	encoder := base32.NewEncoder(encoding, &b)
	encoder.Write(uuid.NewRandom())
	encoder.Close()
	b.Truncate(26)

	return b.String()
}
