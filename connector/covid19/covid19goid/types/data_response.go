package types

// DataResponse represent Data Response
type DataResponse struct {
	LastUpdate string    `json:"last_update"`
	Kasus      Kasus     `json:"kasus"`
	Sembuh     Meninggal `json:"sembuh"`
	Meninggal  Meninggal `json:"meninggal"`
	Perawatan  Meninggal `json:"perawatan"`
}

// Kasus represent cases (Indonesian lang.)
type Kasus struct {
	KondisiPenyerta GejalaClass       `json:"kondisi_penyerta"`
	JenisKelamin    GejalaClass       `json:"jenis_kelamin"`
	KelompokUmur    KasusKelompokUmur `json:"kelompok_umur"`
	Gejala          GejalaClass       `json:"gejala"`
}

// GejalaClass represent indication class (Indonesian lang.)
type GejalaClass struct {
	CurrentData int64             `json:"current_data"`
	MissingData float64           `json:"missing_data"`
	ListData    []GejalaListDatum `json:"list_data"`
}

// GejalaListDatum represent indication list (Indonesian lang.)
type GejalaListDatum struct {
	Key      string  `json:"key"`
	DocCount float64 `json:"doc_count"`
}

// KasusKelompokUmur represent cases grouped by age (Indonesian lang.)
type KasusKelompokUmur struct {
	CurrentData int64                   `json:"current_data"`
	MissingData float64                 `json:"missing_data"`
	ListData    []KelompokUmurListDatum `json:"list_data"`
}

// KelompokUmurListDatum represent list of cases grouped by age (Indonesian lang.)
type KelompokUmurListDatum struct {
	Key      string  `json:"key"`
	DocCount float64 `json:"doc_count"`
	Usia     Usia    `json:"usia"`
}

// Meninggal represent Death (Indonesian lang.)
type Meninggal struct {
	KondisiPenyerta PurpleGejala          `json:"kondisi_penyerta"`
	JenisKelamin    PurpleGejala          `json:"jenis_kelamin"`
	KelompokUmur    MeninggalKelompokUmur `json:"kelompok_umur"`
	Gejala          PurpleGejala          `json:"gejala"`
}

// PurpleGejala represent cases (Indonesian lang.)
type PurpleGejala struct {
	ListData []GejalaListDatum `json:"list_data"`
}

// MeninggalKelompokUmur represent death grouped by age (Indonesian lang.)
type MeninggalKelompokUmur struct {
	ListData []KelompokUmurListDatum `json:"list_data"`
}
