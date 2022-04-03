package kls3

import (
	"errors"
	"testing"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/golang/mock/gomock"
	"github.com/lalamove/konfig"
	"github.com/lalamove/konfig/mocks"
	"github.com/lalamove/konfig/watcher/kwpoll"
	"github.com/stretchr/testify/require"
)


type BufferMatcher struct{
	buffer *aws.WriteAtBuffer
	msg string
}

func (bm *BufferMatcher) Matches(b *aws.WriteAtBuffer) bool{
	return true
}

func (bm *BufferMatcher) String() string{
	return bm.msg
}

type GetObjectInputRequestMatcher struct{
	input *s3.GetObjectInput
	msg string
}

func (g *GetObjectInputRequestMatcher) String() string{
	return g.msg
}

func (g *GetObjectInputRequestMatcher) Matches(x interface{}) bool {

	if v, ok := x.(*s3.GetObjectInput); ok {
		if aws.StringValue(v.Bucket) != aws.StringValue(g.input.Bucket) {
			g.msg = "bucket names are different"
			return false
		}

		if aws.StringValue(v.Key) != aws.StringValue(g.input.Key) {
			g.msg = "object key is different"
			return false
		}

		return true
	}
	return false
}

func TestLoad(t *testing.T) {
	var testCases = []struct {
		name  string
		setUp func(ctrl *gomock.Controller) *Loader
		err   bool
	}{
		{
			name: "single source no error download object",
			setUp: func(ctrl *gomock.Controller) *Loader {
				var c = mocks.NewMockDownloader(ctrl)
				var p1 = mocks.NewMockParser(ctrl)

				var hl = New(&Config{
					Downloader: c,
					Objects: []Object{
						{
							Bucket:    "com.lalamove.konfig",
							Key: "config.json",
							Parser: p1,
						},
					},
				})

				var input = &s3.GetObjectInput{
					Bucket: aws.String("com.lalamove.konfig"),
					Key: aws.String("config.json"),
				}
				c.EXPECT().Download(gomock.Any(),&GetObjectInputRequestMatcher{input:input}).Times(1).Return(
					int64(0),
					nil,
				)

				p1.EXPECT().Parse(gomock.Any(), konfig.Values{}).Times(1).Return(nil)

				return hl
			},
			err: false,
		},
		{
			name: "multiple sources no error download object",
			setUp: func(ctrl *gomock.Controller) *Loader {
				var c = mocks.NewMockDownloader(ctrl)
				var p1 = mocks.NewMockParser(ctrl)
				var p2 = mocks.NewMockParser(ctrl)

				var hl = New(&Config{
					Downloader: c,
					Objects: []Object{
						{
							Bucket:    "com.lalamove.konfig",
							Key: "config1.json",
							Parser: p1,
						},
						{
							Bucket:    "com.lalamove.konfig",
							Key: "config2.toml",
							Parser: p2,
						},
					},
				})

				var req1 = &s3.GetObjectInput{Bucket: aws.String("com.lalamove.konfig"), Key: aws.String("config1.json")}
				var req2 = &s3.GetObjectInput{Bucket: aws.String("com.lalamove.konfig"), Key: aws.String("config2.toml")}

				gomock.InOrder(
					c.EXPECT().Download(gomock.Any(), &GetObjectInputRequestMatcher{input: req1}).Times(1).Return(
						int64(0),
						nil,
					),
					c.EXPECT().Download(gomock.Any(), &GetObjectInputRequestMatcher{input: req2}).Times(1).Return(
						int64(0),
						nil,
					),
				)

				p1.EXPECT().Parse(gomock.Any(), konfig.Values{}).Times(1).Return(nil)
				p2.EXPECT().Parse(gomock.Any(), konfig.Values{}).Times(1).Return(nil)

				return hl
			},
			err: false,
		},
		{
			name: "multiple sources watch no error download object",
			setUp: func(ctrl *gomock.Controller) *Loader {
				var c = mocks.NewMockDownloader(ctrl)
				var p1 = mocks.NewMockParser(ctrl)
				var p2 = mocks.NewMockParser(ctrl)

				var req1 = &s3.GetObjectInput{Bucket: aws.String("com.lalamove.konfig"), Key: aws.String("config1.json")}
				var req2 = &s3.GetObjectInput{Bucket: aws.String("com.lalamove.konfig"), Key: aws.String("config2.toml")}

				gomock.InOrder(
					c.EXPECT().Download(gomock.Any(), &GetObjectInputRequestMatcher{input: req1}).Times(1).Return(
						int64(0),
						nil,
					),
					c.EXPECT().Download(gomock.Any(), &GetObjectInputRequestMatcher{input: req2}).Times(1).Return(
						int64(0),
						nil,
					),
				)

				p1.EXPECT().Parse(gomock.Any(), konfig.Values{}).Times(1).Return(nil)
				p2.EXPECT().Parse(gomock.Any(), konfig.Values{}).Times(1).Return(nil)

				var hl = New(&Config{
					Downloader: c,
					Watch:  true,
					Rater:  kwpoll.Time(100 * time.Millisecond),
					Objects: []Object{
						{
							Bucket:    "com.lalamove.konfig",
							Key: "config1.json",
							Parser: p1,
						},
						{
							Bucket:    "com.lalamove.konfig",
							Key: "config2.toml",
							Parser: p2,
						},
					},
				})

				gomock.InOrder(
					c.EXPECT().Download(gomock.Any(), &GetObjectInputRequestMatcher{input: req1}).Times(1).Return(
						int64(0),
						nil,
					),
					c.EXPECT().Download(gomock.Any(), &GetObjectInputRequestMatcher{input: req2}).Times(1).Return(
						int64(0),
						nil,
					),
				)

				p1.EXPECT().Parse(gomock.Any(), konfig.Values{}).Times(1).Return(nil)
				p2.EXPECT().Parse(gomock.Any(), konfig.Values{}).Times(1).Return(nil)

				return hl
			},
			err: false,
		},
		{
			name: "multiple sources no error download object",
			setUp: func(ctrl *gomock.Controller) *Loader {
				var c = mocks.NewMockDownloader(ctrl)
				var p1 = mocks.NewMockParser(ctrl)
				var p2 = mocks.NewMockParser(ctrl)

				var hl = New(&Config{
					Downloader: c,
					Objects: []Object{
						{
							Bucket:    "com.lalamove.konfig",
							Key: "config1.json",
							Parser: p1,
						},
						{
							Bucket:    "com.lalamove.konfig2",
							Key: "config2.toml",
							Parser: p2,
						},
					},
				})

				var req1 = &s3.GetObjectInput{Bucket: aws.String("com.lalamove.konfig"), Key: aws.String("config1.json")}
				var req2 = &s3.GetObjectInput{Bucket: aws.String("com.lalamove.konfig2"), Key: aws.String("config2.toml")}

				gomock.InOrder(
					c.EXPECT().Download(gomock.Any(), &GetObjectInputRequestMatcher{input: req1}).Times(1).Return(
						int64(0),
						nil,
					),
					c.EXPECT().Download(gomock.Any(), &GetObjectInputRequestMatcher{input: req2}).Times(1).Return(
						int64(0),
						errors.New(""),
					),
				)

				p1.EXPECT().Parse(gomock.Any(), konfig.Values{}).Times(1).Return(nil)

				return hl
			},
			err: true,
		},
		{
			name: "single object error",
			setUp: func(ctrl *gomock.Controller) *Loader {
				var c = mocks.NewMockDownloader(ctrl)
				var p1 = mocks.NewMockParser(ctrl)

				var hl = New(&Config{
					Downloader: c,
					Objects: []Object{
						{
							Bucket:        "com.lalamove.konfig",
							Parser:     p1,
							Key: "config.json",
						},
					},
				})

				var req = &s3.GetObjectInput{Bucket: aws.String("com.lalamove.konfig"), Key: aws.String("config.json")}
				c.EXPECT().Download(gomock.Any(), &GetObjectInputRequestMatcher{input: req}).Times(1).Return(
					int64(0),
					errors.New("Oops"),
				)
				return hl
			},
			err: true,
		},
	}

	for _, testCase := range testCases {
		t.Run(
			testCase.name,
			func(t *testing.T) {
				var ctrl = gomock.NewController(t)
				defer ctrl.Finish()

				konfig.Init(konfig.DefaultConfig())
				var hl = testCase.setUp(ctrl)

				var err = hl.Load(konfig.Values{})
				if testCase.err {
					require.NotNil(t, err, "err should not be nil")
					return
				}
				require.Nil(t, err, "err should be nil")
			},
		)
	}
}

