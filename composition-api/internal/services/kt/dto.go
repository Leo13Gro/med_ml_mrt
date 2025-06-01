package kt

import (
	ht "github.com/ogen-go/ogen/http"
)

type CreateKtArg struct {
	File        ht.MultipartFile
	Description *string
}
