package service

import (
	"context"
	"encoding/json"
	"log"
	"net/url"
	"time"

	"proxy/internal/metrics"
	"proxy/internal/model"

	"github.com/ekomobile/dadata/v2/api/suggest"
	"github.com/ekomobile/dadata/v2/client"
	"github.com/go-redis/redis"
)

type Casher interface {
	Set(key string, value interface{}) error
	Get(key string) (string, error)
}

type Service struct {
	api       *suggest.Api
	apiKey    string
	secretKey string
	Casher    Casher
}

func (s *Service) Cash(input string) ([]*model.Address, error) {
	var res []*model.Address
	jsoninput, err := json.Marshal(input)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	start := time.Now()
	body, err := s.Casher.Get(string(jsoninput))
	metrics.CacheDuration.Observe(time.Since(start).Seconds())

	if err == redis.Nil {
		log.Println("from API")
		start = time.Now()
		bodys, err := s.AddressSearch(input)
		metrics.ApiDuration.Observe(time.Since(start).Seconds())
		if err != nil {
			log.Println(err)
			return nil, err
		}
		jsonbody, err := json.Marshal(bodys)
		if err != nil {
			log.Println(err)
			return nil, err
		}

		s.Casher.Set(string(jsoninput), jsonbody)
		return bodys, nil
	} else if err != nil {
		log.Println(err)
		return nil, err
	}
	log.Println("from cache")
	err = json.Unmarshal([]byte(body), &res)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return res, nil
}

func (s *Service) AddressSearch(input string) ([]*model.Address, error) {
	var res []*model.Address
	rawRes, err := s.api.Address(context.Background(), &suggest.RequestParams{Query: input})
	if err != nil {
		return nil, err
	}
	for _, r := range rawRes {
		if r.Data.City == "" || r.Data.Street == "" {
			continue
		}
		res = append(res, &model.Address{City: r.Data.City, Street: r.Data.Street, House: r.Data.House, Lat: r.Data.GeoLat, Lon: r.Data.GeoLon})
	}
	return res, nil
}

func New(c Casher) Service {
	endpointUrl, _ := url.Parse("https://suggestions.dadata.ru/suggestions/api/4_1/rs/")
	creds := client.Credentials{
		ApiKeyValue:    "a24f884fd9737d9d7604e288667bc4719b6bd9b6",
		SecretKeyValue: "a455650abcfca33d33239b6a286784586a6a0b71",
	}
	api := suggest.Api{
		Client: client.NewClient(endpointUrl, client.WithCredentialProvider(&creds)),
	}

	service := Service{
		api:       &api,
		apiKey:    "a24f884fd9737d9d7604e288667bc4719b6bd9b6",
		secretKey: "a455650abcfca33d33239b6a286784586a6a0b71",
		Casher:    c,
	}
	return service
}
