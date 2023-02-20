package auth

import (
	"github.com/dgrijalva/jwt-go"
	log "github.com/sirupsen/logrus"
	"tiktok/cmd/api/db"
)

var Secret = "tiktok"

var ExpiresTime = 60 * 60 * 24

func Auth(token string) bool {
	if len(token) == 0 {
		log.Error("not have token")
		return false
	}
	auth, err := parseToken(token)
	if err != nil {
		log.Error("token parse error")
		return false
	}	
	log.Info("token parse success")
	return true
}


func parseToken(token string) (*jwt.StandardClaims, error) {
	jwtToken, err := jwt.ParseWithClaims(token, &jwt.StandardClaims{}, func(token *jwt.Token) (i interface{}, e error) {
		return []byte(Secret), nil
	})
	if err == nil && jwtToken != nil {
		if claim, ok := jwtToken.Claims.(*jwt.StandardClaims); ok && jwtToken.Valid {
			return claim, nil
		}
	}
	return nil, err
}

func GenerateToken(username string) string {
	user, err := db.GetUserByUsername(username)
	if err != nil {
		log.Error("not find user")
	}
	token := CreateToken(user)
	log.Info("token: ", token)
	return token
}

func CreateToken(user db.User) string {
	expiresTime := time.Now().Unix() + int64(ExpiresTime)
	id := user.Uid
	claims := jwt.StandardClaims{
		Audience:  user.Name,
		ExpiresAt: expiresTime,
		Id:        strconv.FormatInt(id64, 10),
		IssuedAt:  time.Now().Unix(),
		Issuer:    "tiktok",
		NotBefore: time.Now().Unix(),
		Subject:   "token",
	}
	var jwtSecret = []byte(Secret)
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	if token, err := tokenClaims.SignedString(jwtSecret); err == nil {
		log.Info("generate token success!\n")
		return token
	} else {
		log.Error("generate token fail\n")
		return "fail"
	}
}
