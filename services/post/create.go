package post

import "cawall-be/models"

func (postService *Service) Create(post *models.Post) {
	postService.DB.Create(post)
}
