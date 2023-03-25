package crawler

import (
	"context"
	"github.com/chromedp/chromedp"
	"time"
)

func NewCourseTableCrawler(url string, account string, password string) CourseTableCrawler {
	ctx, _ := chromedp.NewExecAllocator(
		context.Background(),
		append(
			chromedp.DefaultExecAllocatorOptions[:],
			chromedp.UserAgent("Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.114 Safari/537.36 encors/0.0.6"),
			chromedp.Flag("headless", false),
		)...,
	)
	return CourseTableCrawler{
		ctx:      ctx,
		url:      url,
		account:  account,
		password: password,
	}
}

func Authorizer(url string, account string, password string) (err error) {
	crawler := NewCourseTableCrawler(url, account, password)
	var cancel context.CancelFunc
	crawler.ctx, cancel = context.WithTimeout(crawler.ctx, 60*time.Second)
	defer cancel()
	crawler.ctx, _ = chromedp.NewContext(crawler.ctx)

	err = crawler.loginTasks()
	if err != nil {
		return
	}
	return
}

func GetSemesterList(url string, account string, password string) (semesterList []Semester, err error) {
	crawler := NewCourseTableCrawler(url, account, password)
	var cancel context.CancelFunc
	crawler.ctx, cancel = context.WithTimeout(crawler.ctx, 60*time.Second)
	defer cancel()

	crawler.ctx, _ = chromedp.NewContext(crawler.ctx)
	err = crawler.loginTasks()
	if err != nil {
		return
	}

	semesterList, err = crawler.getSemesterList()
	if err != nil {
		return
	}
	return
}

func GetCourseTable(url string, account string, password string, semesterId string) (courseTable CourseTable, err error) {
	crawler := NewCourseTableCrawler(url, account, password)
	var cancel context.CancelFunc
	crawler.ctx, cancel = context.WithTimeout(crawler.ctx, 60*time.Second)
	defer cancel()

	crawler.ctx, _ = chromedp.NewContext(crawler.ctx)
	err = crawler.loginTasks()
	if err != nil {
		return
	}

	err = crawler.selectSemester(semesterId)
	if err != nil {
		return
	}

	courseTable, err = crawler.getCourseTable()
	if err != nil {
		return
	}
	return
}
