package main

import (
	"cloud_asr"
	"encoding/base64"
	"golang.org/x/net/context"
	"io/ioutil"
)

func main() {
	credentials := cloud_asr.AuthRequestDto{"user", 261, "password"}
	client := cloud_asr.NewAPIClient(cloud_asr.NewConfiguration())
	sessionApi := client.SessionApi
	loginResponse, _, _ := sessionApi.Login(context.Background(), credentials)
	sessionId := loginResponse.SessionId
	packageApi := client.PackagesApi
	packageApi.Load(context.Background(), sessionId, "CommonRus")
	data, _ := ioutil.ReadFile("test.wav")
	sEnc := base64.StdEncoding.EncodeToString([]byte(data))
	audio_for_recognition := &cloud_asr.AudioFileDto{sEnc, "audio/wav"}
	recognition_request := cloud_asr.RecognitionRequestDto{audio_for_recognition, "CommonRus"}
	recognitionApi := client.RecognizeApi
	recognition_result, _, _ := recognitionApi.Recognize(context.Background(), sessionId, recognition_request, nil)

}