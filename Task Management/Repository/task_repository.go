// Interface and implementation for task data access operations.

package repository

import (
	"context"
	"time"
	
	domain "Task-Management/Domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type taskRepository struct {
	collection *mongo.Collection
}

func NewTaskRepository(db *mongo.Database) domain.TaskRepository {
	return &taskRepository{
		collection: db.Collection(domain.TaskCollection),
	}
}

func (r *taskRepository) Create(ctx context.Context, task *domain.Task) (*domain.Task, error) {
	task.CreatedAt = time.Now()
	task.UpdatedAt = time.Now()

	result, err := r.collection.InsertOne(ctx, task)
	if err != nil {
		return nil, err
	}

	task.ID = result.InsertedID.(primitive.ObjectID)
	return task, nil
}

func (r *taskRepository) GetByID(ctx context.Context, id primitive.ObjectID) (*domain.Task, error) {
	var task domain.Task
	err := r.collection.FindOne(ctx, bson.M{"_id": id}).Decode(&task)
	if err != nil {
		return nil, err
	}
	return &task, nil
}

func (r *taskRepository) GetByUserID(ctx context.Context, userID primitive.ObjectID) ([]*domain.Task, error) {
	cursor, err := r.collection.Find(ctx, bson.M{"user_id": userID})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var tasks []*domain.Task
	if err := cursor.All(ctx, &tasks); err != nil {
		return nil, err
	}
	return tasks, nil
}

func (r *taskRepository) GetAll(ctx context.Context) ([]*domain.Task, error) {
	cursor, err := r.collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var tasks []*domain.Task
	if err := cursor.All(ctx, &tasks); err != nil {
		return nil, err
	}
	return tasks, nil
}

func (r *taskRepository) Update(ctx context.Context, task *domain.Task) error {
	task.UpdatedAt = time.Now()
	_, err := r.collection.UpdateOne(
		ctx,
		bson.M{"_id": task.ID},
		bson.M{"$set": task},
	)
	return err
}

func (r *taskRepository) Delete(ctx context.Context, id primitive.ObjectID) error {
	_, err := r.collection.DeleteOne(ctx, bson.M{"_id": id})
	return err
}

