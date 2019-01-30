package tts

import (
	"cloud_tts"
	"encoding/base64"
	"golang.org/x/net/context"
	"io/ioutil"
)

func main() {
	credentials := cloud_tts.AuthRequestDto{"user", 261, "password"}
	client := cloud_tts.NewAPIClient(cloud_tts.NewConfiguration())
	sessionApi := client.SessionApi
	loginResponse, _, _ := sessionApi.Login(context.Background(), credentials)
	sessionId := loginResponse.SessionId
	synthesis := client.SynthesizeApi
	synthesisText := &cloud_tts.SynthesizeText{"text/plain", "Text to be synthesised"}
	synthesisRequest := cloud_tts.SynthesizeRequest{synthesisText, "Carol", "audio/wav"}
	synthesisResponse, _, _ := synthesis.Synthesize(context.Background(), sessionId, synthesisRequest, nil)
	sound, _ := base64.StdEncoding.DecodeString(synthesisResponse.Data)
	ioutil.WriteFile("test.wav", sound, 0644)
}
