package main

import (
	"cloud_diarization"
	"encoding/base64"
	"fmt"
	"golang.org/x/net/context"
	"io/ioutil"
)

func main() {
	credentials := cloud_diarization.AuthRequestDto{"user", 261, "password"}
	client := cloud_diarization.NewAPIClient(cloud_diarization.NewConfiguration())
	sessionApi := client.SessionApi
	loginResponse, _, _ := sessionApi.Login(context.Background(), credentials)
	sessionId := loginResponse.SessionId
	audioBytes, _ := ioutil.ReadFile("test.wav")
	base64Bytes := base64.StdEncoding.EncodeToString([]byte(audioBytes))
	audio := cloud_diarization.AudioDto{base64Bytes}
	diarizationApi := client.DiarizationApi
	diarizationResponse, _, _ := diarizationApi.Diarization(context.Background(), sessionId, audio)
	speakers := diarizationResponse.Speakers
}
