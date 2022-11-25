package publickey

import (
	"bytes"
	"context"
	"io"
	"log"
	"net/http"
	"time"
)

const (
	updateDuration = 15 * time.Minute
)

type Service interface {
	Get() []byte
}

type service struct {
	Repository
}

func (s *service) init(ctx context.Context, pemURL string) {
	go func() {
		ticker := time.NewTicker(updateDuration)
		s.update(pemURL)
		for {
			select {
			case <-ticker.C:
				s.update(pemURL)
			case <-ctx.Done():
				return
			}
		}
	}()
}

func (s *service) update(pemURL string) {
	resp, err := http.Get(pemURL)
	if err != nil {
		log.Println(err)
		return
	}

	switch resp.StatusCode {
	case http.StatusOK:
		buf := bytes.NewBuffer([]byte{})
		_, err = io.Copy(buf, resp.Body)
		if err != nil {
			log.Println(err)
			return
		}
		s.Put(buf.Bytes())
	default:
		log.Printf("unrecognized status code: '%d'\n", resp.StatusCode)
	}
}

func NewService(ctx context.Context, pemURL string) Service {
	service := &service{}
	service.init(ctx, pemURL)

	return service
}
