package main

import (
	"fmt"
	snet "github.com/zbh255/ss5-simple/net"
	"log"
	"net"
)

// chinese protocol doc link to: https://wanweibaike.net/wiki-SOCKS
// protocol doc english is RFC1928 & RFC1929
// simple test
func main() {
	listener, err := net.Listen("tcp", "127.0.0.1:1080")
	if err != nil {
		panic(err)
	}
	ss5Server := snet.NewNoAuthServer(listener)
	defer ss5Server.Close()
	_ = ss5Server.Start()
	for {
		ssConn, err := ss5Server.Connection()
		if err != nil {
			log.Printf("[Error]: %s", err.Error())
		}
		go HandlerConnection(ssConn)
	}
}



func HandlerConnection(conn snet.SSConn) {
	defer conn.Close()
	conn.RegisterConnectHandler(msgHandler,dataHandler)
	err := conn.Handler()
	if err != nil {
		fmt.Println(err.Error())
	}
}

func msgHandler(request *snet.Socks5MessageRequest) (*snet.Socks5MessageResponse, error) {
	fmt.Println(request)
	// response timeout
	response := new(snet.Socks5MessageResponse)
	response.Version = request.Version
	response.Reply = snet.SOCKS5_REPLY_SUCCESS
	response.Reserved = 0x00
	response.AddrType = request.AddrType
	response.Adders = request.Adders
	response.Port = request.Port
	return response,nil
}

