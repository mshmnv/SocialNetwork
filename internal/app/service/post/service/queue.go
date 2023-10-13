package service

//func (s *Service) ConsumeNewPost(delivery amqp.Delivery) error {
//	p := &datastruct.Post{}
//	if err := json.Unmarshal(delivery.Body, p); err != nil {
//		logger.Errorf("Error decoding post in consumer: %v", err)
//		return err
//	}
//
//	if err := s.HandleNewPost(p); err != nil {
//		logger.Errorf("Error handling new post: %v", err)
//		return err
//	}
//	return nil
//}
