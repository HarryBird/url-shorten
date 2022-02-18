package data

import (
	"context"
	"url-shorten/app/shorten/internal/biz"
	"url-shorten/app/shorten/internal/data/dao/model"
	"url-shorten/app/shorten/internal/data/dao/query"

	"github.com/HarryBird/mo-kit/msgr"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

var _ biz.ShortenRepo = (*shortenRepo)(nil)

type shortenRepo struct {
	data *Data
	log  *log.Helper
}

func NewShortenRepo(data *Data, logger log.Logger) biz.ShortenRepo {
	return &shortenRepo{
		data: data,
		log:  log.NewHelper(log.With(logger, "mod", "repo.shorten")),
	}
}

func (r *shortenRepo) Create(ctx context.Context, url *biz.ShortenURL) (*biz.ShortenURL, error) {
	fname := "Create"
	dao := query.Use(r.data.db).URLShortened

	r.log.WithContext(ctx).Infof("%s check shorten url existed: urlcode=%v", msgr.W(fname), url.URLCode)
	existedShortenURL, err := dao.WithContext(ctx).Where(dao.URLCode.Eq(url.URLCode)).First()

	// 短链已存在
	if err == nil {
		r.log.WithContext(ctx).Infof("%s shorten url had existed: %+v", msgr.W(fname), existedShortenURL)
		if existedShortenURL.URLFull == url.URLFull {
			return r.modelToBiz(ctx, existedShortenURL), nil
		}

		// 短链存在，但原始链接不同，hash 冲突
		return nil, errors.New("repo: shorten url hash conflict")
	}

	// 短链不存在，创建
	if errors.Is(err, gorm.ErrRecordNotFound) {
		newShortenURL, err := r.create(ctx, url)
		if err != nil {
			return nil, err
		}

		return r.modelToBiz(ctx, newShortenURL), nil
	}

	return nil, errors.WithMessage(err, "repo: query shorten url existed fail")
}

func (r *shortenRepo) create(ctx context.Context, url *biz.ShortenURL) (*model.URLShortened, error) {
	fname := "create"
	dao := query.Use(r.data.db).URLShortened

	// biz to model
	u := r.bizToModel(ctx, url)

	// store
	r.log.WithContext(ctx).Infof("%s prepare data: %+v", msgr.W(fname), u)
	if err := dao.WithContext(ctx).Create(u); err != nil {
		return nil, errors.Wrap(err, "repo: create shorten url fail")
	}
	r.log.WithContext(ctx).Infof("%s after insert db: %+v", msgr.W(fname), u)

	return u, nil
}

func (r *shortenRepo) bizToModel(ctx context.Context, url *biz.ShortenURL) *model.URLShortened {
	return &model.URLShortened{
		URLFull:  url.URLFull,
		URLHost:  url.URLHost,
		URLURI:   url.URLUri,
		URLQuery: url.URLQuery,
		URLCode:  url.URLCode,
	}
}

func (r *shortenRepo) modelToBiz(ctx context.Context, url *model.URLShortened) *biz.ShortenURL {
	return &biz.ShortenURL{
		URLFull:  url.URLFull,
		URLHost:  url.URLHost,
		URLUri:   url.URLURI,
		URLQuery: url.URLQuery,
		URLCode:  url.URLCode,
	}
}
