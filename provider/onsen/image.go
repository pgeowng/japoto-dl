package onsen

import (
	"fmt"

	"github.com/pgeowng/japoto-dl/model"
	"github.com/pgeowng/japoto-dl/provider/common"
	"github.com/pkg/errors"
)

func (show *OnsenShow) LoadImage(loader model.Loader, workdir model.WorkdirBase) error {

	url := show.ProgramInfo.Image.Url
	if len(url) == 0 {
		return errors.New("empty image url")
	}

	imageBody, err := loader.Raw(url, gopts)
	if err != nil {
		return errors.Wrap(err, "onsen.img")
	}

	ext := common.GuessContentType(imageBody)
	filename := fmt.Sprintf("%s--%s%s", show.Provider(), show.ShowId(), ext)

	return errors.Wrap(workdir.SaveRaw(filename, imageBody), "onsen.img")
}
