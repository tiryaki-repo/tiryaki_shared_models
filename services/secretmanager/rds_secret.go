package secretmanager

import (
	"context"
	"encoding/json"
)

type RDSSecret struct {
	Password string `json:"password"`
	Username string `json:"username"`
	Host     string `json:"host"`
	Port     int    `json:"port"`
	Engine   string `json:"engine"`
}

func (s *SecretManager) GetRDSSecret(ctx context.Context, sn string) (RDSSecret, error) {
	var secret = RDSSecret{}
	secretStr, err := s.getSecret(ctx, sn)
	if err != nil {
		return RDSSecret{}, err
	}

	err = json.Unmarshal([]byte(*secretStr), &secret)
	if err != nil {
		return RDSSecret{}, err
	}

	return secret, err
}
