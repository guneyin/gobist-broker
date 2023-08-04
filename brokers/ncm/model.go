package ncm

type parseDTO struct {
	UserName  string `json:"username"`
	Password  string `json:"password"`
	OtpCode   string `json:"otpCode"`
	DateBegin string `json:"dateBegin"`
	DateEnd   string `json:"dateEnd"`
	Token     string `json:"token"`
}

type loginRequestDTO struct {
	UserName string `json:"username"`
	Password string `json:"password"`
	OtpCode  string `json:"otpCode"`
}

type loginSuccessDTO struct {
	Data struct {
		AccessToken string `json:"accessToken"`
		ExpiresIn   int    `json:"expiresIn"`
		UserData    struct {
			SicilNo  string `json:"sicilNo"`
			Ad       string `json:"ad"`
			Soyad    string `json:"soyad"`
			SonGiris string `json:"sonGiris"`
			KalanGun int    `json:"kalanGun"`
		} `json:"userData"`
	} `json:"data"`
	Success    bool   `json:"success"`
	Message    string `json:"message"`
	StatusCode int    `json:"statusCode"`
}

type transactionRequestDTO struct {
	SymbolName      string `json:"symbolName"`
	FROMISLEMTARIHI string `json:"FROM_ISLEMTARIHI"`
	TOISLEMTARIHI   string `json:"TO_ISLEMTARIHI"`
	MENKULNO        int    `json:"MENKUL_NO"`
}

type transactionResponseDTO struct {
	Data struct {
		R1 []struct {
			ISLEMTARIHI      string  `json:"ISLEM_TARIHI"`
			SEANS            int     `json:"SEANS"`
			MENKULKODU       string  `json:"MENKUL_KODU"`
			EMIR             string  `json:"EMIR"`
			FIYAT            float64 `json:"FIYAT"`
			ISLEMTURU        string  `json:"ISLEM_TURU"`
			BOLUM            string  `json:"BOLUM"`
			KURTAJORANI      float64 `json:"KURTAJ_ORANI"`
			SUBENO           int     `json:"SUBE_NO"`
			BROKER           int     `json:"BROKER"`
			IMKBISLEMNO      string  `json:"IMKB_ISLEM_NO"`
			EfektifAlis      float64 `json:"EfektifAlis"`
			TUTAR            float64 `json:"TUTAR"`
			MIKTAR           float64 `json:"MIKTAR"`
			KOMISYON         float64 `json:"KOMISYON"`
			BSMV             float64 `json:"BSMV"`
			ALISTUTAR        float64 `json:"ALIS_TUTAR"`
			SATISTUTAR       float64 `json:"SATIS_TUTAR"`
			ALISMIKTAR       float64 `json:"ALIS_MIKTAR"`
			SATISMIKTAR      float64 `json:"SATIS_MIKTAR"`
			NETMIKTAR        float64 `json:"NET_MIKTAR"`
			NETTUTAR         float64 `json:"NET_TUTAR"`
			SYSNO            int     `json:"SYS_NO"`
			MUSTERINO        int     `json:"MUSTERI_NO"`
			UNVANI           string  `json:"UNVANI"`
			IMKBSAAT         string  `json:"IMKB_SAAT"`
			KAYITSAAT        string  `json:"KAYIT_SAAT"`
			ALISKOMISYON     float64 `json:"ALIS_KOMISYON"`
			SATISKOMISYON    float64 `json:"SATIS_KOMISYON"`
			BORSAPAYIKESILEN float64 `json:"BORSA_PAYI_KESILEN"`
			BORSAPAYI        float64 `json:"BORSA_PAYI"`
			NET              float64 `json:"NET"`
			EmirVeren        int     `json:"EmirVeren"`
		} `json:"R1"`
		R2 []struct {
			ALISTLTOPLAM        float64 `json:"ALIS_TL_TOPLAM"`
			SATISTLTOPLAM       float64 `json:"SATIS_TL_TOPLAM"`
			ISLEMHACMI          float64 `json:"ISLEM_HACMI"`
			SATIMALIMFARKI      float64 `json:"SATIM_ALIM_FARKI"`
			KOMISYON            float64 `json:"KOMISYON"`
			BORSAPAYIKESILEN    float64 `json:"BORSA_PAYI_KESILEN"`
			BORSAPAYIHESAPLANAN float64 `json:"BORSA_PAYI_HESAPLANAN"`
			NETFARK             float64 `json:"NET_FARK"`
			BSMV                float64 `json:"BSMV"`
		} `json:"R2"`
		Output []struct {
			ReturnValue int `json:"ReturnValue"`
		} `json:"Output"`
	} `json:"data"`
	Success    bool   `json:"success"`
	Message    string `json:"message"`
	StatusCode int    `json:"statusCode"`
}

type menkulItem struct {
	MENKULKODU     string  `json:"MENKUL_KODU"`
	VALORTARIHI    string  `json:"VALOR_TARIHI"`
	ANA            int     `json:"ANA"`
	ACIKLAMA       string  `json:"ACIKLAMA"`
	CIKAN          float64 `json:"CIKAN"`
	GIREN          float64 `json:"GIREN"`
	KALAN          float64 `json:"KALAN"`
	FIS            string  `json:"FIS"`
	ISLEM          string  `json:"ISLEM"`
	FN             int     `json:"FN"`
	GUNUNTARIHI    string  `json:"GUNUN_TARIHI"`
	MENKULUNVANI   string  `json:"MENKUL_UNVANI"`
	KAYITANI       string  `json:"KAYIT_ANI"`
	SAAT           string  `json:"SAAT"`
	MUSTERINO      int     `json:"MUSTERI_NO"`
	UNVANI         string  `json:"UNVANI"`
	IDX            int     `json:"IDX"`
	MENKULGRUPKODU string  `json:"MENKUL_GRUP_KODU"`
}
