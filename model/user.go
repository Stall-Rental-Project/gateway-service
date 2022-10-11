package model

import (
	"encoding/json"
	grpc "gateway-service/client/common"
	"gateway-service/client/market"
)

type User struct {
	UserId      string       `json:"user_id"`
	FirstName   string       `json:"first_name"`
	MiddleName  string       `json:"middle_name"`
	LastName    string       `json:"last_name"`
	Email       string       `json:"email"`
	Status      grpc.Status  `json:"status"`
	Markets     []UserMarket `json:"markets"`
	Roles       []string     `json:"roles"`
	RoleIds     []string     `json:"role_ids"`
	RoleCodes   []string     `json:"role_codes"`
	Permissions []string     `json:"permissions"`
	MarketCodes []string     `json:"market_codes"`
}

func (s User) MarshalBinary() ([]byte, error) {
	return json.Marshal(s)
}

func (s User) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, &s)
}

type UserMarket struct {
	MarketId string            `json:"market_id"`
	Name     string            `json:"name"`
	Code     string            `json:"code"`
	Type     market.MarketType `json:"type"`
}

func (s UserMarket) MarshalBinary() ([]byte, error) {
	return json.Marshal(s)
}

func (s UserMarket) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, &s)
}

type UserPreference struct {
}

func (s UserPreference) MarshalBinary() ([]byte, error) {
	return json.Marshal(s)
}

func (s UserPreference) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, &s)
}
