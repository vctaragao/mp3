package frame

import (
	"errors"
	"fmt"
)

type Identifier int

const (
	TIT2 = iota // Title/songname/content description]
	TPE1        // Lead performer(s)/Soloist(s)]
	TYER        // Year
	TCON        // Content type
	TENC        // Encoded by
	TBPM        // BPM (Beats per minute)
)

func (i Identifier) String() string {
	return []string{"TIT2", "TPE1", "TYER", "TCON", "TENC", "TBPM"}[i]
}

func IdentifierFromString(identifier string) (Identifier, error) {
	id, ok := map[string]Identifier{
		"TIT2": TIT2,
		"TPE1": TPE1,
		"TYER": TYER,
		"TCON": TCON,
		"TENC": TENC,
		"TBPM": TBPM,
	}[identifier]
	if !ok {
		return 0, errors.New(fmt.Sprintf("invalid or unmapped identifier: %s", identifier))
	}

	return id, nil
}
