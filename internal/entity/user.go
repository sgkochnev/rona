package entity

type User struct {
	Username     string `json:"username" bson:"_id"`
	Email        string `json:"email" bson:"email"`
	Password     string `json:"password" bson:"password"`
	Fingerprint  string `json:"fingerprint" bson:"fingerprint"`
	RefreshToken `json:"refresh_token" bson:"refresh_token"`
	//RefreshToken string `json:"refresh_token" bson:"refresh_token"`
}

//type Fingerprint struct {
//	Value    string
//	Username string
//}
