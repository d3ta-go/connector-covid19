package types

// ProvResponse represent province Response
type ProvResponse struct {
	LastDate      string      `json:"last_date"`
	CurrentData   float64     `json:"current_data"`
	MissingData   float64     `json:"missing_data"`
	TanpaProvinsi int64       `json:"tanpa_provinsi"`
	ListData      []ListDatum `json:"list_data"`
}

// ListDatum represent general list information (Indonesian lang.)
type ListDatum struct {
	Key             string         `json:"key"`
	DocCount        float64        `json:"doc_count"`
	JumlahKasus     int64          `json:"jumlah_kasus"`
	JumlahSembuh    int64          `json:"jumlah_sembuh"`
	JumlahMeninggal int64          `json:"jumlah_meninggal"`
	JumlahDirawat   int64          `json:"jumlah_dirawat"`
	JenisKelamin    []JenisKelamin `json:"jenis_kelamin"`
	KelompokUmur    []KelompokUmur `json:"kelompok_umur"`
	Lokasi          Lokasi         `json:"lokasi"`
	Penambahan      PenambahanProv `json:"penambahan"`
}

// JenisKelamin represent gender/sex type (Indonesian lang.)
type JenisKelamin struct {
	Key      JenisKelaminKey `json:"key"`
	DocCount int64           `json:"doc_count"`
}

// KelompokUmur represent group by age (Indonesian lang.)
type KelompokUmur struct {
	Key      KelompokUmurKey `json:"key"`
	DocCount int64           `json:"doc_count"`
	Usia     Usia            `json:"usia"`
}

// Lokasi represent location (Indonesian lang.)
type Lokasi struct {
	Lon float64 `json:"lon"`
	Lat float64 `json:"lat"`
}

// PenambahanProv represent new cases by province (Indonesian lang.)
type PenambahanProv struct {
	Positif   int64 `json:"positif"`
	Sembuh    int64 `json:"sembuh"`
	Meninggal int64 `json:"meninggal"`
}

// JenisKelaminKey represent gender/sex type key (Indonesian lang.)
type JenisKelaminKey string

const (
	// LakiLaki represent Male (Indonesian lang.)
	LakiLaki JenisKelaminKey = "LAKI-LAKI"
	// Perempuan represent Female (Indonesian lang.)
	Perempuan JenisKelaminKey = "PEREMPUAN"
)

// KelompokUmurKey represent group by age key (Indonesian lang.)
type KelompokUmurKey string

const (
	// The05 represent 0-5 age range (Indonesian lang.)
	The05 KelompokUmurKey = "0-5"
	// The617 represent 6-17 age range (Indonesian lang.)
	The617 KelompokUmurKey = "6-17"
	// The1830 represent 18-30 age range (Indonesian lang.)
	The1830 KelompokUmurKey = "18-30"
	// The3145 represent 31-45 age range (Indonesian lang.)
	The3145 KelompokUmurKey = "31-45"
	// The4659 represent 46-59 age range (Indonesian lang.)
	The4659 KelompokUmurKey = "46-59"
	// The60 represent >=60 age range (Indonesian lang.)
	The60 KelompokUmurKey = "≥ 60"
)
