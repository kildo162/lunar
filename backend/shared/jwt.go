package shared

import (
	"crypto/hmac"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"strings"
	"time"
)

type Data struct {
	SSID  string `json:"ssid"`
	Exp   int64  `json:"exp"`
	Nonce string `json:"nonce"`
	Role  string `json:"role"`
}

type JWT struct {
	SecretKey string
}

func NewJWT(secretKey string) *JWT {
	return &JWT{SecretKey: secretKey}
}

// [header.data.signature]
// expired: time duration
func (j *JWT) CreateToken(ssid string, expired time.Duration, role string) (string, Error) {

	expiredDuration := expired
	if expiredDuration == 0 || expiredDuration > 24*time.Hour || expiredDuration < 5*time.Minute {
		expiredDuration = 5 * time.Minute
	}

	expirationTime := time.Now().Add(expiredDuration).Unix()

	nonce, err := generateNonce()
	if err != nil {
		log.Println("failed to create token - nonce: " + err.Error())
		return "", SERVER_ERROR
	}

	claims := Data{
		SSID:  ssid,
		Exp:   expirationTime,
		Nonce: nonce,
		Role:  role,
	}

	header := map[string]interface{}{
		"alg": "HS256",
		"typ": "JWT",
	}

	headerBytes, err := json.Marshal(header)
	if err != nil {
		log.Println("failed to create token - header: " + err.Error())
		return "", SERVER_ERROR
	}

	encodedHeader := base64.URLEncoding.EncodeToString(headerBytes)

	bodyBytes, err := json.Marshal(claims)
	if err != nil {
		log.Println("failed to create token - verify: " + err.Error())
		return "", SERVER_ERROR
	}

	encodedData := base64.URLEncoding.EncodeToString(bodyBytes)

	signatureInput := fmt.Sprintf("%s.%s", encodedHeader, encodedData)
	signature := j.signHMACSHA256(signatureInput)

	token := fmt.Sprintf("%s.%s.%s", encodedHeader, encodedData, signature)

	return token, nil
}

func (j *JWT) VerifyToken(token string) (Data, Error) {
	parts := strings.Split(token, ".")
	if len(parts) != 3 {
		return Data{}, PERMISSION_ERROR
	}

	encodedHeader := parts[0]
	encodedData := parts[1]
	signature := parts[2]

	signatureInput := fmt.Sprintf("%s.%s", encodedHeader, encodedData)
	if !j.verifyHMACSHA256(signatureInput, signature) {
		log.Println("invalid signature")
		return Data{}, PERMISSION_ERROR
	}

	headerBytes, err := base64.URLEncoding.DecodeString(encodedHeader)
	if err != nil {
		log.Println("failed to decode token - header: " + err.Error())
		return Data{}, PERMISSION_ERROR
	}

	claimsBytes, err := base64.URLEncoding.DecodeString(encodedData)
	if err != nil {
		log.Println("failed to decode token - header: " + err.Error())
		return Data{}, PERMISSION_ERROR
	}

	var header map[string]interface{}
	err = json.Unmarshal(headerBytes, &header)
	if err != nil {
		log.Println("failed to decode token - header: " + err.Error())
		return Data{}, PERMISSION_ERROR
	}

	if alg, ok := header["alg"].(string); !ok || alg != "HS256" {
		log.Println("unsupported signing algorithm")
		return Data{}, PERMISSION_ERROR
	}

	var claims Data
	err = json.Unmarshal(claimsBytes, &claims)
	if err != nil {
		log.Println("failed to verify token: " + err.Error())
		return Data{}, PERMISSION_ERROR
	}

	if time.Now().Unix() > claims.Exp {
		return Data{}, PERMISSION_ERROR
	}

	return claims, nil
}

func (j *JWT) signHMACSHA256(data string) string {
	keyBytes := []byte(j.SecretKey)

	mac := hmac.New(sha256.New, keyBytes)
	mac.Write([]byte(data))
	signature := mac.Sum(nil)

	return base64.URLEncoding.EncodeToString(signature)
}

func (j *JWT) verifyHMACSHA256(data string, signature string) bool {
	expectedSignature := j.signHMACSHA256(data)
	return signature == expectedSignature
}

// generateNonce generates a random nonce string for JWT token.
// The nonce is used to prevent replay attacks.
func generateNonce() (string, error) {
	nonceBytes := make([]byte, 16)
	_, err := rand.Read(nonceBytes)
	if err != nil {
		return "", err
	}
	return base64.URLEncoding.EncodeToString(nonceBytes), nil
}
