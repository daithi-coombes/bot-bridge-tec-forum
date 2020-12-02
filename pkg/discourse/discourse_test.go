package discourse

import (
	"bytes"
	"io/ioutil"
	"log"
	"net/http"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUpdatePost(t *testing.T) {

	mockClient := helperGetClient(func(req *http.Request) *http.Response {
		resp, _ := helperLoadFixture("fixture_DiscoursePutPost.repsonse.json")
		return &http.Response{
			StatusCode: 200,
			Body:       ioutil.NopCloser(bytes.NewBuffer(resp)),
			Header:     make(http.Header),
		}
	})
	underTest := NewDiscourse(
		"http://example.com",
		"435947dc08168807254198c9b8701b282e2f8b2722caaef7ae3171900418e18b",
		mockClient,
	)

	post := Post{
		ID:  14,
		Raw: `<p><blockquote>Propsal submitted, check status here: http://alksdfjl.alsdkfj.alaksdjf</blockquote></p><hr/><p>this is the updated content (Post.Raw)</p><blockquote>This is the blockquote?</blockquote>`,
	}
	actual, err := underTest.UpdatePost(post)
	if err != nil {
		t.Error(err)
	}
	expected := Response{
		Post: post,
	}

	assert.Equal(t, expected.Post.Raw, actual.Post.Cooked)
}

func TestGetPost(t *testing.T) {

	mockClient := helperGetClient(func(req *http.Request) *http.Response {
		resp, _ := helperLoadFixture("fixture_DiscorseGetPost.response.json")
		return &http.Response{
			StatusCode: 200,
			Body:       ioutil.NopCloser(bytes.NewBuffer(resp)),
			Header:     make(http.Header),
		}
	})
	underTest := NewDiscourse(
		"http://example.com",
		"435947dc08168807254198c9b8701b282e2f8b2722caaef7ae3171900418e18b",
		mockClient,
	)

	postURL := "http://example.com/t/once-upon-a-plop/11.json"
	actual, err := underTest.GetPost(postURL)
	if err != nil {
		t.Error(err)
	}

	expected := Post{
		ID:     14,
		Cooked: "<p>there was madness going on</p>",
	}
	assert.Equal(t, expected, actual)
}

type RoundTripFunc func(req *http.Request) *http.Response

func (f RoundTripFunc) RoundTrip(req *http.Request) (*http.Response, error) {
	return f(req), nil
}
func helperGetClient(fn RoundTripFunc) *http.Client {
	return &http.Client{
		Transport: RoundTripFunc(fn),
	}
}
func helperLoadFixture(filename string) ([]byte, error) {
	var j []byte
	absPath, err := filepath.Abs("../../tests")
	if err != nil {
		return j, err
	}
	fixture := absPath + "/" + filename
	log.Printf("loading fixture: %s\n", fixture)
	j, err = ioutil.ReadFile((fixture))
	if err != nil {
		return j, err
	}
	return j, nil
}
