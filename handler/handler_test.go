package handler

import (
	"bytes"
	"context"
	"database/sql"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/labstack/echo/v4"

	"github.com/stretchr/testify/assert"

	"github.com/volatiletech/null"

	"github.com/c8112002/twitter_clone_go/models"
	"github.com/c8112002/twitter_clone_go/router"
	"github.com/c8112002/twitter_clone_go/store"
	"github.com/c8112002/twitter_clone_go/utils"
	"github.com/volatiletech/sqlboiler/boil"

	"github.com/c8112002/twitter_clone_go/db"
	_ "github.com/go-sql-driver/mysql"
)

var (
	h      *Handler
	d      *sql.DB
	us     *store.UserStore
	ts     *store.TweetStore
	ctx    context.Context
	update = flag.Bool("update", false, "update .golden files")
)

func setup() {
	var err error
	d, err = db.TestDB(false, utils.Location())

	if err != nil {
		panic(err.Error())
	}

	_ = router.New()

	ctx = context.Background()
	us = store.NewUserStore(d, ctx)
	ts = store.NewTweetStore(d, ctx)

	h = NewHandler(us, ts)

	if err := db.MigrateTestDB(d); err != nil {
		panic(err.Error())
	}

	if err := loadFixtures(); err != nil {
		panic(err.Error())
	}
}

func tearDown() {
	if err := db.DropTestDB(d); err != nil {
		panic(err.Error())
	}
}

func loadFixtures() error {

	if err := loadUsers(); err != nil {
		return err
	}

	if err := loadTweets(); err != nil {
		return err
	}

	return nil
}

func loadUsers() error {
	tx, err := d.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	users := []models.User{
		models.User{
			Name:    "鈴木 一郎",
			Icon:    "https://icon/1",
			Profile: "こんにちは。鈴木一郎です。",
		},
		models.User{
			Name:    "佐藤 二郎",
			Icon:    "https://icon/2",
			Profile: "こんにちは。佐藤二郎です。",
		},
		models.User{
			Name:    "田中 三郎",
			Icon:    "https://icon/3",
			Profile: "こんにちは。田中三郎です。",
		},
		models.User{
			Name:    "高橋 四郎",
			Icon:    "https://icon/4",
			Profile: "こんにちは。高橋四郎です。",
		},
		models.User{
			Name:      "橋本 五郎",
			Icon:      "https://icon/5",
			Profile:   "こんにちは。橋本五郎です。",
			DeletedAt: null.NewTime(time.Now(), true),
		},
	}

	for _, u := range users {
		if err := u.Insert(ctx, tx, boil.Infer()); err != nil {
			return err
		}
	}

	if err := tx.Commit(); err != nil {
		return err
	}

	return nil
}

func loadTweets() error {
	tx, err := d.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	nt := 40
	nu := 5

	for i := 0; i < nt; i++ {
		tid := i + 1

		var deletedAt null.Time
		if tid <= 30 {
			deletedAt = null.NewTime(time.Now(), false)
		} else {
			deletedAt = null.NewTime(time.Now(), true)
		}

		t := models.Tweet{
			UserID:    uint(i%nu + 1),
			Tweet:     fmt.Sprintf("ツイート %d", tid),
			DeletedAt: deletedAt,
		}

		if err := t.Insert(ctx, tx, boil.Infer()); err != nil {
			return err
		}
	}

	if err := tx.Commit(); err != nil {
		return err
	}

	return nil
}

func newEchoContext(r *http.Request, w http.ResponseWriter) echo.Context {
	e := echo.New()
	return e.NewContext(r, w)
}

func newRequest(method, target string, body io.Reader) *http.Request {
	req := httptest.NewRequest(method, target, body)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	return req
}

func newContext(r *http.Request, w http.ResponseWriter) echo.Context {
	e := echo.New()
	return e.NewContext(r, w)
}

func assertResponse(t *testing.T, res *http.Response, code int, goldenFilePath string) {
	t.Helper()

	assertResponseHeader(t, res, code)
	assertResponseBody(t, res, goldenFilePath)
}

func assertResponseHeader(t *testing.T, res *http.Response, code int) {
	t.Helper()

	if code != res.StatusCode {
		t.Errorf("expected status code is '%d',\n but actual given code is '%d'", code, res.StatusCode)
	}

	if expected := "application/json; charset=UTF-8"; res.Header.Get("Content-Type") != expected {
		t.Errorf("unexpected response Content-Type,\n expected: %#v,\n but given #%v", expected, res.Header.Get("Content-Type"))
	}
}

func assertResponseBody(t *testing.T, res *http.Response, goldenFilePath string) {
	t.Helper()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Fatalf("unexpected error by ioutil.ReadAll() '%#v'", err)
	}

	var actual bytes.Buffer
	err = json.Indent(&actual, body, "", "  ")
	if err != nil {
		t.Fatalf("unexpected error by json.Indent '%#v'", err)
	}

	if *update {
		updateGoldenFile(t, actual, goldenFilePath)
	}

	rs := getStringFromTestFile(t, goldenFilePath)

	assert.JSONEq(t, rs, actual.String())
}

func getStringFromTestFile(t *testing.T, path string) string {
	t.Helper()

	bt, err := ioutil.ReadFile(path)
	if err != nil {
		t.Fatalf("unexpected error while opening file '%#v'", err)
	}
	return string(bt)
}

func updateGoldenFile(t *testing.T, actual bytes.Buffer, path string) {
	t.Helper()

	dir := filepath.Dir(path)
	if err := os.MkdirAll(dir, 0744); err != nil {
		t.Fatalf("failed to make the directory for golden file: %s", err)
	}
	if err := ioutil.WriteFile(path, actual.Bytes(), 0644); err != nil {
		t.Fatalf("failed to update golden file: %s", err)
	}

	t.Log(path + " is updated.")
}
