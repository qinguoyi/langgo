package plugins

import (
	"errors"
	"github.com/olivere/elastic/v7"
	"go.uber.org/zap"
	"langgo/app/pkg/common"
	"langgo/bootstrap"
	"langgo/config"
	"sync"
)

var lgES = new(LangGoES)

// LangGoES .
type LangGoES struct {
	Once     *sync.Once
	ESClient *elastic.Client
	Index    string
}

// NewES .
func (lg *LangGoES) NewES() *elastic.Client {
	if lgES.ESClient != nil {
		return lg.ESClient
	} else {
		return lg.New().(*elastic.Client)
	}
}

func newLangGoES() *LangGoES {
	return &LangGoES{
		ESClient: &elastic.Client{},
		Once:     &sync.Once{},
	}
}

// Name .
func (lg *LangGoES) Name() string {
	return "ES"
}

// New .
func (lg *LangGoES) New() interface{} {
	lgES = newLangGoES()
	lgES.initES(bootstrap.NewConfig(""))
	return lg.ESClient
}

// Health .
func (lg *LangGoES) Health() {}

// Close .
func (lg *LangGoES) Close() {}

func init() {
	p := &LangGoES{}
	RegisteredPlugin(p)
}

func (lg *LangGoES) initES(conf *config.Configuration) {
	lg.Once.Do(func() {
		var err error
		if !common.IsAnyBlank(conf.ES.Url, conf.ES.Index) {
			lgES.Index = conf.ES.Index
			lgES.ESClient, err = elastic.NewClient(
				elastic.SetURL(conf.ES.Url),
				elastic.SetHealthcheck(false),
				elastic.SetSniff(false),
			)
		} else {
			err = errors.New("es config not found. ")
		}

		if err != nil {
			bootstrap.NewLogger().Logger.Error("es connect failed,", zap.Any("err", err.Error))
		}
	})
}
