package dao

import (
	"quick_find_backend/common"
	"quick_find_backend/topic/model"
)

func (d *Dao) Insert(topic *model.Topic) error {
	_, err := d.sql.Insert(topic)
	if err != nil {
		common.Logger.Error("[d.topic] Push(%#v) error:%s", topic, err.Error())
		return err
	}
	return nil
}

func (d *Dao) List(start, count int) ([]*model.TopicJoinUser, error) {
	topic := make([]*model.TopicJoinUser, 0)

	if err := d.sql.Limit(count, start).
		Table("topic").
		Join("inner","user","user.uid=topic.uid").
		Find(&topic); err != nil {
		common.Logger.Error("[d.topic] List(%d,%d) error:%s", start, count, err.Error())
		return nil, err
	}
	return topic, nil
}





