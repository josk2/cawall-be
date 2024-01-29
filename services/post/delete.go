package post

import "cawall-be/models"

func (postService *Service) Delete(post *models.Post) {
	postService.DB.Delete(post)
}
