package httd

import (
	"bytes"
	"compress/gzip"
	"io/ioutil"
	"net/http"
	"testing"
	"time"
)

func missingImplError(t *testing.T, interfaceName string) {
	t.Error("client{} does not implement " + interfaceName + " interface")
}

func TestClientImplementInterfaces(t *testing.T) {
	client := &Client{}
	if _, implemented := interface{}(client).(Requester); !implemented {
		missingImplError(t, "Requester")
	}
	if _, implemented := interface{}(client).(Getter); !implemented {
		missingImplError(t, "Getter")
	}
	if _, implemented := interface{}(client).(Poster); !implemented {
		missingImplError(t, "Poster")
	}
	if _, implemented := interface{}(client).(Puter); !implemented {
		missingImplError(t, "Puter")
	}
	if _, implemented := interface{}(client).(Patcher); !implemented {
		missingImplError(t, "Patcher")
	}
	if _, implemented := interface{}(client).(Deleter); !implemented {
		missingImplError(t, "Deleter")
	}
}

func TestDecodingResponseBody(t *testing.T) {
	expected := "oashoasihdosado4o5ry8wy34hr8w3yr88y3r9283y"
	client := &Client{}
	resp := &http.Response{
		Body: ioutil.NopCloser(bytes.NewBufferString(expected)),
	}
	defer resp.Body.Close()

	body, err := client.decodeResponseBody(resp)
	if err != nil {
		t.Error(err)
	}

	if string(body) != expected {
		t.Errorf("decoding failed. Got %s, wants %s", string(body), expected)
	}
}

func TestDecodingResponseBodyWithGZIP(t *testing.T) {
	expected := "9ng574g8573g394g3874gf837g"
	client := &Client{}
	resp := &http.Response{
		Body:   ioutil.NopCloser(bytes.NewBufferString(expected)),
		Header: make(http.Header, 0),
	}
	resp.Header.Set(ContentEncoding, GZIPCompression)

	// expect to fail as body is not gzip compressed yet
	_, err := client.decodeResponseBody(resp)
	if err == nil {
		t.Error("successfully decoded a none gzip encoded message using the gzip algorithm")
	}
	resp.Body.Close()

	// compress body using gzip
	var b bytes.Buffer
	gz := gzip.NewWriter(&b)
	if _, err := gz.Write([]byte(expected)); err != nil {
		gz.Close()
		t.Error("could not compress content using gzip")
	}
	gz.Close()

	resp.Body = ioutil.NopCloser(bytes.NewBufferString(b.String()))
	defer resp.Body.Close()

	// try decompressing gzip
	body, err := client.decodeResponseBody(resp)
	if err != nil {
		t.Error(err)
	}

	if string(body) != expected {
		t.Errorf("decoding failed. Got %s, wants %s", string(body), expected)
	}

}

func TestAdjustReset(t *testing.T) {
	reset := int64(1 * 1000)
	adjuster := func(d time.Duration) time.Duration {
		return d
	}

	newReset := adjustReset(reset, adjuster)
	if reset != newReset {
		t.Errorf("expects reset and newReset to be equal. Got %d, wants %d", newReset, reset)
	}

	// for reaction endpoint we will get 1 second rate limits, but we want to
	// set these to 250ms
	adjuster = func(d time.Duration) time.Duration {
		if d.Seconds() == 1 {
			d = time.Duration(250) * time.Millisecond
		}
		return d
	}

	newReset = adjustReset(reset, adjuster)
	wants := int64(250)
	if wants != newReset {
		t.Errorf("expects newReset to be 250ms. Got %d, wants %d", newReset, wants)
	}
}
