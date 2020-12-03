package discourse

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/daithi-coombes/bot-bridge-tec-forum/pkg/dao"
)

// Discourse The main struct for this package
type Discourse struct {
	c        *http.Client
	endpoint string
	apiKey   string
}

// Post The struct for handling post object from discorse
type Post struct {
	ID         float64 `json:"id"`
	Raw        string  `json:"raw,omitempty"`
	RawOld     string  `json:"raw_old,omitempty"`
	Cooked     string  `json:"cooked,omitempty"`
	EditReason string  `json:"edit_reason,omitempty"`
}

// PostStream The PUT payload
type PostStream struct {
	Posts []Post `json:"posts"`
}

// Response The response struct
type Response struct {
	PostStream PostStream `json:"post_stream,omitempty"`
	Post       Post       `json:"post,omitempty"`
}

// NewDiscourse Factory method
func NewDiscourse(endpoint string, apiKey string, client *http.Client) Discourse {
	return Discourse{
		c:        client,
		endpoint: endpoint,
		apiKey:   apiKey,
	}
}

// GetPost Get the json representation of a post from its URL
func (d *Discourse) GetPost(url string) (Post, error) {

	var post Post

	log.Printf("Getting post: %s\n", url)
	req, err := http.NewRequest("GET", url+".json", nil)
	if err != nil {
		return post, err
	}
	d.SetHeaders(req)

	resp, err := d.c.Do(req)
	if err != nil {
		return post, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return post, err
	}

	postStream := Response{}
	if err := json.Unmarshal(body, &postStream); err != nil {
		return post, err
	}

	if len(postStream.PostStream.Posts) == 0 {
		return post, fmt.Errorf("No posts found")
	}

	return postStream.PostStream.Posts[0], nil
}

// HandleProposal Built to be used as goroutine
func (d *Discourse) HandleProposal(p dao.ProposalAdded) error {

	log.Printf("Handling proposal: %s\n", p.Link)
	// 1. when recieve proposal
	post, err := d.GetPost(string(p.Link))
	if err != nil {
		return err
	}

	newBody := "<p><blockquote>Propsal submitted, check status here: http://alksdfjl.alsdkfj.alaksdjf</blockquote></p><hr/>" + post.Cooked
	post.Raw = newBody
	post.EditReason = "updating with submitted proposal"
	_, err = d.UpdatePost(post)
	if err != nil {
		log.Fatal(err)
	}

	return nil
}

// UpdatePost Save changes to a post
func (d *Discourse) UpdatePost(post Post) (Response, error) {

	var update Response
	log.Println("updateing post...")

	url := d.endpoint + "/posts/" + fmt.Sprintf("%d", int(post.ID)) + ".json"
	putBody, err := json.Marshal(post)
	if err != nil {
		return update, err
	}
	req, err := http.NewRequest("PUT", url, bytes.NewBuffer(putBody))
	log.Printf("PUT %s\n", url)
	if err != nil {
		return update, err
	}
	d.SetHeaders(req)

	resp, err := d.c.Do(req)
	if err != nil {
		return update, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return update, err
	}

	if err := json.Unmarshal(body, &update); err != nil {
		return update, err
	}

	return update, nil
}

// SetHeaders Set the headers for requests
func (d *Discourse) SetHeaders(req *http.Request) {
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Api-Key", d.apiKey)
	req.Header.Set("Api-Username", "system")
}
