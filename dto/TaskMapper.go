package dto

import "draft-zadania-1/models"

func ToTask(dto UpdateTaskDTO) models.Task {
	return models.Task{
		Title:       dto.Title,
		Description: dto.Description,
		DueDate:     dto.DueDate,
		Status:      models.Status(dto.Status),
		UserId:      dto.UserId,
	}
}

func ToTaskCreate(dto CreateTaskDTO) models.Task {
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
		Title:       task.Title,
		Description: task.Description,
		DueDate:     task.DueDate,
		Status:      Status(task.Status),
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
