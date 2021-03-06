package main

import (
	"io"
	"net/http"
	"os"
	"time"
)

func main() {
	httpGet()
}

func httpGet() {
	c := http.Client{Timeout: time.Duration(3) * time.Second}
	req, err := http.NewRequest(http.MethodGet, url, nil)
	checkErr(err)
	setUpHeader(req)
	resp, err := c.Do(req)
	checkErr(err)
	_, err = io.Copy(os.Stdout, resp.Body)
	checkErr(err)
	defer resp.Body.Close()
}

var (
	url         = `https://tiki.vn/api/v2/orders?page=0&limit=10`
	accessToken = `eyJhbGciOiJSUzI1NiJ9.eyJzdWIiOiIxNDU4MjI3MCIsImlhdCI6MTY0NjgxODg1MywiZXhwIjoxNjQ2OTkxNjUzLCJpc3MiOiJodHRwczovL3Rpa2kudm4iLCJjdXN0b21lcl9pZCI6IjE0NTgyMjcwIiwiY2xpZW50X2lkIjoidGlraS1zc28iLCJzY29wZSI6InNzbyJ9.FaBcCbH9YNpeL5BQeVzIsLFkPDksT1o_qWYdjIiFBOfjNhO974iYgHcKC4UtCvD1z_KwTjPL0sQSiVHrvVA8AcN6yiTVOB0g50aAuC9HkqYNMd54rfz-R9XabNT-fRCzPOh9RL-30GhtTvXCkGrPT7_yy5kujCbJedOeutabn7bwbz4KkVlhn-SHftjHma32Omw-fwN-hZu5vVTJ9OKt0yTFqeHtf4xge4Oea7uajU2k53Dz5kga6mrRiP7R_nYKc8zzUaKtkNNDwN3g-2btD_3YLXZrHF8xGM1yHA00U6dqA7rOh9cCt0S9zs-2GeRLPC22grS8uPOo7RZ2k6w-G66fqEKHzUcaLTcSdNzZIXK7-ekRnMq2WkllH95F_Qnj1sNfHuwU_BO5Mcd2FHPlTGe3La7zXw7lGMkE9_E4SkZhpVA3Aq9Zepd4PVOg_lHp3p9I9aJEiSk5ULGjOfnVSbnDg93fgluLzVULVzNRC6YQHOBS4Z8dt3UzfvYAPZZYC2ue3zT18pRiAU9cudwGFej8uGpUcuVtB8YfYbx7NOCCGOXu_GG-nNdiebccrvnDQTP_bdRyrD7gTdBYfMcomP38gKV0Gv3Y-kl2l1pZdLKO7Wlevfy_867dC2u5-SQWzvy5PnA0ACIB48KdrJWJVwzXoHNifqtj7WQF8hMwv_4`
	userAngent  = `Mozilla/5.0 (Macintosh; Intel Mac OS X 10.15; rv:98.0) Gecko/20100101 Firefox/98.0`
	cookie      = `_trackity=ab78e012-1ccb-fb52-5bab-9ed60fbbba62; TOKENS={%22access_token%22:%22eyJhbGciOiJSUzI1NiJ9.eyJzdWIiOiIxNDU4MjI3MCIsImlhdCI6MTY0NjgxODg1MywiZXhwIjoxNjQ2OTkxNjUzLCJpc3MiOiJodHRwczovL3Rpa2kudm4iLCJjdXN0b21lcl9pZCI6IjE0NTgyMjcwIiwiY2xpZW50X2lkIjoidGlraS1zc28iLCJzY29wZSI6InNzbyJ9.FaBcCbH9YNpeL5BQeVzIsLFkPDksT1o_qWYdjIiFBOfjNhO974iYgHcKC4UtCvD1z_KwTjPL0sQSiVHrvVA8AcN6yiTVOB0g50aAuC9HkqYNMd54rfz-R9XabNT-fRCzPOh9RL-30GhtTvXCkGrPT7_yy5kujCbJedOeutabn7bwbz4KkVlhn-SHftjHma32Omw-fwN-hZu5vVTJ9OKt0yTFqeHtf4xge4Oea7uajU2k53Dz5kga6mrRiP7R_nYKc8zzUaKtkNNDwN3g-2btD_3YLXZrHF8xGM1yHA00U6dqA7rOh9cCt0S9zs-2GeRLPC22grS8uPOo7RZ2k6w-G66fqEKHzUcaLTcSdNzZIXK7-ekRnMq2WkllH95F_Qnj1sNfHuwU_BO5Mcd2FHPlTGe3La7zXw7lGMkE9_E4SkZhpVA3Aq9Zepd4PVOg_lHp3p9I9aJEiSk5ULGjOfnVSbnDg93fgluLzVULVzNRC6YQHOBS4Z8dt3UzfvYAPZZYC2ue3zT18pRiAU9cudwGFej8uGpUcuVtB8YfYbx7NOCCGOXu_GG-nNdiebccrvnDQTP_bdRyrD7gTdBYfMcomP38gKV0Gv3Y-kl2l1pZdLKO7Wlevfy_867dC2u5-SQWzvy5PnA0ACIB48KdrJWJVwzXoHNifqtj7WQF8hMwv_4%22%2C%22refresh_token%22:%22TKIAEnvoh15yQhVDfhowTBwJMmXN7N8aiGx6oOBOdpHH6ldWLEgbr3p-JOhtM3eKUFBzupRZgFPXbdQdi9el%22%2C%22token_type%22:%22bearer%22%2C%22expires_in%22:172800%2C%22expires_at%22:1646991653562%2C%22customer_id%22:14582270}; delivery_zone=Vk4wMzkwMDYwMDE=; amp_99d374=kb-M6pUZpR7aUEKxccRo9g.MTQ1ODIyNzA=..1ftn0pt61.1ftn0r4s8.5.8.d; _ga=GA1.2.556715556.1646818818; _gid=GA1.2.426044649.1646818818; _gcl_au=1.1.155953249.1646818818; _ga_GSD4ETCY1D=GS1.1.1646818818.1.1.1646818857.0; tiki_client_id=556715556.1646818818; _gat=1; _fbp=fb.1.1646818819209.1068670845; _hjSessionUser_522327=eyJpZCI6ImQ1MTk2YTI3LWEyNjQtNTViNS05ZmM5LThhYzUwODM3Mjc1MCIsImNyZWF0ZWQiOjE2NDY4MTg4MjA4MTksImV4aXN0aW5nIjp0cnVlfQ==; _hjFirstSeen=1; _hjIncludedInSessionSample=0; _hjSession_522327=eyJpZCI6IjliYzFkOTVmLTc5MWUtNGU3MS1iZGNkLWQ3MzgzNGI2ZTRkYSIsImNyZWF0ZWQiOjE2NDY4MTg4MjI2OTMsImluU2FtcGxlIjpmYWxzZX0=; _hjAbsoluteSessionInProgress=0; cto_bundle=WsbFW18lMkJtRmkxa0tsTVo5aGolMkYzRWZnMCUyQnB1S0FBMWpGZWJIdllza2NRVkhMd1FDYzExQWRLNDhaTFVrWkw4bFJHQmd2S2IlMkJFOFVKZWdhYTU0V3ZNTmVvJTJGRW1CQnYlMkYlMkJWVk95VXl5TEhJaDNwbUt1enlxdmdVbnFYcDglMkYyZDMxeWNpJTJGUXgyeVcwdE1haVhBNlVoSlhOR21HS3JCenFHQTBVVzBMWVkzNXNEaVRXN3RYQzBwYjVYbm8lMkJYMmV1SUpiJTJGZE92; __iid=749; __iid=749; __su=0; __su=0; G_ENABLED_IDPS=google; TIKI_FE_VARIANT=0d9ab821-16ce-4bb8-8237-a284211af1b5variant_code11654831-27f3-4ccb-bec5-bcd4c44e0ee8; TIKI_ACCESS_TOKEN=eyJhbGciOiJSUzI1NiJ9.eyJzdWIiOiIxNDU4MjI3MCIsImlhdCI6MTY0NjgxODg1MywiZXhwIjoxNjQ2OTkxNjUzLCJpc3MiOiJodHRwczovL3Rpa2kudm4iLCJjdXN0b21lcl9pZCI6IjE0NTgyMjcwIiwiY2xpZW50X2lkIjoidGlraS1zc28iLCJzY29wZSI6InNzbyJ9.FaBcCbH9YNpeL5BQeVzIsLFkPDksT1o_qWYdjIiFBOfjNhO974iYgHcKC4UtCvD1z_KwTjPL0sQSiVHrvVA8AcN6yiTVOB0g50aAuC9HkqYNMd54rfz-R9XabNT-fRCzPOh9RL-30GhtTvXCkGrPT7_yy5kujCbJedOeutabn7bwbz4KkVlhn-SHftjHma32Omw-fwN-hZu5vVTJ9OKt0yTFqeHtf4xge4Oea7uajU2k53Dz5kga6mrRiP7R_nYKc8zzUaKtkNNDwN3g-2btD_3YLXZrHF8xGM1yHA00U6dqA7rOh9cCt0S9zs-2GeRLPC22grS8uPOo7RZ2k6w-G66fqEKHzUcaLTcSdNzZIXK7-ekRnMq2WkllH95F_Qnj1sNfHuwU_BO5Mcd2FHPlTGe3La7zXw7lGMkE9_E4SkZhpVA3Aq9Zepd4PVOg_lHp3p9I9aJEiSk5ULGjOfnVSbnDg93fgluLzVULVzNRC6YQHOBS4Z8dt3UzfvYAPZZYC2ue3zT18pRiAU9cudwGFej8uGpUcuVtB8YfYbx7NOCCGOXu_GG-nNdiebccrvnDQTP_bdRyrD7gTdBYfMcomP38gKV0Gv3Y-kl2l1pZdLKO7Wlevfy_867dC2u5-SQWzvy5PnA0ACIB48KdrJWJVwzXoHNifqtj7WQF8hMwv_4; TIKI_USER=8GMJDM3pmjUaWTPehSNmsEVTa8Jc%2FF0xOo3Il5D9Yo%2B9A3tSpnfekbwwdfkMKqCfqcMchpe2K3NY`
)

func setUpHeader(req *http.Request) {
	req.Header.Add("User-Agent", userAngent)
	req.Header.Add("Accept", "application/json, text/plain, */*")
	req.Header.Add("x-access-token", accessToken)
	req.Header.Add("Sec-Fetch-Mode", "cors")
	req.Header.Add("Sec-Fetch-Site", "same-origin")
	req.Header.Add("Referer", "https://tiki.vn/sales/order/history?src=header_my_account")
	req.Header.Add("Connection", "keep-alive")
	req.Header.Add("Cookie", cookie)
}

func checkErr(err error) {
	if err != nil {
		panic(err.Error())
	}
}
