package service

import "blog_web/db/dao"

/*
* @Author: mgh
* @Date: 2022/3/19 11:33
* @Desc:
 */

type BGImageService struct {
	bgImageDao *dao.BGImageDao
}

func NewBGImageService() *BGImageService {
	return &BGImageService{
		bgImageDao: dao.NewBGImageDao(),
	}
}

func (bg *BGImageService) GetAllUrl() ([]string, error) {
	return bg.bgImageDao.FindAllURL()
}

func (bg *BGImageService) AddUrl(url string) error {
	return bg.bgImageDao.AddUrl(url)
}

func (bg *BGImageService) DeleteUrl(id uint32) error {
	return bg.bgImageDao.DeleteUrl(id)
}

func (bg *BGImageService) UpdateUrl(id uint32, url string) error {
	return bg.bgImageDao.UpdateUrl(id, url)
}