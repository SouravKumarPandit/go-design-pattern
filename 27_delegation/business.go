package main

import (
	"fmt"
	"strings"
)

type IBusinessService interface {
	doProcessing()
}

type EJBService struct {
}

func (ejb EJBService) doProcessing() {
	fmt.Println("Processing Task by invoking EJB service")
}

type JMSService struct {
}

func (ejb JMSService) doProcessing() {
	fmt.Println("Processing Task by invoking JMS service")
}

type BusinessLookup struct {
}

func (bl BusinessLookup) getBusinessService(serviceType string) IBusinessService {
	if strings.EqualFold("EJB", serviceType) {
		return EJBService{}
	} else {
		return JMSService{}
	}

}

type BusinessDelegate struct {
	lookupService BusinessLookup
	businessService IBusinessService
	serviceType string
}

func NewBusinessDelegate() BusinessDelegate {
	return BusinessDelegate{lookupService: BusinessLookup{}}
}

func (bsd *BusinessDelegate) setServiceType(serviceType string)  {
	bsd.serviceType=serviceType
}

func (bsd BusinessDelegate) doTask()  {
	bsd.businessService = bsd.lookupService.getBusinessService(bsd.serviceType)
	bsd.businessService.doProcessing()

}

type Client struct {
	businessDelegate *BusinessDelegate
}

func NewClient(bs *BusinessDelegate) Client {
	return Client{businessDelegate: bs}
}

func (c Client) doTask() {
	c.businessDelegate.doTask()
}