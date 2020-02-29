package service

import (
	"fmt"
	"log"
	"net/http"
	"os"
	//"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/langaner/crawlerdetector"

	dao "url-shortener-go/dao"
	//sdk "url-shortener-go/sdk"
	utilities "url-shortener-go/utilities"
)


func skipEvent(c *gin.Context) bool {
	skipEvent := c.Query("s")
	return skipEvent != "T"
}

func isCrawlerForUA(c *gin.Context) bool {
	detector := crawlerdetector.New()
	isCrawler := detector.IsCrawler(c.Request.Header.Get("User-Agent"))
	return isCrawler
}

// HandleLinkCode 处理短链Code
func HandleLinkCode(c *gin.Context) {
	var shortCode = c.Param("shortCode")
	var link, _ = getLinkByID(shortCode)
	// 找不到短链记录
	if link.ID == 0 {
		log.Println(fmt.Sprintf("shortCode:%s not found.", shortCode))
		c.JSON(404, gin.H{
			"code":  -1,
			"error": shortCode + " not found",
		})
		return
	}

	//shortCodePrefix := os.Getenv("LINK_CODES_PREFIX")
	//urlShortenerHost := os.Getenv("URL_SHORTER_DOMAIN")
	// 其他业务场景需要判断是不是爬虫，以及触发一些其他的事件
	// isCrawler := isCrawlerForUA(c)
	// if isCrawler {
	// 	log.Println(fmt.Sprintf("shortCode:%s,User-Agent:%s,Request is crawler,auto redirect.", shortCode, c.Request.Header.Get("User-Agent")))
	// }
	// var bcLinkCodes, _ = c.Cookie(shortCodePrefix)
	// if strings.Contains(bcLinkCodes, shortCode) {
	// 	log.Println(fmt.Sprintf("shortCode:%s in bcLinkCodes,auto redirect.", shortCode))
	// }
	// // 非爬虫 + 没有访问过 触发MQ时间
	// if !isCrawler && !strings.Contains(bcLinkCodes, shortCode) && !skipEvent(c){
	// 	log.Println(fmt.Sprintf("shortCode:%s not in bcLinkCodes,try to send share_link event.", shortCode))
	// 	bcLinkCodes = bcLinkCodes + shortCode + "|"
	// 	c.SetCookie(shortCodePrefix, bcLinkCodes, 60*60*24*30, "", urlShortenerHost, false, false)
	// 	// sdk.SendShareLinkMessage(*link)
	// }

	c.Redirect(http.StatusFound, link.URL)
}

// HandleCreateLink 处理新增短链
func HandleCreateLink(c *gin.Context) {
	var reqInfo dao.DBLink
	err := c.BindJSON(&reqInfo)
	if err != nil {
		log.Println(err)
		c.JSON(400, gin.H{"code": -1, "description": "Post Data Err"})
		return
	}
	var link, _ = createLink(reqInfo.OpenID, reqInfo.URL)
	// shortCodePrefix := os.Getenv("LINK_CODES_PREFIX")
	// urlShortenerHost := os.Getenv("URL_SHORTER_DOMAIN")
	// cookieMaxAge := 60 * 60 * 24 * 30
	// var bcLinkCodes, _ = c.Cookie(shortCodePrefix)
	// bcLinkCodes = bcLinkCodes + link.ShortCode + "|"
	// c.SetCookie(shortCodePrefix, bcLinkCodes, cookieMaxAge, "", urlShortenerHost, false, false)
	//log.Println(fmt.Sprintf("create link success,shortCode:%s,reqInfo.OpenID:%s,reqInfo.URL:%s", link.ShortCode, reqInfo.OpenID, reqInfo.URL))
	urlShortenerHost := os.Getenv("URL_SHORTER_DOMAIN")
	data := make(map[string]interface{})
	data["shareLink"] = urlShortenerHost + "/" + link.ShortCode
	data["shortCode"] = link.ShortCode
	c.JSON(200, gin.H{"code": 0, "data": data})
}

// createLink 创建一个Link
func createLink(openID string, url string) (link *dao.DBLink, err error) {
	link = new(dao.DBLink)
	link, _ = dao.GetLinkByURLAndOpenID(openID, url)
	if link.ID > 0 {
		log.Println(fmt.Sprintf("openID:%s, url:%s in db, skip it.", openID, url))
		return
	}
	link.OpenID = openID
	link.URL = url
	link.CreateTime = time.Now()
	link.UpdateTime = time.Now()
	err = dao.CreateLink(link)
	return
}

// getLinkByID 查找一个Link
func getLinkByID(shortCode string) (link *dao.DBLink, err error) {
	n := 62
	var linkID int = utilities.AnyToDecimal(shortCode, n)
	link, err = dao.GetLinkByID(int64(linkID))
	return
}
