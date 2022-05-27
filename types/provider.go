package types

import "encoding/json"

type Provider int64

const (
	UndefinedProvider Provider = iota
	Google
	Apple
)

func (p Provider) String() string {
	switch p {
	case Google:
		return "google"
	case Apple:
		return "apple"
	}
	return ""
}

func (p Provider) EnumIndex() int {
	return int(p)
}

func (p Provider) IsNil() bool {
	if p == UndefinedProvider {
		return true
	} else {
		return false
	}
}

func (p Provider) FromString(provider string) Provider {
	return map[string]Provider{
		"apple":  Apple,
		"google": Google,
		"":       UndefinedProvider,
	}[provider]
}

func ProviderFromString(provider string) Provider {
	return map[string]Provider{
		"apple":  Apple,
		"google": Google,
		"":       UndefinedProvider,
	}[provider]
}

func (p Provider) MarshalJSON() ([]byte, error) {
	return json.Marshal(p.String())
}

func (p *Provider) UnmarshalJSON(b []byte) error {
	var s string
	err := json.Unmarshal(b, &s)
	if err != nil {
		return err
	}
	*p = p.FromString(s)
	return nil
}