func dataHandler(request []byte) ([]byte, error) {
	fmt.Println(string(request))
	responseBytes := []byte("<!DOCTYPE html>\n<html><!--STATUS OK--><head><meta http-equiv=\"Content-Type\" content=\"text/html; charset=utf-8\" /><meta http-equiv=\"Cache-control\" content=\"no-cache\" /><meta name=\"viewport\" content=\"width=device-width,minimum-scale=1.0,maximum-scale=1.0,user-scalable=no\"/><style type=\"text/css\">body {margin: 0;text-align: center;font-size: 14px;font-family: Arial,Helvetica,LiHei Pro Medium;color: #262626;}form {position: relative;margin: 12px 15px 91px;height: 41px;}img {border: 0}.wordWrap{margin-right: 85px;}#word {background-color: #FFF;border: 1px solid #6E6E6E;color: #000;font-size: 18px;height: 27px;padding: 6px;width: 100%;-webkit-appearance: none;-webkit-border-radius: 0;border-radius: 0;}.bn {background-color: #F5F5F5;border: 1px solid #787878;font-size: 16px;font-weight: bold;height: 41px;letter-spacing: -1px;line-height: 41px;padding: 0;position: absolute;right: 0;text-align: center;top: 0;width: 72px;-webkit-appearance: none;-webkit-box-sizing: border-box;box-sizing: border-box;-webkit-border-radius: 0;border-radius: 0;}.lg {margin-top: 30px;}a {text-decoration: none;color: #46446d;}.a {margin-top: 20px;}.d {margin: 58px 0 8px;}.b {margin-bottom: 19px;}.a a, .b a{text-align: center;display: inline-block;width: 58px;}.c a{margin: 0 10px;}.f{margin-top: 20px;color: #b4b4b4;font-size: 13px;}a.ip {padding-right: 18px;}span.ut {padding-left: 18px;border-left: 1px solid #dadada;}p:last-child a {color: #b4b4b4;}</style><title>百度一下,你就知道</title><script>function form_submit(){time=(new Date()).getTime().toString();time=time.substring(time.length-7);document.getElementsByName('ts')[0].value = time;return true;}</script></head><body><script type='text/javascript'>!function(){if(/#+.*?wd=([^&|$]+)/.test(location.href)&&RegExp.$1){var a=RegExp.$1,c=location.search,h=\"/s\"+c;h+=c&&c.indexOf(\"?\")>=0?\"&\":\"?\",h+=\"word=\"+a,location.replace(h);var b=document.createElement(\"meta\");b.setAttribute(\"http-equiv\",\"refresh\"),b.setAttribute(\"content\",\"0;URL=\"+h),document.head.appendChild(b),document.body.style.display=\"none\"}}();</script><div class=\"wrap\"><div class=\"lg\"><img src=\"//www.baidu.com/img/flexible/logo/utouch.png\" alt=\"百度首页\"/></div><form action=\"/from=844b/s\" method=\"get\" onsubmit=\"return form_submit()\"><div class=\"wordWrap\"><input type=\"text\" name=\"word\" maxlength=\"64\" size=\"17\" id=\"word\"/><input type=\"hidden\" value=\"ib\" name=\"sa\"/><input type=\"hidden\" value=\"0\" name=\"ts\"/><input type=\"hidden\" name=\"rsv_pq\" value=\"3756941177\"/><input type=\"hidden\" name=\"rsv_t\" value=\"125bVWdpWWF1oty4twl6PRJt3qPeXUG4dQjUMmLbZsXu7W5D16lMw6Fq7Q\"/><input type=\"hidden\" name=\"ms\" value=\"1\"/></div><input type=\"submit\" value=\"百度一下\" class=\"bn\"/></form><div class=\"a\"><a href=\"http://wapwenku.baidu.com/?statcms=index_wenku&amp;ssid=0&amp;from=844b&amp;bd_page_type=1&amp;uid=0&amp;pu=sz%401321_1001%2Cta%40utouch_2__24_90.0&amp;news?idx=30000&amp;itj=311\">文库</a><a href=\"http://image.baidu.com/search/wiseindex?tn=wiseindex&amp;wiseps=1\">图片</a><a href=\"http://zhidao.baidu.com/?idx=30000&amp;ssid=0&amp;from=844b&amp;bd_page_type=1&amp;uid=0&amp;pu=sz%401321_1001%2Cta%40utouch_2__24_90.0&amp;itj=34&amp;device=mobile\">知道</a><a href=\"http://m.news.baidu.com/news\">新闻</a><a href=\"http://wapbaike.baidu.com/?ssid=0&amp;from=844b&amp;bd_page_type=1&amp;uid=0&amp;pu=sz%401321_1001%2Cta%40utouch_2__24_90.0&amp;itj=31\">百科</a></div><div class=\"a\"><a href=\"http://mobile.baidu.com/simple/?action=index\">应用</a><a href=\"http://wapmap.baidu.com/?ssid=0&amp;from=844b&amp;bd_page_type=1&amp;uid=0&amp;pu=sz%401321_1001%2Cta%40utouch_2__24_90.0&amp;idx=30000&amp;itj=35&amp;wtj=wi\">地图</a><a href=\"http://tieba.baidu.com/?idx=30000&amp;ssid=0&amp;from=844b&amp;bd_page_type=1&amp;uid=0&amp;pu=sz%401321_1001%2Cta%40utouch_2__24_90.0&amp;itj=33&amp;mo_device=1\">贴吧</a><a href=\"http://m.hao123.com/?ssid=0&amp;from=844b&amp;bd_page_type=1&amp;uid=0&amp;pu=sz%401321_1001%2Cta%40utouch_2__24_90.0&amp;idx=30000&amp;itj=39\">hao123</a><a href=\"http%3A//m.baidu.com/pub/u_more.php?idx=30000&amp;ssid=0&amp;from=844b&amp;bd_page_type=1&amp;uid=0&amp;pu=sz%401321_1001%2Cta%40utouch_2__24_90.0&amp;itj=310&amp;device_os_id=2&amp;tj=utouch\">更多</a></div><div class=\"d\"><div class=\"b\"><a href=\"http://duokoo.baidu.com/novel/?fr=home&amp;ssid=0&amp;from=844b&amp;bd_page_type=1&amp;uid=0&amp;pu=sz%401321_1001%2Cta%40utouch_2__24_90.0\">小说</a><a href=\"http://duokoo.baidu.com/vgame/?fr=home&amp;ssid=0&amp;from=844b&amp;bd_page_type=1&amp;uid=0&amp;pu=sz%401321_1001%2Cta%40utouch_2__24_90.0\">游戏</a><a href=\"http://duokoo.baidu.com/vdl/?fr=home&amp;ssid=0&amp;from=844b&amp;bd_page_type=1&amp;uid=0&amp;pu=sz%401321_1001%2Cta%40utouch_2__24_90.0\">下载</a></div><div class=\"c\">下载：<a href=\"https://downpack.baidu.com/litebaiduboxapp_AndroidPhone_1021514q.apk\">百度客户端</a><a href=\"http://mobile.baidu.com/simple?action=index\">百度应用</a><a href=\"http://mo.baidu.com/map/?from=1708\">地图</a></div><div class=\"f\"><noscript><style>.ls {display: none;}</style></noscript><p class=\"ls\"><a id=\"switch-to-iphone\" class=\"ip\" href=\"http://m.baidu.com/?from=844b&amp;pu=sz%401320_480&amp;wpo=btmbase\">完整版</a></p><p><a href='http://m.baidu.com/pub/help.php?pn=25&amp;ssid=0&amp;from=844b&amp;uid=&amp;pu=sz%401321_1001%2Cta%40utouch_2__24_90.0&amp;bd_page_type=1'>Baidu&nbsp;京ICP证030173号</a></p></div><img src=\"//hpd.baidu.com/v.gif?tid=110&ref=index_utouch&from=0&pu=sz%401321_1001%2Cta%40utouch_2__24_90.0&qid=3756941177&logid=3756941177&ssid=0&ct=1&cst=1&logFrom=utouch&logInfo=utouch\" style=\"display:none;\"/><script>(function(){var switch_to_iphone = document.getElementById('switch-to-iphone');function setCookie(name, value, expires, path, domain, secure){var today = new Date();today.setTime(today.getTime());if (expires){expires = expires * 1000 * 60 * 60 * 24;}var expires_date = new Date(today.getTime() + (expires));document.cookie = name + \"=\" +escape( value ) +((expires) ? \";expires=\" + expires_date.toGMTString() : \"\" ) +((path) ? \";path=\" + path : \"\" ) +((domain) ? \";domain=\" + domain : \"\" ) +((secure) ? \";secure\" : \"\" );}switch_to_iphone.onclick = function(e){setCookie('utn', '1', 365, '/', '.baidu.com');};})();</script></div></div><script>window.addEventListener('load',function(){/*config*/var config = {'moplus' :{'url' : 'http://127.0.0.1:6259/getapn?','newurl' : 'http://127.0.0.1:40310/getapn?','params' : {'callback' : 'getapn','mcmdf' : 'inapp_test'}},'push':{'version': {'url' : 'http://127.0.0.1:7777/getPushServiceVersion?','params' : {'callback' : 'getPushServiceVersion'}},'network' : {'url' : 'http://127.0.0.1:7777/getNetworkType?','params' : {'callback' : 'getNetworkType'}}}};var conParams = function(obj){if(obj==undefined || typeof(obj)!==typeof({}))return '';var param = '';for(var k in obj){param += k + '=' + obj[k] + '&';}return param.slice(0,-1);};var flag = false, timeflag = false;var getJsonp = function(url, success, timeout, callback){var script = document.createElement(\"script\");script.onerror = script.onload = function(response){timeflag = true;script.onload = null;script.parentNode.removeChild(script);};window[callback] = success;script.src = url;document.body.appendChild(script);if(timeout > 0){setTimeout(function(){if(!timeflag){script.onload=null;script.parentNode.removeChild(script);}},timeout);}};/*set cookie*/var setCookie = function(name, value, path){var cookie = name + \"=\" + encodeURIComponent(value);cookie += \"; max-age=\" + 48000 + \"; path=\" + path;document.cookie = cookie;};/*html5 recogenized*/if(!flag && navigator.connection){nettype = navigator.connection.type;/*5:4g*/if(nettype == 5){setCookie('net', '4g-3', '/');flag = true;}}/*moplus*/if(!flag){var moplusSuccess = function(data){if (data && data.error == 0 && data.subtype && data.subtype.toLowerCase() == 'lte'){setCookie('net', '4g-4', '/');flag = true;}/*push*/if(!flag){var push = config.push;var versionSuccess = function(data){if(data && data.error == 0 && data.pushVersion >= 23){var networksuccess = function(data){if(data && data.error == 0 && data.networkType && data.networkType.search(/_13/) != -1){setCookie('net','4g-5', '/');}};var net = push.network;getJsonp(net.url + conParams(net.params), networksuccess, 2000, net.params.callback);}};var ver = push.version;getJsonp(ver.url + conParams(ver.params), versionSuccess, 2000, ver.params.callback);}};var moplus = config.moplus;getJsonp(moplus.url + conParams(moplus.params), moplusSuccess, 2000, moplus.params.callback);getJsonp(moplus.newurl + conParams(moplus.params), moplusSuccess, 2000, moplus.params.callback);}});</script></body></html>")
	return responseBytes,nil
}