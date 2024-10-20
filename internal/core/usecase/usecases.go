package usecase

import (
	"github.com/Sreeram-ganesan/my-blog/internal/core/outport"
)

type UseCases struct {
	AddrBook outport.AddrBook
	// other output/secondary ports can be added here
	BlogBook outport.BlogBook
}
