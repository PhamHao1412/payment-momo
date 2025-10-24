package pkg

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

func HmacSHA256(raw, key string) string {
	h := hmac.New(sha256.New, []byte(key))
	h.Write([]byte(raw))
	return hex.EncodeToString(h.Sum(nil))
}

func HttpPostJSON[Resp any](url string, body any) (*Resp, error) {
	b, _ := json.Marshal(body)
	req, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(b))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{Timeout: 15 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	data, _ := io.ReadAll(resp.Body)
	if resp.StatusCode >= 300 {
		return nil, fmt.Errorf("http %d: %s", resp.StatusCode, string(data))
	}
	var out Resp
	if err := json.Unmarshal(data, &out); err != nil {
		return nil, fmt.Errorf("decode error: %w; body=%s", err, string(data))
	}
	return &out, nil
}
