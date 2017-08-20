package detect

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var expected = []struct {
	ua       string
	platform string
}{
	// iPhone
	{"Mozilla/5.0 (iPhone; U; CPU like Mac OS X; en) AppleWebKit/420.1 (KHTML, like Gecko) Version/3.0 Mobile/4A102 Safari/419", "iOS"},
	{"Mozilla/5.0 (iPhone; CPU iPhone OS 7_0_3 like Mac OS X) AppleWebKit/537.51.1 (KHTML, like Gecko) Version/7.0 Mobile/11B511 Safari/9537.53", "iOS"},
	// iPad
	{"Mozilla/5.0 (iPad; U; CPU OS 3_2 like Mac OS X; en-us) AppleWebKit/531.21.10 (KHTML, like Gecko) Version/4.0.4 Mobile/7B367 Safari/531.21.10", "iOS"},
	// iPod
	{"Mozilla/5.0 (iPod; U; CPU like Mac OS X; en) AppleWebKit/420.1 (KHTML, like Gecko) Version/3.0 Mobile/4A102 Safari/419", "iOS"},
	// Android misc.
	{"Opera/9.80 (Android 4.2.1; Linux; Opera Mobi/ADR-1212030829) Presto/2.11.355 Version/12.10", "Android"},
	{"Mozilla/5.0 (Linux; Android 4.2.1; Galaxy Nexus Build/JOP40D) AppleWebKit/535.19 (KHTML, like Gecko) Chrome/18.0.1025.166 Mobile Safari/535.19", "Android"},
	{"Mozilla/5.0 (Linux; U; Android 1.5; de-; HTC Magic Build/PLAT-RC33) AppleWebKit/528.5+ (KHTML, like Gecko) Version/3.1.2 Mobile Safari/525.20.1", "Android"},
	{"Mozilla/5.0 (Android; Tablet; rv:26.0) Gecko/26.0 Firefox/26.0", "Android"},
	{"Mozilla/5.0 (Android; Mobile; rv:17.0) Gecko/17.0 Firefox/17.0", "Android"},
	// BlackBerry
	{"Mozilla/5.0 (BlackBerry; U; BlackBerry 9800; en) AppleWebKit/534.1+ (KHTML, Like Gecko) Version/6.0.0.141 Mobile Safari/534.1+", "BlackBerry"},
	// BB10
	{"Mozilla/5.0 (BB10; Touch) AppleWebKit/537.3+ (KHTML, like Gecko) Version/10.0.9.388 Mobile Safari/537.3+", "Unknown Platform"},
	// webOS
	{"Mozilla/5.0 (webOS/1.4.0; U; en-US) AppleWebKit/532.2 (KHTML, like Gecko) Version/1.0 Safari/532.2 Pre/1.1", "Unknown Platform"},
	// Symbian
	{"Mozilla/5.0 (SymbianOS/9.1; U; [en-us]) AppleWebKit/413 (KHTML, like Gecko) Safari/413", "Symbian OS"},
	{"Mozilla/5.0 (SymbianOS/9.4; U; Series60/5.0 Profile/MIDP-2.1 Configuration/CLDC-1.1) AppleWebKit/525 (KHTML, like Gecko) Version/3.0 Safari/525", "Symbian OS"},
	// Firefox OS
	{"Mozilla/5.0 (Mobile; rv:26.0) Gecko/26.0 Firefox/26.0", "Unknown Platform"},
	{"Mozilla/5.0 (Tablet; rv:26.0) Gecko/26.0 Firefox/26.0", "Unknown Platform"},
	// Windows phone
	{"Mozilla/4.0 (compatible; MSIE 7.0; Windows Phone OS 7.0; Trident/3.1; IEMobile/7.0; SAMSUNG; SGH-i917)", "Windows Phone"},
	{"Mozilla/5.0 (compatible; MSIE 9.0; Windows Phone OS 7.5; Trident/5.0; IEMobile/9.0)", "Windows Phone"},
	{"HTC_Touch_3G Mozilla/4.0 (compatible; MSIE 6.0; Windows CE; IEMobile 7.11)", "Unknown Windows OS"},
	// Desktop
	{"Opera/9.80 (X11; Linux x86_64) Presto/2.12.388 Version/12.10", "Linux"},
	{"Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/535.19 (KHTML, like Gecko) Chrome/18.0.1025.168 Safari/535.19", "Windows 7"},
	{"", "Unknown Platform"},
	{"Mozilla/5.0 (Macintosh; U; Intel Mac OS X 10.6; en-US; rv:1.9.2.13) Gecko/20101203 Firefox/3.6.13", "Mac OS X"},
	{"alksjdlakdj", "Unknown Platform"},
	{"Mozilla/5.0 (compatible; MSIE 10.0; Windows NT 6.2; ARM; Trident/6.0; Touch)", "Windows 8"},
	{"Mozilla/4.0 (compatible; MSIE 7.0; Windows NT 6.3; Trident/7.0)", "Windows 8.1"},
	{"Mozilla/5.0 (Windows NT 6.1; WOW64; Trident/7.0; MDDRJS; rv:11.0) like Gecko", "Windows 7"},
	{"Mozilla/5.0 (Windows NT 6.1; WOW64; Trident/7.0; rv:11.0) like Gecko", "Windows 7"},
	{"Mozilla/5.0 (Windows NT 6.3; Trident/7.0; rv:11.0) like Gecko", "Windows 8.1"},
	{"Mozilla/5.0 (Windows NT 6.1; Trident/7.0; rv:11.0) like Gecko", "Windows 7"},
	{"Mozilla/5.0 (Windows NT 6.3; Trident/7.0; rv:11.0) like Gecko", "Windows 8.1"},
	{"Mozilla/4.0 (compatible; MSIE 7.0; Windows NT 6.1; WOW64; Trident/6.0; SLCC2; .NET CLR 2.0.50727; .NET CLR 3.5.30729; .NET CLR 3.0.30729; Media Center PC 6.0; .NET4.0C; .NET4.0E; InfoPath.3; MS-RTC LM 8)", "Windows 7"},
	{"Mozilla/4.0 (compatible; MSIE 7.0; Windows NT 6.1; WOW64; Trident/4.0; SLCC2; .NET CLR 2.0.50727; .NET CLR 3.5.30729; .NET CLR 3.0.30729; Media Center PC 6.0; .NET4.0C; .NET4.0E; InfoPath.3; MS-RTC LM 8)", "Windows 7"},
	{"Mozilla/4.0 (compatible; MSIE 7.0; Windows NT 6.2; ARM; Trident/6.0; Touch; .NET4.0E; .NET4.0C; Tablet PC 2.0)", "Windows 8"},
	{"Mozilla/5.0 (compatible; MSIE 10.0; Windows NT 6.2; Trident/6.0)", "Windows 8"},
}

func TestSetPlatform(t *testing.T) {
	for _, test := range expected {
		u := New(test.ua)
		u.setPlatform()
		t.Run(fmt.Sprintf("(UserAgent: %s) platform is %s", test.ua, test.platform), func(t *testing.T) {
			assert.Equal(t, test.platform, u.PlatForm)
		})
	}
}
