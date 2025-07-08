package dto

import "draft-zadania-1/models"

func ToTask(dto TaskDTO) models.Task {
	return models.Task{
		Title:       dto.Title,
		Description: dto.Description,
		DueDate:     dto.DueDate,
		Status:      models.Status(dto.Status),
		UserId:      dto.UserId,
	}
}

func ToResponseTaskDTO(task models.Task) ResponseTaskDTO {
	return ResponseTaskDTO{
		Id:          task.Id,
		Title:       task.Title,
		Description: task.Description,
		DueDate:     task.DueDate,
		Status:      task.Status.String(),
		UserId:      task.UserId,
	}
}

func ToResponseTaskDTOs(tasks []models.Task) []ResponseTaskDTO {
	var dtos []ResponseTaskDTO
	for _, t := range tasks {
		dtos = append(dtos, ToResponseTaskDTO(t))
	}
	return dtos
}
