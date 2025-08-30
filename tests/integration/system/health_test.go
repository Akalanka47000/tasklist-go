package system_test

import (
	"fmt"
	"github.com/samber/lo"
	. "github.com/smartystreets/goconvey/convey"
	"io"
	"net/http"
	"tasklist/tests/setup"
	"testing"
)

func TestSystemHealthHandler(t *testing.T) {
	t.Parallel()

	Convey("returns ok", t, func(c C) {
		app := ts.Prepare(t, c)

		req, _ := http.NewRequest(http.MethodGet, "/system/health", nil)
		res, err := app.Test(req)

		So(err, ShouldBeNil)

		defer res.Body.Close()

		So(res.StatusCode, ShouldEqual, http.StatusOK)

		body := string(lo.Ok(io.ReadAll(res.Body)))

		So(body, ShouldEqual, fmt.Sprintf("%s - OK", app.Config().AppName))
	})
}
