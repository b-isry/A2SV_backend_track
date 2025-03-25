package data

import (
	"context"
	"time"

	"github.com/Bisrat/task-manager/errors"
	"github.com/Bisrat/task-manager/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type TaskService struct {
	collection *mongo.Collection
}

func NewTaskService() *TaskService {
	return &TaskService{
		collection: GetCollection(),
	}
}

func validateTask(task models.Task) error {
	if task.Title == "" {
		return &errors.ValidationError{
			Field:   "title",
			Message: "title is required",
		}
	}
	if task.Description == "" {
		return &errors.ValidationError{
			Field:   "description",
			Message: "description is required",
		}
	}
	if task.DueDate.IsZero() {
		return &errors.ValidationError{
			Field:   "due_date",
			Message: "due date is required",
		}
	}
	if task.Status == "" {
		return &errors.ValidationError{
			Field:   "status",
			Message: "status is required",
		}
	}
	return nil
}

func (s *TaskService) GetAllTasks() ([]models.Task, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cursor, err := s.collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, &errors.DatabaseError{
			Operation: "find all tasks",
			Err:       err,
		}
	}
	defer cursor.Close(ctx)

	var tasks []models.Task
	if err = cursor.All(ctx, &tasks); err != nil {
		return nil, &errors.DatabaseError{
			Operation: "decode tasks",
			Err:       err,
		}
	}
	return tasks, nil
}

func (s *TaskService) GetTaskByID(id string) (models.Task, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return models.Task{}, &errors.ValidationError{
			Field:   "id",
			Message: "invalid task ID format",
		}
	}

	var task models.Task
	err = s.collection.FindOne(ctx, bson.M{"_id": objectID}).Decode(&task)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return models.Task{}, &errors.NotFoundError{
				Resource: "task",
				ID:       id,
			}
		}
		return models.Task{}, &errors.DatabaseError{
			Operation: "find task by ID",
			Err:       err,
		}
	}
	return task, nil
}

func (s *TaskService) CreateTask(task models.Task) error {
	if err := validateTask(task); err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err := s.collection.InsertOne(ctx, task)
	if err != nil {
		return &errors.DatabaseError{
			Operation: "create task",
			Err:       err,
		}
	}
	return nil
}

func (s *TaskService) UpdateTask(id string, updatedTask models.Task) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return &errors.ValidationError{
			Field:   "id",
			Message: "invalid task ID format",
		}
	}

	if updatedTask.Title == "" && updatedTask.Description == "" && updatedTask.Status == "" {
		return &errors.ValidationError{
			Field:   "update fields",
			Message: "at least one field must be provided for update",
		}
	}

	update := bson.M{}
	if updatedTask.Title != "" {
		if err := validateTask(models.Task{Title: updatedTask.Title}); err != nil {
			return err
		}
		update["title"] = updatedTask.Title
	}
	if updatedTask.Description != "" {
		if err := validateTask(models.Task{Description: updatedTask.Description}); err != nil {
			return err
		}
		update["description"] = updatedTask.Description
	}
	if updatedTask.Status != "" {
		if err := validateTask(models.Task{Status: updatedTask.Status}); err != nil {
			return err
		}
		update["status"] = updatedTask.Status
	}

	result, err := s.collection.UpdateOne(
		ctx,
		bson.M{"_id": objectID},
		bson.M{"$set": update},
	)
	if err != nil {
		return &errors.DatabaseError{
			Operation: "update task",
			Err:       err,
		}
	}

	if result.MatchedCount == 0 {
		return &errors.NotFoundError{
			Resource: "task",
			ID:       id,
		}
	}
	return nil
}

func (s *TaskService) DeleteTask(id string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return &errors.ValidationError{
			Field:   "id",
			Message: "invalid task ID format",
		}
	}

	result, err := s.collection.DeleteOne(ctx, bson.M{"_id": objectID})
	if err != nil {
		return &errors.DatabaseError{
			Operation: "delete task",
			Err:       err,
		}
	}

	if result.DeletedCount == 0 {
		return &errors.NotFoundError{
			Resource: "task",
			ID:       id,
		}
	}
	return nil
}
