package mantras

import (
	"encoding/json"
	"errors"
	"log"
	"math/rand"
	"os"

	"github.com/rs/xid"
)

type Mantra struct {
	Text string `json:"text"`
	ID   string `json:"id"`
}

var (
	mantras          = make([]Mantra, 0)
	ErrMissingText   = errors.New("Mantras must have text, got an empty string.")
	ErrMissingID     = errors.New("Must specify a mantra ID to delete, got an empty string.")
	ErrNotFound      = errors.New("Mantra with that ID not found.")
	MANTRA_SAVE_PATH = "./data/mantras.json"
)

func init() {
	data, err := os.ReadFile(MANTRA_SAVE_PATH)

	if err != nil {
		return
	}

	err = json.Unmarshal(data, &mantras)

	if err != nil {
		log.Fatal("Error parsing mantras!", err)
	}
}

func saveMantras() {
	data, err := json.Marshal(mantras)
	if err != nil {
		log.Fatal("Failed to marshal mantras: ", err)
	}

	err = os.WriteFile(MANTRA_SAVE_PATH, data, 0640)
	if err != nil {
		log.Fatal("Failed to save mantras: ", err)
	}
}

func GetMantras() []Mantra {
	return mantras
}

func AddMantra(text string) error {
	// don't add blank mantras
	if text == "" {
		return ErrMissingText
	}

	mantras = append(mantras, Mantra{
		Text: text,
		ID:   xid.New().String(),
	})

	saveMantras()

	return nil
}

func UpdateMantra(text, id string) error {
	if text == "" {
		return ErrMissingText
	}

	foundMantra := false

	for i, m := range mantras {
		if m.ID == id {
			foundMantra = true
			m.Text = text
			mantras[i] = m
		}
	}

	if !foundMantra {
		return ErrNotFound
	}

	saveMantras()
	return nil
}

func DeleteMantra(id string) error {
	if id == "" {
		return ErrMissingID
	}

	m := make([]Mantra, 0)

	for _, mantra := range mantras {
		if mantra.ID != id {
			m = append(m, mantra)
		}
	}

	mantras = m

	saveMantras()
	return nil
}

// get a random mantra, the bool signifies if there actually are mantras,
// if the bool is false no message should be sent, there's nothing to say
func GetRandomMantra() (Mantra, bool) {
	if len(mantras) == 0 {
		return Mantra{}, false
	}
	i := rand.Intn(len(mantras))
	return mantras[i], true
}
