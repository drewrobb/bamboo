package haproxy

import (
	conf "github.com/QubitProducts/bamboo/configuration"
	"github.com/QubitProducts/bamboo/services/marathon"
	"github.com/QubitProducts/bamboo/services/service"
)

type templateData struct {
	Apps     marathon.AppList
	Services map[string]service.Service
	HAProxy  conf.HAProxy
}

func GetTemplateData(config *conf.Configuration, storage service.Storage) (*templateData, error) {

	apps, err := marathon.FetchApps(config.Marathon, config)

	if err != nil {
		return nil, err
	}

	services, err := storage.All()
	if err != nil {
		return nil, err
	}

	return &templateData{apps, services, config.HAProxy}, nil
}
