package data

type Params struct {
	AccentPhrases      []AccentPhrases `json:"accent_phrases"`
	SpeedScale         float64         `json:"speedScale"`
	PitchScale         float64         `json:"pitchScale"`
	IntonationScale    float64         `json:"intonationScale"`
	VolumeScale        float64         `json:"volumeScale"`
	PrePhonemeLength   float64         `json:"prePhonemeLength"`
	PostPhonemeLength  float64         `json:"postPhonemeLength"`
	OutputSamplingRate int             `json:"outputSamplingRate"`
	OutputStereo       bool            `json:"outputStereo"`
	Kana               string          `json:"kana"`
}

type Mora struct {
	Text            string   `json:"text"`
	Consonant       *string  `json:"consonant"`
	ConsonantLength *float64 `json:"consonant_length"`
	Vowel           string   `json:"vowel"`
	VowelLength     float64  `json:"vowel_length"`
	Pitch           float64  `json:"pitch"`
}

type AccentPhrases struct {
	Moras           []Mora `json:"moras"`
	Accent          int    `json:"accent"`
	PauseMora       *Mora  `json:"pause_mora"`
	IsInterrogative bool   `json:"is_interrogative"`
}

type Speakers []struct {
	Name        string   `json:"name"`
	SpeakerUUID string   `json:"speaker_uuid"`
	Styles      []Styles `json:"styles"`
	Version     string   `json:"version"`
}

type Styles struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}
