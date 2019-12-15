package quick_find_services_topic

import (
	"context"
	"quick_find_backend/topic/dao"
	"quick_find_backend/topic/model"
	"time"
)

type TopicService struct {
	topicDao *dao.Dao
}

func NewTopicService(dao *dao.Dao)*TopicService{
	return &TopicService{
		topicDao:dao,
	}
}

func (ts *TopicService) Push(context context.Context, res *Publish, resp *PushResponse) error {
	push := &model.Topic{
		Uid:        res.Uid,
		Title:      res.Title,
		Content:    res.Content,
		Like:       0,
		CreateTime: time.Now().Format("2006-01-02 15:04:05"),
	}
	if err := ts.topicDao.Insert(push); err != nil {
		return err
	}
	return nil
}

func(ts *TopicService)List(context context.Context, res *ListRequest,resp *TopicList)error{
	start:=(res.Page-1)*res.Count
	list,err:=ts.topicDao.List(int(start),int(res.Count))
	if err!=nil{
		return err
	}

	*resp=TopicList{
		List:make([]*TopinJoinName,len(list)),
	}

	for i:=0;i<len(list);i++{
		single:=&TopinJoinName{
			Title:list[i].Title,
			Uid:list[i].Uid,
			Content:list[i].Content,
			Like:list[i].Like,
			Name:list[i].Name,
		}
		resp.List[i]=single
	}


	return nil
}