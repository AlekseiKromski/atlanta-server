package gin_server

import (
	"alekseikromski.com/atlanta/core"
	"alekseikromski.com/atlanta/modules/storage"
	"fmt"
)

func (s *Server) Require() []string {
	return []string{
		"storage",
	}
}

func (s *Server) getStorageFromRequirement(requirements map[string]core.Module) (storage.Storage, error) {
	storage, ok := requirements["storage"].(storage.Storage)
	if !ok {
		s.Log("All requirements list")
		for k, v := range requirements {
			s.Log("Requirement", k, v.Signature())
		}
		return nil, fmt.Errorf("requiremnt list has wrong storage requirement")
	}

	return storage, nil
}
