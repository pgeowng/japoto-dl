package onsen

import (
	"reflect"

	"github.com/mitchellh/mapstructure"
	"github.com/pgeowng/japoto-dl/model"
	"github.com/pkg/errors"
)

func (p *OnsenInfo) GetFeed() ([]model.Show, error) {
	mapObj := make([]map[string]interface{}, 0)
	err := p.loader.JSON("https://onsen.ag/web_api/programs/", &mapObj, nil)
	if err != nil {
		return nil, errors.Wrap(err, "onsen.feed.get")
	}

	resObj := []FeedRawShow{}
	conf := &mapstructure.DecoderConfig{
		ErrorUnused: true,
		Result:      &resObj,
	}

	mapstr, err := mapstructure.NewDecoder(conf)
	if err != nil {
		return nil, errors.Wrap(err, "onsen.feed.mapstr")
	}

	err = mapstr.Decode(mapObj)
	if err != nil {
		return nil, errors.Wrap(err, "onsen.feed.map")
	}

	// fmt.Printf("%#v", resObj)

	for i := range resObj {
		for j := range resObj[i].Contents {
			resObj[i].Contents[j].showRef = &resObj[i]
		}
	}

	result := make([]model.Show, 0)
	for i := range resObj {
		v := reflect.ValueOf(&resObj[i]).Interface()
		c := v.(model.Show)
		result = append(result, c)
	}

	// for _, show := range resObj {
	// 	for _, ep := range show.Contents {
	// 		fmt.Println(ep)
	// 	}
	// }

	return result, nil
}