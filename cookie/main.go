package main

import (
	"fmt"
	"net/http"
	"time"
)

const (
	NAME        = "name"
	NAME_DOMAIN = "nameDomain"
	NAME_COUNT  = "nameCount"
	NAME_EXPIRE = "nameExpire"
	NAME_SECURE = "nameSecure"
	NAME_RAW    = "nameRaw"
)

var links string = "<div><a href=\"/\">top</a> <a href=\"/a\">/a</a> <a href=\"/a/b\">/a/b</a> <a href=\"/set\">cookieをセット</a> <a href=\"/delete\">cookieを削除</a></div>"

// Default Request Handler
func Handler(w http.ResponseWriter, r *http.Request) {
	// クッキーの取得
	value := cookieValue(r, NAME, "No Data.")
	valueDomain := cookieValue(r, NAME_DOMAIN, "No Data.")
	valueCount := cookieValue(r, NAME_COUNT, "No Data.")
	valueExpire := cookieValue(r, NAME_EXPIRE, "No Data.")
	valueSecure := cookieValue(r, NAME_SECURE, "ssl通信時のみ取得できます")
	valueRaw := cookieValue(r, NAME_RAW, "No Data.")

	s := links + "cookie : " + value + "<br>" +
		"cookieDomain : " + valueDomain + "<br>" +
		"cookieCount:" + valueCount + "<br>" +
		"cookieExpire:" + valueExpire + "<br>" +
		"cookieSecure:" + valueSecure + "<br>" +
		"cookieRaw:" + valueRaw + "<br>"

	fmt.Fprintf(w, s)
}

// Default Request Handler
func HandlerSet(w http.ResponseWriter, r *http.Request) {
	// クッキーの登録
	//	nowPlus := time.Now().Add(5 * time.Second)
	//	Path       string    // optional
	//	Domain     string    // optional
	//	Expires    time.Time // optional
	//	RawExpires string    // for reading cookies only
	//
	//	// MaxAge=0 means no 'Max-Age' attribute specified.
	//	// MaxAge<0 means delete cookie now, equivalently 'Max-Age: 0'
	//	// MaxAge>0 means Max-Age attribute present and given in seconds
	//	MaxAge   int
	//	Secure   bool
	//	HttpOnly bool
	//	Raw      string
	//	Unparsed []string // Raw text of unparsed attribute-value pairs
	cookies := []*http.Cookie{
		// Nameは取得時のキーとして作用しますが、取得時の状況によって対象が変化します
		// cookiの本来的なパラメータとしては、Name,Path,Domainで一つのキーとして作用しています
		// 方法はあるのかもしれませんが、valueに日本語は指定できないみたいです
		// Pathは前方一致
		&http.Cookie{Name: NAME, Value: "value"},
		&http.Cookie{Name: NAME, Value: "valueA", Path: "/a"},
		&http.Cookie{Name: NAME, Value: "valueB", Path: "/b"},
		&http.Cookie{Name: NAME, Value: "valueAB", Path: "/a/b"},
		&http.Cookie{Name: NAME_DOMAIN, Value: "valueDomain", Domain: "127.0.0.1"},
		// キーの生存可能経過時間
		&http.Cookie{Name: NAME_COUNT, Value: "MaxAge is can Access second", MaxAge: 5},
		// キーの削除時間指定
		&http.Cookie{Name: NAME_EXPIRE, Value: "5 second", Expires: time.Now().Add(time.Second * 5)},
		// httpsのみ
		&http.Cookie{Name: NAME_SECURE, Value: "secure", Secure: true},
		// ???
		&http.Cookie{Name: NAME_RAW, Value: "raw", Raw: "raw"},
	}

	for _, c := range cookies {
		http.SetCookie(w, c)
	}

	fmt.Fprintf(w, links+"cookiに値をセットしました")
}

// Default Request Handler
func HandlerDelete(w http.ResponseWriter, r *http.Request) {

	// クッキーの削除
	cookies := []*http.Cookie{
		&http.Cookie{Name: NAME, MaxAge: -1},
		&http.Cookie{Name: NAME, Path: "/a", MaxAge: -1},
		&http.Cookie{Name: NAME, Path: "/b", MaxAge: -1},
		&http.Cookie{Name: NAME_DOMAIN, Domain: "127.0.0.1", MaxAge: -1},
		&http.Cookie{Name: NAME_SECURE, MaxAge: -1},
		&http.Cookie{Name: NAME_RAW, MaxAge: -1},
	}

	for _, c := range cookies {
		http.SetCookie(w, c)
	}

	fmt.Fprintf(w, links+"cookieを削除しました<br>ただし、/a/bの値は残しているので,確認してみてください")
}

func cookieValue(r *http.Request, name string, def string) (value string) {
	cookie, err := r.Cookie(name)
	if err == nil {
		value = cookie.Value
	} else {
		value = def
	}
	return
}

func main() {
	http.HandleFunc("/", Handler)
	http.HandleFunc("/a", Handler)
	http.HandleFunc("/a/b", Handler)
	http.HandleFunc("/set", HandlerSet)
	http.HandleFunc("/delete", HandlerDelete)

	http.ListenAndServe("127.0.0.1:8080", nil)
}
