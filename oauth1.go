package discogs

import (
	"crypto/hmac"
	"crypto/rand"
	"crypto/sha1"
	"encoding/base64"
	"encoding/hex"
	"net/url"
	"sort"
	"strconv"
	"strings"
	"time"
)

func oauth1AuthorizationHeader(method, rawURL, consumerKey, consumerSecret, token, tokenSecret string, extra map[string]string) string {
	params := map[string]string{
		"oauth_consumer_key":     consumerKey,
		"oauth_nonce":            oauthNonce(),
		"oauth_signature_method": "HMAC-SHA1",
		"oauth_timestamp":        strconv.FormatInt(time.Now().Unix(), 10),
		"oauth_version":          "1.0",
	}
	if token != "" {
		params["oauth_token"] = token
	}
	for k, v := range extra {
		params[k] = v
	}

	baseString := oauthSignatureBaseString(method, rawURL, params)
	params["oauth_signature"] = oauthSignHMACSHA1(baseString, consumerSecret, tokenSecret)

	// Emit only oauth_* params.
	keys := make([]string, 0, len(params))
	for k := range params {
		if strings.HasPrefix(k, "oauth_") {
			keys = append(keys, k)
		}
	}
	sort.Strings(keys)

	var b strings.Builder
	b.WriteString("OAuth ")
	for i, k := range keys {
		if i > 0 {
			b.WriteString(", ")
		}
		b.WriteString(oauthPercentEncode(k))
		b.WriteString(`="`)
		b.WriteString(oauthPercentEncode(params[k]))
		b.WriteString(`"`)
	}
	return b.String()
}

func oauthSignatureBaseString(method, rawURL string, params map[string]string) string {
	u, _ := url.Parse(rawURL)
	query := u.Query()
	u.RawQuery = ""
	u.Fragment = ""

	type kv struct{ k, v string }
	kvs := make([]kv, 0, len(params)+len(query))
	for k, v := range params {
		if k == "oauth_signature" {
			continue
		}
		kvs = append(kvs, kv{k: oauthPercentEncode(k), v: oauthPercentEncode(v)})
	}
	for k, vs := range query {
		for _, v := range vs {
			kvs = append(kvs, kv{k: oauthPercentEncode(k), v: oauthPercentEncode(v)})
		}
	}
	sort.Slice(kvs, func(i, j int) bool {
		if kvs[i].k == kvs[j].k {
			return kvs[i].v < kvs[j].v
		}
		return kvs[i].k < kvs[j].k
	})

	var normalized strings.Builder
	for i, p := range kvs {
		if i > 0 {
			normalized.WriteByte('&')
		}
		normalized.WriteString(p.k)
		normalized.WriteByte('=')
		normalized.WriteString(p.v)
	}

	return strings.ToUpper(method) + "&" + oauthPercentEncode(u.String()) + "&" + oauthPercentEncode(normalized.String())
}

func oauthSignHMACSHA1(baseString, consumerSecret, tokenSecret string) string {
	key := oauthPercentEncode(consumerSecret) + "&" + oauthPercentEncode(tokenSecret)
	mac := hmac.New(sha1.New, []byte(key))
	_, _ = mac.Write([]byte(baseString))
	return base64.StdEncoding.EncodeToString(mac.Sum(nil))
}

func oauthNonce() string {
	var b [16]byte
	_, _ = rand.Read(b[:])
	return hex.EncodeToString(b[:])
}

func oauthPercentEncode(s string) string {
	esc := url.QueryEscape(s)
	esc = strings.ReplaceAll(esc, "+", "%20")
	esc = strings.ReplaceAll(esc, "%7E", "~")
	return esc
}

