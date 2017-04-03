package main

//go:generate go-bindata -pkg $GOPACKAGE -o assets.go assets/

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"image/gif"
	"image/png"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
)

type slackResponse struct {
	OK    bool   `json:"ok"`
	Error string `json:"error"`
}

func setPhoto(token string, image []byte) error {
	var body bytes.Buffer
	writer := multipart.NewWriter(&body)
	defer writer.Close()
	_ = writer.WriteField("token", token)
	part, err := writer.CreateFormFile("image", "parrot.png")
	if err != nil {
		return err
	}
	_, err = io.Copy(part, bytes.NewBuffer(image))
	resp, err := http.Post("https://slack.com/api/users.setPhoto", writer.FormDataContentType(), &body)
	if err != nil {
		return err
	}
	var slackResp slackResponse
	dec := json.NewDecoder(resp.Body)
	dec.Decode(&slackResp)
	if !slackResp.OK {
		return fmt.Errorf("slack: %s", slackResp.Error)
	}
	return nil
}

func main() {
	var slackToken, parrot string
	flag.StringVar(&slackToken, "slack-token", "",
		"A Slack API token, generatable at https://api.slack.com/custom-integrations/legacy-tokens")
	flag.StringVar(&parrot, "parrot", "",
		"Path to a parrot gif to use. If not provided, will default to the original party parrot.")
	flag.Parse()
	if slackToken == "" {
		fmt.Fprintf(os.Stderr, "flag is required: -slack-token\n")
		flag.Usage()
		os.Exit(2)
	}

	fmt.Println("profile parrot is running!")

	var reader io.Reader
	if parrot == "" {
		parrotBytes := MustAsset("assets/parrot.gif")
		reader = bytes.NewReader(parrotBytes)
	} else {
		f, err := os.Open(parrot)
		if err != nil {
			fmt.Fprintf(os.Stderr, "error: %s\n", err)
			os.Exit(1)
		}
		reader = f
	}
	parrotGIF, err := gif.DecodeAll(reader)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %s\n", err)
		os.Exit(1)
	}

	// convert images to pngs
	convertedImages := make([][]byte, len(parrotGIF.Image))
	for i, img := range parrotGIF.Image {
		buf := &bytes.Buffer{}
		if err := png.Encode(buf, img); err != nil {
			fmt.Fprintf(os.Stderr, "error: %s\n", err)
			os.Exit(1)
		}
		b, err := ioutil.ReadAll(buf)
		if err != nil {
			fmt.Fprintf(os.Stderr, "error: %s\n", err)
			os.Exit(1)
		}
		convertedImages[i] = b
	}

	i := 0
	for {
		err := setPhoto(slackToken, convertedImages[i])
		if err != nil {
			fmt.Fprintf(os.Stderr, "error: %s\n", err)
			os.Exit(1)
		}
		i = (i + 1) % len(parrotGIF.Image)
	}
}