func TestNew(t *testing.T) {
	t.Run(
		"default http client",
		func(t *testing.T) {
			var ctrl = gomock.NewController(t)
			defer ctrl.Finish()

			var p = mocks.NewMockParser(ctrl)

			var hl = New(&Config{
				Objects: []Object{
					{
						Bucket:    "com.lalamove.konfig",
						Key: "config.json",
						Parser: p,
					},
				},
			})

			require.Equal(t, defaultDownloader, hl.cfg.Downloader)
		},
	)
	t.Run(
		"panic no sources",
		func(t *testing.T) {
			var ctrl = gomock.NewController(t)
			defer ctrl.Finish()

			require.Panics(t, func() {
				New(&Config{
					Objects: []Object{},
				})
			})
		},
	)
}

func TestLoaderMethods(t *testing.T) {

	var ctrl = gomock.NewController(t)
	defer ctrl.Finish()

	var p = mocks.NewMockParser(ctrl)

	var hl = New(&Config{
		Name:          "s3",
		MaxRetry:      1,
		RetryDelay:    1 * time.Second,
		StopOnFailure: true,
		Objects: []Object{
			{
				Bucket:    "com.lalamove.konfig",
				Key: "config.json",
				Parser: p,
			},
		},
	})

	require.True(t, hl.StopOnFailure())
	require.Equal(t, "s3", hl.Name())
	require.Equal(t, 1*time.Second, hl.RetryDelay())
	require.Equal(t, 1, hl.MaxRetry())
}
