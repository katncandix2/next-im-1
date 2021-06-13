package oauth

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

const (
	DevClientID     = "c94c5bb56cc743f35e1e"
	DevClientSecret = "ad0cd8328a31293f21d5b227bbe30fded4188cbc"

	//todo online ClientID Secret 登记地址 https://github.com/settings/applications/new
	AccessTokenUrl = "https://github.com/login/oauth/access_token?client_id=%v&client_secret=%v&code=%v"

	AuthCodeUrl = "https://github.com/login/oauth/authorize?client_id=%v&redirect_uri=http://localhost:8080/oauth/redirect"
)

type Token struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	Scope       string `json:"scope"`
}

type UserGitHubMeta struct {
	Login     string `json:"login"`
	ID        int64  `json:"id"`
	NodeID    string `json:"node_id"`
	AvatarUrl string `json:"avatar_url"`
}

//code 在调试时可以直接访问如下地址获取
//https://github.com/login/oauth/authorize?client_id=c94c5bb56cc743f35e1e&redirect_uri=http://localhost:8080/oauth/redirect
func GetAccessToken(code string) *Token {
	url := fmt.Sprintf(AccessTokenUrl, DevClientID, DevClientSecret, code)
	payload := strings.NewReader("")
	client := &http.Client{}
	req, _ := http.NewRequest(http.MethodGet, url, payload)
	req.Header.Add("accept", "application/json")
	res, _ := client.Do(req)
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
	token := &Token{}
	_ = json.Unmarshal(body, token)
	return token
}

//示例json
//{"login":"katncandix2","id":10647224,"node_id":"MDQ6VXNlcjEwNjQ3MjI0","avatar_url":"https://avatars.githubusercontent.com/u/10647224?v=4","gravatar_id":"","url":"https://api.github.com/users/katncandix2","html_url":"https://github.com/katncandix2","followers_url":"https://api.github.com/users/katncandix2/followers","following_url":"https://api.github.com/users/katncandix2/following{/other_user}","gists_url":"https://api.github.com/users/katncandix2/gists{/gist_id}","starred_url":"https://api.github.com/users/katncandix2/starred{/owner}{/repo}","subscriptions_url":"https://api.github.com/users/katncandix2/subscriptions","organizations_url":"https://api.github.com/users/katncandix2/orgs","repos_url":"https://api.github.com/users/katncandix2/repos","events_url":"https://api.github.com/users/katncandix2/events{/privacy}","received_events_url":"https://api.github.com/users/katncandix2/received_events","type":"User","site_admin":false,"name":"guruiqin","company":"tencent","blog":"tencent","location":"beijing","email":"254971939@qq.com","hireable":null,"bio":"机器学习🤔\r\n","twitter_username":null,"public_repos":62,"public_gists":0,"followers":3,"following":35,"created_at":"2015-01-22T03:52:54Z","updated_at":"2021-06-13T08:07:43Z"}
func GetUserMeta(code string) *UserGitHubMeta {
	token := GetAccessToken(code)
	url := "https://api.github.com/user"
	client := &http.Client{}
	req, _ := http.NewRequest(http.MethodGet, url, nil)

	req.Header.Add("Accept", "application/vnd.github.v3+json")
	req.Header.Add("Authorization", "token "+token.AccessToken)
	res, _ := client.Do(req)
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
	umeta := &UserGitHubMeta{}
	json.Unmarshal(body, umeta)
	return umeta
}
