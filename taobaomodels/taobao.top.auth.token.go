package taobaomodels

type TopAuthToken struct {
	W1ExpiresIn           int    `json:"w1_expires_in"`
	RefreshTokenValidTime int64  `json:"refresh_token_valid_time"`
	TaobaoUserNick        string `json:"taobao_user_nick"`
	ReExpiresIn           int    `json:"re_expires_in"`
	ExpireTime            int64  `json:"expire_time"`
	TokenType             string `json:"token_type"`
	SubTaobaoUserId       string `json:"sub_taobao_user_id"`
	AccessToken           string `json:"access_token"`
	TaobaoOpenUid         string `json:"taobao_open_uid"`
	W1Valid               int64  `json:"w1_valid"`
	RefreshToken          string `json:"refresh_token"`
	W2ExpiresIn           int    `json:"w2_expires_in"`
	SubTaobaoUserNick     string `json:"sub_taobao_user_nick"`
	W2Valid               int64  `json:"w2_valid"`
	R1ExpiresIn           int    `json:"r1_expires_in"`
	R2ExpiresIn           int    `json:"r2_expires_in"`
	R2Valid               int64  `json:"r2_valid"`
	R1Valid               int64  `json:"r1_valid"`
	TaobaoOpenSubUid      string `json:"taobao_open_sub_uid"`
	TaobaoUserId          string `json:"taobao_user_id"`
	ExpiresIn             int    `json:"expires_in"`
}
