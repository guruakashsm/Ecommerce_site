package service

import (
	"bytes"
	"crypto/hmac"
	"crypto/rand"
	"crypto/sha1"
	"encoding/base32"
	"encoding/binary"
	"log"
	"math"
	"math/big"
	"strconv"
	"testing"
	"time"
)

const (
	RETURN_DIGITS int = 6
	SECRET_LENGTH int = 16
)

func ValidateOTP(otp, secret string) bool {
	correctOTP, err := GetOTP(secret)
	log.Println(correctOTP)
	return (otp == correctOTP) && (err == nil)
}


func GetOTP(secret string) (string, error) {
	return computeTOTP(secret, timestamp())
}


func GenerateSecret() (string, error) {
	return generateBase32CryptoString(SECRET_LENGTH)
}


func TestTOTP(t *testing.T) {
	secret := "ABCDEFGHIJKLMNOPQRSTUVWXYZ234567"
	timeIntervals := int64(47156746)
	correctOTP := "109431"
	otp, err := computeTOTP(secret, timeIntervals)
	if err != nil {
		t.Errorf("computeTOTP returns an error for valid input: %s", err.Error())
	}
	if otp != correctOTP {
		t.Errorf("TOTP(%s, %d) = %s, what %s", secret, timeIntervals, otp, correctOTP)
	}
}

func computeTOTP(secret string, time int64) (string, error) {
	key, err := decodeSecret(secret)
	if err != nil {
		return "", err
	}

	msg := encodeTime(time)
	hash := computeHMAC(msg, key)

	offset := hash[len(hash)-1] & 0x0F
	binary := (int(hash[offset]&0x7F) << 24) |
		(int(hash[offset+1]&0xFF) << 16) |
		(int(hash[offset+2]&0xFF) << 8) |
		int(hash[offset+3]&0xFF)
	otp := binary % int(math.Pow10(RETURN_DIGITS))

	result := strconv.Itoa(otp)
	for len(result) < RETURN_DIGITS {
		result = "0" + result
	}

	return result, nil
}



func timestamp() int64 {
	return time.Now().Unix() / 30
}

func computeHMAC(data, secret []byte) []byte {
	mac := hmac.New(sha1.New, secret)
	mac.Write(data)
	return mac.Sum(nil)
}

func decodeSecret(secret string) ([]byte, error) {
	return base32.StdEncoding.DecodeString(secret)
}



func encodeTime(time int64) []byte {
	buf := new(bytes.Buffer)
	binary.Write(buf, binary.BigEndian, time)
	return buf.Bytes()
}

func generateBase32CryptoString(length int) (string, error) {
	str := "ABCDEFGHIJKLMNOPQRSTUVWXYZ234567"
	buf := &bytes.Buffer{}
	for i := 0; i < length; i++ {
		choice, err := rand.Int(rand.Reader, big.NewInt(int64(len(str))))
		if err != nil {
			return "", err
		}
		buf.WriteString(string(str[uint8(choice.Int64())]))
	}
	return buf.String(), nil
}